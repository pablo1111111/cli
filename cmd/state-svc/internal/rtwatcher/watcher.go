package rtwatcher

import (
	"encoding/json"
	"os"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/ActiveState/cli/internal/analytics/client/sync"
	anaConst "github.com/ActiveState/cli/internal/analytics/constants"
	"github.com/ActiveState/cli/internal/analytics/dimensions"
	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/runbits/panics"
)

const defaultInterval = 1 * time.Minute
const CfgKey = "runtime-watchers"

type Watcher struct {
	an       *sync.Client
	cfg      *config.Instance
	watching []entry
	stop     chan struct{}
	interval time.Duration
}

func New(cfg *config.Instance, an *sync.Client) *Watcher {
	w := &Watcher{an: an, stop: make(chan struct{}, 1), cfg: cfg, interval: defaultInterval}

	if watchersJson := w.cfg.GetString(CfgKey); watchersJson != "" {
		watchers := []entry{}
		err := json.Unmarshal([]byte(watchersJson), &watchers)
		if err != nil {
			logging.Error("Could not unmarshal watchersL %s", errs.JoinMessage(err))
		} else {
			w.watching = watchers
		}
	}

	if v := os.Getenv(constants.HeartbeatIntervalEnvVarName); v != "" {
		vv, err := strconv.Atoi(v)
		if err != nil {
			logging.Warning("Invalid value for %s: %s", constants.HeartbeatIntervalEnvVarName, v)
		} else {
			w.interval = time.Duration(vv) * time.Millisecond
		}
	}

	go w.ticker(w.check)
	return w
}

func (w *Watcher) ticker(cb func()) {
	defer panics.LogPanics(recover(), debug.Stack())

	ticker := time.NewTicker(w.interval)
	for {
		select {
		case <-ticker.C:
			cb()
		case <-w.stop:
			return
		}
	}
}

func (w *Watcher) check() {
	logging.Debug("Checking for runtime processes")

	for i := range w.watching {
		e := w.watching[i] // Must use index, because we are deleting indexes further down
		running, err := e.IsRunning()
		if err != nil {
			logging.Error("Could not check if runtime process is running: %s", errs.JoinMessage(err))
			// Don't return yet, the conditional below still needs to clear this entry
		}
		if !running {
			w.watching = append(w.watching[:i], w.watching[i+1:]...)
			continue
		}

		w.RecordUsage(e)
	}
}

func (w *Watcher) RecordUsage(e entry) {
	logging.Debug("Recording usage of %s (%d)", e.Exec, e.PID)
	w.an.Event(anaConst.CatRuntimeUsage, anaConst.ActRuntimeHeartbeat, e.Dims)
}

func (w *Watcher) Close() error {
	logging.Debug("Closing runtime watcher")

	w.stop <- struct{}{}

	if len(w.watching) > 0 {
		watchingJson, err := json.Marshal(w.watching)
		if err != nil {
			return errs.Wrap(err, "Could not marshal watchers")
		}
		return w.cfg.Set(CfgKey, watchingJson)
	}

	return nil
}

func (w *Watcher) Watch(pid int, exec string, dims *dimensions.Values) {
	logging.Debug("Watching %s (%d)", exec, pid)
	e := entry{pid, exec, dims}
	w.watching = append(w.watching, e)
	w.RecordUsage(e)
}
