// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type GrandCompany struct {
	GameContentLinks struct {
		BuddyEquip struct {
			GrandCompany []interface{} `json:"GrandCompany,omitempty"`
		}
		GCScripShopCategory struct {
			GrandCompany []interface{} `json:"GrandCompany,omitempty"`
		}
		GCShop struct {
			GrandCompany []interface{} `json:"GrandCompany,omitempty"`
		}
		Item struct {
			GrandCompany []interface{} `json:"GrandCompany,omitempty"`
		}
		Quest struct {
			GrandCompany []interface{} `json:"GrandCompany,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	Url string `json:"Url,omitempty"`
}