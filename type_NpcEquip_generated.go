// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type NpcEquip struct {
	DyeLeftRing int `json:"DyeLeftRing,omitempty"`
	GameContentLinks struct {
		BNpcBase struct {
			NpcEquip []interface{} `json:"NpcEquip,omitempty"`
		}
		ENpcBase struct {
			NpcEquip []interface{} `json:"NpcEquip,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	ModelHands int `json:"ModelHands,omitempty"`
	DyeHead int `json:"DyeHead,omitempty"`
	ModelNeck int `json:"ModelNeck,omitempty"`
	DyeNeck int `json:"DyeNeck,omitempty"`
	DyeMainHand int `json:"DyeMainHand,omitempty"`
	ModelLeftRing int `json:"ModelLeftRing,omitempty"`
	ModelMainHand string `json:"ModelMainHand,omitempty"`
	Url string `json:"Url,omitempty"`
	Visor int `json:"Visor,omitempty"`
	DyeHands int `json:"DyeHands,omitempty"`
	ModelRightRing int `json:"ModelRightRing,omitempty"`
	ModelOffHand string `json:"ModelOffHand,omitempty"`
	ModelWrists int `json:"ModelWrists,omitempty"`
	DyeRightRing int `json:"DyeRightRing,omitempty"`
	ModelEars int `json:"ModelEars,omitempty"`
	ModelLegs int `json:"ModelLegs,omitempty"`
	DyeEars int `json:"DyeEars,omitempty"`
	DyeOffHand int `json:"DyeOffHand,omitempty"`
	DyeWrists int `json:"DyeWrists,omitempty"`
	ModelBody int `json:"ModelBody,omitempty"`
	DyeLegs int `json:"DyeLegs,omitempty"`
	DyeFeet int `json:"DyeFeet,omitempty"`
	ModelFeet int `json:"ModelFeet,omitempty"`
	ModelHead int `json:"ModelHead,omitempty"`
	DyeBody int `json:"DyeBody,omitempty"`
}