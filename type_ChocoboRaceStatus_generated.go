// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type ChocoboRaceStatus struct {
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	ID int `json:"ID,omitempty"`
	Status struct {
		MaxStacks int `json:"MaxStacks,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		Transfiguration int `json:"Transfiguration,omitempty"`
		Description string `json:"Description,omitempty"`
		Description_fr string `json:"Description_fr,omitempty"`
		IconID int `json:"IconID,omitempty"`
		LockControl int `json:"LockControl,omitempty"`
		LockMovement int `json:"LockMovement,omitempty"`
		IsPermanent int `json:"IsPermanent,omitempty"`
		LockActions int `json:"LockActions,omitempty"`
		VFX int `json:"VFX,omitempty"`
		CanDispel int `json:"CanDispel,omitempty"`
		Description_en string `json:"Description_en,omitempty"`
		ID int `json:"ID,omitempty"`
		InflictedByActor int `json:"InflictedByActor,omitempty"`
		Invisibility int `json:"Invisibility,omitempty"`
		IsFcBuff int `json:"IsFcBuff,omitempty"`
		Name string `json:"Name,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		Category int `json:"Category,omitempty"`
		Description_de string `json:"Description_de,omitempty"`
		Description_ja string `json:"Description_ja,omitempty"`
		HitEffect int `json:"HitEffect,omitempty"`
		Icon string `json:"Icon,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
	}
	StatusTarget string `json:"StatusTarget,omitempty"`
	StatusTargetID int `json:"StatusTargetID,omitempty"`
	Url string `json:"Url,omitempty"`
}
