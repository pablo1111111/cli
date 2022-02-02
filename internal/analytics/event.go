package analytics

import "github.com/ActiveState/cli/internal/analytics/dimensions"

type Events []Event

type Event struct {
	Category   string             `json:"category"`
	Action     string             `json:"action"`
	Label      string             `json:"label"`
	Dimensions *dimensions.Values `json:"dimensions"`
}
