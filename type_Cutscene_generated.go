// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type Cutscene struct {
	Path_en string `json:"Path_en,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks struct {
		CompleteJournal struct {
			Cutscene0 []interface{} `json:"Cutscene0,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Path string `json:"Path,omitempty"`
}
