// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type EurekaAethernet struct {
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	ID int `json:"ID,omitempty"`
	Location struct {
		NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
		NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
		NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		ID int `json:"ID,omitempty"`
		Icon string `json:"Icon,omitempty"`
		Name string `json:"Name,omitempty"`
		NameNoArticle string `json:"NameNoArticle,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
	}
	LocationTarget string `json:"LocationTarget,omitempty"`
	LocationTargetID int `json:"LocationTargetID,omitempty"`
	Url string `json:"Url,omitempty"`
}