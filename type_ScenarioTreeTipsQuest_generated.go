// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type ScenarioTreeTipsQuest struct {
	LevelTarget string `json:"LevelTarget,omitempty"`
	LevelTargetID int `json:"LevelTargetID,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks struct {
		ScenarioTreeTips struct {
			Tips1 []interface{} `json:"Tips1,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Level int `json:"Level,omitempty"`
}