package contracts

import (
	_ "embed"
	"encoding/json"
)

//go:embed constants.json
var rawConstants []byte

type RegencyDef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ViolationDef struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Desc  string `json:"desc"`
}

type UploadLimits struct {
	MaxImageSizeBytes int64    `json:"max_image_size_bytes"`
	MaxDocSizeBytes   int64    `json:"max_doc_size_bytes"`
	AllowedImageMimes []string `json:"allowed_image_mimes"`
	AllowedDocMimes   []string `json:"allowed_doc_mimes"`
}

type SharedConstants struct {
	Roles             map[string]string `json:"roles"`
	ComplaintStatuses map[string]string `json:"complaint_statuses"`
	ViolationOptions  []ViolationDef    `json:"violation_options"`
	Regencies         []RegencyDef      `json:"regencies"`
	UploadLimits      UploadLimits      `json:"upload_limits"`
}

var Shared SharedConstants

func init() {
	if err := json.Unmarshal(rawConstants, &Shared); err != nil {
		panic("failed to parse embedded contracts/constants.json: " + err.Error())
	}
}
