// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type TerritoryType struct {
	Aetheryte int `json:"Aetheryte,omitempty"`
	GameContentLinks struct {
		Aetheryte struct {
			Territory []interface{} `json:"Territory,omitempty"`
		}
		FishParameter struct {
			TerritoryType []interface{} `json:"TerritoryType,omitempty"`
		}
		GatheringPoint struct {
			TerritoryType []interface{} `json:"TerritoryType,omitempty"`
		}
		InstanceContent struct {
			TerritoryType []interface{} `json:"TerritoryType,omitempty"`
		}
		Level struct {
			Territory []interface{} `json:"Territory,omitempty"`
		}
		Map struct {
			TerritoryType []interface{} `json:"TerritoryType,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Url string `json:"Url,omitempty"`
	WeatherRate int `json:"WeatherRate,omitempty"`
	ArrayEventHandler int `json:"ArrayEventHandler,omitempty"`
	Bg string `json:"Bg,omitempty"`
	Bg_en string `json:"Bg_en,omitempty"`
	Name string `json:"Name,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	Map int `json:"Map,omitempty"`
	PlaceName int `json:"PlaceName,omitempty"`
	PlaceNameRegion int `json:"PlaceNameRegion,omitempty"`
	PlaceNameZone int `json:"PlaceNameZone,omitempty"`
	TerritoryIntendedUse int `json:"TerritoryIntendedUse,omitempty"`
}
