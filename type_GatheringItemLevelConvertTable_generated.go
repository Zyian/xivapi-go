// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type GatheringItemLevelConvertTable struct {
	ID int `json:"ID,omitempty"`
	Stars int `json:"Stars,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks struct {
		GatheringItem struct {
			GatheringItemLevel []interface{} `json:"GatheringItemLevel,omitempty"`
		}
	}
	GatheringItemLevel int `json:"GatheringItemLevel,omitempty"`
}