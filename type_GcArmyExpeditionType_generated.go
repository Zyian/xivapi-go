// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type GcArmyExpeditionType struct {
	Name_fr string `json:"Name_fr,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks struct {
		GcArmyExpedition struct {
			GcArmyExpeditionType []interface{} `json:"GcArmyExpeditionType,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
}
