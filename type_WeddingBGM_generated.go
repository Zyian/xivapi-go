// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type WeddingBGM struct {
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	SongName_ja string `json:"SongName_ja,omitempty"`
	SongTarget string `json:"SongTarget,omitempty"`
	SongName_en string `json:"SongName_en,omitempty"`
	SongName_fr string `json:"SongName_fr,omitempty"`
	SongTargetID int `json:"SongTargetID,omitempty"`
	Url string `json:"Url,omitempty"`
	ID int `json:"ID,omitempty"`
	Song struct {
		File_en string `json:"File_en,omitempty"`
		ID int `json:"ID,omitempty"`
		File string `json:"File,omitempty"`
	}
	SongName string `json:"SongName,omitempty"`
	SongName_de string `json:"SongName_de,omitempty"`
}