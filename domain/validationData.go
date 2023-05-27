package domain

import "sync"

var ValidErrors = map[int]struct{}{
	-50024: {},
}

type ValidationData struct {
	DataVersion        float32                 `json:"dataVersion,omitempty"`
	URL                string                  `json:"url,omitempty"`
	Messages           []ValidationDataMessage `json:"messages,omitempty"`
	ParseFailed        bool                    `json:"parseFailed,omitempty"`
	DataID             int                     `json:"dataID,omitempty"`
	ValidatorVersion   string                  `json:"validatorVersion,omitempty"`
	ValidatorTimestamp string                  `json:"validatorTimestamp,omitempty"`
	PlaylistKind       string                  `json:"playlistKind,omitempty"`
}

type ValidationDataMessage struct {
	ErrorComment          string `json:"errorComment,omitempty"`
	ErrorDomain           string `json:"errorDomain,omitempty"`
	ErrorStatusCode       int    `json:"errorStatusCode,omitempty"`
	ErrorRequirementLevel int    `json:"errorRequirementLevel,omitempty"`
	ErrorDetail           string `json:"errorDetail,omitempty"`
}

type SafeValidationMap struct {
	Mu            sync.Mutex
	ValidationMap map[string]bool
}
