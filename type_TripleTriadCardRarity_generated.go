// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type TripleTriadCardRarity struct {
	GameContentLinks struct {
		TripleTriadCardResident struct {
			TripleTriadCardRarity []interface{} `json:"TripleTriadCardRarity,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Stars int `json:"Stars,omitempty"`
	Url string `json:"Url,omitempty"`
}
