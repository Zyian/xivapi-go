// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type FieldMarker struct {
	VFXTargetID int `json:"VFXTargetID,omitempty"`
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	ID int `json:"ID,omitempty"`
	Icon string `json:"Icon,omitempty"`
	IconID int `json:"IconID,omitempty"`
	Url string `json:"Url,omitempty"`
	VFX struct {
		ID int `json:"ID,omitempty"`
		Location string `json:"Location,omitempty"`
		Location_en string `json:"Location_en,omitempty"`
	}
	VFXTarget string `json:"VFXTarget,omitempty"`
}
