// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type TripleTriadRule struct {
	Name string `json:"Name,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks struct {
		TripleTriad struct {
			TripleTriadRule0 []interface{} `json:"TripleTriadRule0,omitempty"`
			TripleTriadRule1 []interface{} `json:"TripleTriadRule1,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
}
