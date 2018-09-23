// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type Achievement struct {
	Item int `json:"Item,omitempty"`
	PostAchievements []interface{} `json:"PostAchievements,omitempty"`
	QuestRequirements []interface{} `json:"QuestRequirements,omitempty"`
	Type int `json:"Type,omitempty"`
	Description string `json:"Description,omitempty"`
	IconID int `json:"IconID,omitempty"`
	Data3 int `json:"Data3,omitempty"`
	Description_de string `json:"Description_de,omitempty"`
	Icon string `json:"Icon,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	PreAchievements []interface{} `json:"PreAchievements,omitempty"`
	AchievementCategoryTarget string `json:"AchievementCategoryTarget,omitempty"`
	ClassJobRequirements interface{} `json:"ClassJobRequirements,omitempty"`
	Data2 int `json:"Data2,omitempty"`
	Description_ja string `json:"Description_ja,omitempty"`
	ID int `json:"ID,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
	Data0 int `json:"Data0,omitempty"`
	Data1 int `json:"Data1,omitempty"`
	Points int `json:"Points,omitempty"`
	QuestRequirementsAll bool `json:"QuestRequirementsAll,omitempty"`
	Data4 int `json:"Data4,omitempty"`
	Data7 int `json:"Data7,omitempty"`
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	Name string `json:"Name,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Url string `json:"Url,omitempty"`
	Data5 int `json:"Data5,omitempty"`
	Data6 int `json:"Data6,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	Title int `json:"Title,omitempty"`
	AchievementCategoryTargetID int `json:"AchievementCategoryTargetID,omitempty"`
	Data8 int `json:"Data8,omitempty"`
	GamePatch struct {
		ExVersion int `json:"ExVersion,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		Banner string `json:"Banner,omitempty"`
		ID int `json:"ID,omitempty"`
		Name string `json:"Name,omitempty"`
		Name_cn string `json:"Name_cn,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		Name_kr string `json:"Name_kr,omitempty"`
		ReleaseDate int `json:"ReleaseDate,omitempty"`
		Version int `json:"Version,omitempty"`
	}
	Description_fr string `json:"Description_fr,omitempty"`
	Order int `json:"Order,omitempty"`
	AchievementCategory struct {
		AchievementKindTarget string `json:"AchievementKindTarget,omitempty"`
		AchievementKindTargetID int `json:"AchievementKindTargetID,omitempty"`
		ID int `json:"ID,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		AchievementKind struct {
			Name string `json:"Name,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			ID int `json:"ID,omitempty"`
		}
		Name string `json:"Name,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
	}
	Description_en string `json:"Description_en,omitempty"`
}
