// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type DeepDungeonStatus struct {
	Url string `json:"Url,omitempty"`
	LogMessageTarget string `json:"LogMessageTarget,omitempty"`
	Name struct {
		Description_ja string `json:"Description_ja,omitempty"`
		ID int `json:"ID,omitempty"`
		Icon string `json:"Icon,omitempty"`
		Name string `json:"Name,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		Description_de string `json:"Description_de,omitempty"`
		Description_en string `json:"Description_en,omitempty"`
		Description_fr string `json:"Description_fr,omitempty"`
		IconID int `json:"IconID,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		Description string `json:"Description,omitempty"`
	}
	ScreenImageTarget string `json:"ScreenImageTarget,omitempty"`
	ScreenImageTargetID int `json:"ScreenImageTargetID,omitempty"`
	NameTarget string `json:"NameTarget,omitempty"`
	NameTargetID int `json:"NameTargetID,omitempty"`
	ScreenImage struct {
		ImageID int `json:"ImageID,omitempty"`
		ID int `json:"ID,omitempty"`
		Image string `json:"Image,omitempty"`
	}
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	ID int `json:"ID,omitempty"`
	LogMessage struct {
		LogKindTarget string `json:"LogKindTarget,omitempty"`
		LogKindTargetID int `json:"LogKindTargetID,omitempty"`
		Text string `json:"Text,omitempty"`
		Text_en string `json:"Text_en,omitempty"`
		ID int `json:"ID,omitempty"`
		Text_de string `json:"Text_de,omitempty"`
		Text_fr string `json:"Text_fr,omitempty"`
		Text_ja string `json:"Text_ja,omitempty"`
		LogKind struct {
			Format_en string `json:"Format_en,omitempty"`
			Format_fr string `json:"Format_fr,omitempty"`
			ID int `json:"ID,omitempty"`
			LogKindCategoryText int `json:"LogKindCategoryText,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Example_de string `json:"Example_de,omitempty"`
			Example_en string `json:"Example_en,omitempty"`
			Format string `json:"Format,omitempty"`
			Format_ja string `json:"Format_ja,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Example_ja string `json:"Example_ja,omitempty"`
			Format_de string `json:"Format_de,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			Example string `json:"Example,omitempty"`
			Example_fr string `json:"Example_fr,omitempty"`
		}
	}
	LogMessageTargetID int `json:"LogMessageTargetID,omitempty"`
}
