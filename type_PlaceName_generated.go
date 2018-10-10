// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type PlaceName struct {
	Name string `json:"Name,omitempty"`
	NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks struct {
		Aetheryte struct {
			PlaceName []interface{} `json:"PlaceName,omitempty"`
		}
		Leve struct {
			PlaceNameIssued []interface{} `json:"PlaceNameIssued,omitempty"`
			PlaceNameStart []interface{} `json:"PlaceNameStart,omitempty"`
		}
		Map struct {
			PlaceName []interface{} `json:"PlaceName,omitempty"`
		}
		MapMarker struct {
			PlaceNameSubtext []interface{} `json:"PlaceNameSubtext,omitempty"`
		}
		TerritoryType struct {
			PlaceName []interface{} `json:"PlaceName,omitempty"`
		}
	}
	GamePatch interface{} `json:"GamePatch,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	Icon string `json:"Icon,omitempty"`
	NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
	ID int `json:"ID,omitempty"`
	NameNoArticle string `json:"NameNoArticle,omitempty"`
	NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
	NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
}