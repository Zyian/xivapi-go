// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type GCScripShopCategory struct {
	GrandCompany struct {
		Name_de string `json:"Name_de,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		ID int `json:"ID,omitempty"`
		Name string `json:"Name,omitempty"`
	}
	GrandCompanyTarget string `json:"GrandCompanyTarget,omitempty"`
	GrandCompanyTargetID int `json:"GrandCompanyTargetID,omitempty"`
	ID int `json:"ID,omitempty"`
	SubCategory int `json:"SubCategory,omitempty"`
	Tier int `json:"Tier,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
}
