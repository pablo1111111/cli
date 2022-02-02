package dimensions

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/installation/storage"
	"github.com/ActiveState/cli/internal/instanceid"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/machineid"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/rtutils/p"
	"github.com/ActiveState/cli/internal/singleton/uniqid"
	"github.com/ActiveState/cli/pkg/platform/authentication"
	"github.com/ActiveState/cli/pkg/sysinfo"
	"github.com/imdario/mergo"
)

type Values struct {
	Version          *string `json:"version,omitempty"`
	BranchName       *string `json:"branchName,omitempty"`
	UserID           *string `json:"userID,omitempty"`
	OSName           *string `json:"osName,omitempty"`
	OSVersion        *string `json:"osVersion,omitempty"`
	InstallSource    *string `json:"installSource,omitempty"`
	MachineID        *string `json:"machineID,omitempty"`
	UniqID           *string `json:"uniqID,omitempty"`
	SessionToken     *string `json:"sessionToken,omitempty"`
	UpdateTag        *string `json:"updateTag,omitempty"`
	ProjectNameSpace *string `json:"projectNameSpace,omitempty"`
	OutputType       *string `json:"outputType,omitempty"`
	ProjectID        *string `json:"projectID,omitempty"`
	Flags            *string `json:"flags,omitempty"`
	Trigger          *string `json:"trigger,omitempty"`
	Headless         *string `json:"headless,omitempty"`
	InstanceID       *string `json:"instanceID,omitempty"`
	CommitID         *string `json:"commitID,omitempty"`
	Timestamp        *string `json:"timestamp,omitempty"`

	preProcessor func(*Values) error
}

func NewDefaultDimensions(pjNamespace, sessionToken, updateTag string) *Values {
	installSource, err := storage.InstallSource()
	if err != nil {
		logging.Error("Could not detect installSource: %s", errs.Join(err, " :: ").Error())
	}

	machineID := machineid.UniqID()
	if machineID == machineid.UnknownID || machineID == machineid.FallbackID {
		logging.Error("unknown machine id: %s", machineID)
	}
	deviceID := uniqid.Text()

	var userIDString string
	userID := authentication.LegacyGet().UserID()
	if userID != nil {
		userIDString = userID.String()
	}

	osName := sysinfo.OS().String()
	osVersion := "unknown"
	osvInfo, err := sysinfo.OSVersion()
	if err != nil {
		logging.Errorf("Could not detect osVersion: %v", err)
	}
	if osvInfo != nil {
		osVersion = osvInfo.Version
	}

	return &Values{
		p.StrP(constants.Version),
		p.StrP(constants.BranchName),
		p.StrP(userIDString),
		p.StrP(osName),
		p.StrP(osVersion),
		p.StrP(installSource),
		p.StrP(machineID),
		p.StrP(deviceID),
		p.StrP(sessionToken),
		p.StrP(updateTag),
		p.StrP(pjNamespace),
		p.StrP(string(output.PlainFormatName)),
		p.StrP(""),
		p.StrP(CalculateFlags()),
		p.StrP(""),
		p.StrP(""),
		p.StrP(instanceid.ID()),
		p.StrP(""),
		p.StrP(time.Now().String()),
		nil,
	}
}

func (m *Values) Merge(mergeWith ...*Values) {
	for _, dim := range mergeWith {
		if err := mergo.Merge(m, dim, mergo.WithOverride); err != nil {
			logging.Critical("Could not merge dimension maps: %s", errs.JoinMessage(err))
		}
	}
}

func (v *Values) RegisterPreProcessor(f func(*Values) error) {
	v.preProcessor = f
}

func (v *Values) PreProcess() error {
	if v.preProcessor != nil {
		if err := v.preProcessor(v); err != nil {
			return errs.Wrap(err, "PreProcessor failed: %s", errs.JoinMessage(err))
		}
	}

	if p.PStr(v.UniqID) == machineid.FallbackID {
		return errs.New("machine id was set to fallback id when creating analytics event")
	}

	return nil
}

func (v *Values) Marshal() (string, error) {
	dimMarshalled, err := json.Marshal(v)
	if err != nil {
		return "", errs.Wrap(err, "Could not marshal dimensions")
	}
	return string(dimMarshalled), nil
}

func CalculateFlags() string {
	flags := []string{}
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		}
	}
	return strings.Join(flags, " ")
}
