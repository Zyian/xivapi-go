// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type GCShop struct {
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	GrandCompany struct {
		Name_en string `json:"Name_en,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		ID int `json:"ID,omitempty"`
		Name string `json:"Name,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
	}
	GrandCompanyTarget string `json:"GrandCompanyTarget,omitempty"`
	GrandCompanyTargetID int `json:"GrandCompanyTargetID,omitempty"`
	ID int `json:"ID,omitempty"`
	Url string `json:"Url,omitempty"`
}
