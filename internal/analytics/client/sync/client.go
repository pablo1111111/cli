package sync

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"runtime/debug"
	"sync"

	"github.com/ActiveState/cli/internal/analytics"
	anaConsts "github.com/ActiveState/cli/internal/analytics/constants"
	"github.com/ActiveState/cli/internal/analytics/dimensions"
	"github.com/ActiveState/cli/internal/condition"
	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/installation/storage"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/machineid"
	"github.com/ActiveState/cli/internal/rtutils/p"
	"github.com/ActiveState/cli/internal/singleton/uniqid"
	"github.com/ActiveState/cli/internal/updater"
	"github.com/ActiveState/cli/pkg/platform/authentication"
	"github.com/ActiveState/cli/pkg/sysinfo"
)

type Reporter interface {
	ID() string
	Event(category, action, label string, dimensions *dimensions.Values) error
}

// Client instances send analytics events to GA and S3 endpoints without delay. It is only supposed to be used inside the `state-svc`.  All other processes should use the DefaultClient.
type Client struct {
	customDimensions *dimensions.Values
	eventWaitGroup   *sync.WaitGroup
	reporters        []Reporter
	url              string
	client           *http.Client
}

var _ analytics.Dispatcher = &Client{}

// New initializes the analytics instance with all custom dimensions known at this time
func New(cfg *config.Instance, auth *authentication.Auth) *Client {
	a := &Client{
		eventWaitGroup: &sync.WaitGroup{},
		// This is a temporary URL
		url:    "https://platform.activestate.com/state-analytics/events",
		client: &http.Client{},
	}

	installSource, err := storage.InstallSource()
	if err != nil {
		logging.Error("Could not detect installSource: %s", errs.Join(err, " :: ").Error())
	}

	machineID := machineid.UniqID()
	if machineID == machineid.UnknownID || machineID == machineid.FallbackID {
		logging.Error("unknown machine id: %s", machineID)
	}
	deviceID := uniqid.Text()

	osName := sysinfo.OS().String()
	osVersion := "unknown"
	osvInfo, err := sysinfo.OSVersion()
	if err != nil {
		logging.Errorf("Could not detect osVersion: %v", err)
	}
	if osvInfo != nil {
		osVersion = osvInfo.Version
	}

	var sessionToken string
	var tag string
	if cfg != nil {
		sessionToken = cfg.GetString(anaConsts.CfgSessionToken)
		var ok bool
		tag, ok = os.LookupEnv(constants.UpdateTagEnvVarName)
		if !ok {
			tag = cfg.GetString(updater.CfgUpdateTag)
		}
	}

	userID := ""
	if auth != nil && auth.UserID() != nil {
		userID = string(*auth.UserID())
	}

	customDimensions := &dimensions.Values{
		Version:       p.StrP(constants.Version),
		BranchName:    p.StrP(constants.BranchName),
		OSName:        p.StrP(osName),
		OSVersion:     p.StrP(osVersion),
		InstallSource: p.StrP(installSource),
		MachineID:     p.StrP(machineID),
		UniqID:        p.StrP(deviceID),
		SessionToken:  p.StrP(sessionToken),
		UpdateTag:     p.StrP(tag),
		UserID:        p.StrP(userID),
		Flags:         p.StrP(dimensions.CalculateFlags()),
	}

	a.customDimensions = customDimensions

	return a
}

func (a *Client) NewReporter(rep Reporter) {
	a.reporters = append(a.reporters, rep)
}

func (a *Client) Wait() {
	a.eventWaitGroup.Wait()
}

func (a *Client) report(category, action, label string, d *dimensions.Values) error {
	e := analytics.Event{category, action, label, d}
	evs := analytics.Events{[]analytics.Event{e}}
	b, err := json.Marshal(evs)
	if err != nil {
		return errs.Wrap(err, "Could not marshall analytics event")
	}

	req, err := http.NewRequest(http.MethodPost, a.url, bytes.NewBuffer(b))
	if err != nil {
		return errs.Wrap(err, "Could not create new HTTP request")
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return errs.Wrap(err, "Could not send analytics request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errs.New("Analytics proxy responded with unexected HTTP code: %d", resp.StatusCode)
	}

	return nil
}

func (a *Client) Event(category string, action string, dims ...*dimensions.Values) {
	a.EventWithLabel(category, action, "", dims...)
}

func (a *Client) EventWithLabel(category string, action, label string, dims ...*dimensions.Values) {
	if a.customDimensions == nil {
		if condition.InUnitTest() {
			return
		}
		if !condition.BuiltViaCI() {
			panic("Trying to send analytics without configuring the Analytics instance.")
		}
		logging.Critical("Trying to send analytics event without configuring the Analytics instance.")
		return
	}

	_actualDims := *a.customDimensions
	actualDims := &_actualDims
	for _, dim := range dims {
		actualDims.Merge(dim)
	}

	if err := actualDims.PreProcess(); err != nil {
		logging.Critical("Analytics dimensions cannot be processed properly: %s", errs.JoinMessage(err))
	}

	a.eventWaitGroup.Add(1)
	// We do not wait for the events to be processed, just scheduling them
	go func() {
		defer a.eventWaitGroup.Done()
		defer handlePanics(recover(), debug.Stack())
		a.report(category, action, label, actualDims)
	}()
}

func handlePanics(err interface{}, stack []byte) {
	if err == nil {
		return
	}
	logging.Error("Panic in state-svc analytics: %v", err)
	logging.Debug("Stack: %s", string(stack))
}
