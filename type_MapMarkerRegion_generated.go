// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type MapMarkerRegion struct {
	GameContentLinks struct {
		MapMarker struct {
			MapMarkerRegion []interface{} `json:"MapMarkerRegion,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Url string `json:"Url,omitempty"`
	X int `json:"X,omitempty"`
}
