// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type QuickChatTransient struct {
	ID int `json:"ID,omitempty"`
	TextOutput string `json:"TextOutput,omitempty"`
	TextOutput_de string `json:"TextOutput_de,omitempty"`
	TextOutput_en string `json:"TextOutput_en,omitempty"`
	TextOutput_fr string `json:"TextOutput_fr,omitempty"`
	TextOutput_ja string `json:"TextOutput_ja,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks struct {
		QuickChat struct {
			QuickChatTransient []interface{} `json:"QuickChatTransient,omitempty"`
		}
	}
}
