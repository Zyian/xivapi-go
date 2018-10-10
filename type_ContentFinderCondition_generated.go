// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type ContentFinderCondition struct {
	ContentMemberTypeTarget string `json:"ContentMemberTypeTarget,omitempty"`
	Description string `json:"Description,omitempty"`
	Description_de string `json:"Description_de,omitempty"`
	ContentTypeTarget string `json:"ContentTypeTarget,omitempty"`
	InstanceContentTargetID int `json:"InstanceContentTargetID,omitempty"`
	Name string `json:"Name,omitempty"`
	ContentTypeTargetID int `json:"ContentTypeTargetID,omitempty"`
	Icon string `json:"Icon,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	ClassJobLevelRequired int `json:"ClassJobLevelRequired,omitempty"`
	ContentLinkType int `json:"ContentLinkType,omitempty"`
	Description_ja string `json:"Description_ja,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	ContentType struct {
		Name_fr string `json:"Name_fr,omitempty"`
		ID int `json:"ID,omitempty"`
		Icon string `json:"Icon,omitempty"`
		IconDutyFinder string `json:"IconDutyFinder,omitempty"`
		IconDutyFinderID int `json:"IconDutyFinderID,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		IconID int `json:"IconID,omitempty"`
		Name string `json:"Name,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
	}
	InstanceContentTarget string `json:"InstanceContentTarget,omitempty"`
	ItemLevelRequired int `json:"ItemLevelRequired,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	ClassJobLevelSync int `json:"ClassJobLevelSync,omitempty"`
	ContentMemberTypeTargetID int `json:"ContentMemberTypeTargetID,omitempty"`
	ContentMemberType struct {
		TanksPerParty int `json:"TanksPerParty,omitempty"`
		HealersPerParty int `json:"HealersPerParty,omitempty"`
		ID int `json:"ID,omitempty"`
		MeleesPerParty int `json:"MeleesPerParty,omitempty"`
		RangedPerParty int `json:"RangedPerParty,omitempty"`
	}
	Description_en string `json:"Description_en,omitempty"`
	Description_fr string `json:"Description_fr,omitempty"`
	InstanceContent struct {
		BGM struct {
			File string `json:"File,omitempty"`
			File_en string `json:"File_en,omitempty"`
			ID int `json:"ID,omitempty"`
		}
		BossCurrencyC0 int `json:"BossCurrencyC0,omitempty"`
		BossCurrencyC1 int `json:"BossCurrencyC1,omitempty"`
		PartyCondition int `json:"PartyCondition,omitempty"`
		BossCurrencyC3 int `json:"BossCurrencyC3,omitempty"`
		FinalBossCurrencyB int `json:"FinalBossCurrencyB,omitempty"`
		InstanceContentBuff int `json:"InstanceContentBuff,omitempty"`
		NewPlayerBonusA int `json:"NewPlayerBonusA,omitempty"`
		TerritoryTypeTargetID int `json:"TerritoryTypeTargetID,omitempty"`
		TerritoryType struct {
			PlaceNameRegion int `json:"PlaceNameRegion,omitempty"`
			TerritoryIntendedUse int `json:"TerritoryIntendedUse,omitempty"`
			WeatherRate int `json:"WeatherRate,omitempty"`
			ArrayEventHandler int `json:"ArrayEventHandler,omitempty"`
			Bg_en string `json:"Bg_en,omitempty"`
			Map int `json:"Map,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			PlaceNameZone int `json:"PlaceNameZone,omitempty"`
			Aetheryte int `json:"Aetheryte,omitempty"`
			Bg string `json:"Bg,omitempty"`
			ID int `json:"ID,omitempty"`
		}
		TimeLimitMin int `json:"TimeLimitMin,omitempty"`
		BNpcBaseBossTargetID int `json:"BNpcBaseBossTargetID,omitempty"`
		BossCurrencyA0 int `json:"BossCurrencyA0,omitempty"`
		BossCurrencyC4 int `json:"BossCurrencyC4,omitempty"`
		InstanceContentTextDataObjectiveEnd struct {
			Text string `json:"Text,omitempty"`
			Text_de string `json:"Text_de,omitempty"`
			Text_en string `json:"Text_en,omitempty"`
			Text_fr string `json:"Text_fr,omitempty"`
			Text_ja string `json:"Text_ja,omitempty"`
			ID int `json:"ID,omitempty"`
		}
		InstanceContentTextDataObjectiveStartTarget string `json:"InstanceContentTextDataObjectiveStartTarget,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		BossCurrencyA4 int `json:"BossCurrencyA4,omitempty"`
		InstanceContentTextDataBossEnd int `json:"InstanceContentTextDataBossEnd,omitempty"`
		InstanceContentTypeTarget string `json:"InstanceContentTypeTarget,omitempty"`
		InstanceContentTypeTargetID int `json:"InstanceContentTypeTargetID,omitempty"`
		Name string `json:"Name,omitempty"`
		InstanceContentTextDataObjectiveStart struct {
			Text_fr string `json:"Text_fr,omitempty"`
			Text_ja string `json:"Text_ja,omitempty"`
			ID int `json:"ID,omitempty"`
			Text string `json:"Text,omitempty"`
			Text_de string `json:"Text_de,omitempty"`
			Text_en string `json:"Text_en,omitempty"`
		}
		InstanceContentTextDataObjectiveStartTargetID int `json:"InstanceContentTextDataObjectiveStartTargetID,omitempty"`
		BGMTarget string `json:"BGMTarget,omitempty"`
		BossCurrencyA3 int `json:"BossCurrencyA3,omitempty"`
		BossCurrencyB0 int `json:"BossCurrencyB0,omitempty"`
		BossCurrencyB2 int `json:"BossCurrencyB2,omitempty"`
		FinalBossCurrencyC int `json:"FinalBossCurrencyC,omitempty"`
		InstanceContentTextDataObjectiveEndTarget string `json:"InstanceContentTextDataObjectiveEndTarget,omitempty"`
		WeekRestriction int `json:"WeekRestriction,omitempty"`
		BossCurrencyA1 int `json:"BossCurrencyA1,omitempty"`
		BossCurrencyB4 int `json:"BossCurrencyB4,omitempty"`
		BossCurrencyC2 int `json:"BossCurrencyC2,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		NewPlayerBonusB int `json:"NewPlayerBonusB,omitempty"`
		InstanceContentType struct {
			ID int `json:"ID,omitempty"`
		}
		Name_de string `json:"Name_de,omitempty"`
		BGMTargetID int `json:"BGMTargetID,omitempty"`
		BNpcBaseBoss struct {
			ArrayEventHandler int `json:"ArrayEventHandler,omitempty"`
			BNpcCustomize int `json:"BNpcCustomize,omitempty"`
			ID int `json:"ID,omitempty"`
			ModelChara int `json:"ModelChara,omitempty"`
			NpcEquip int `json:"NpcEquip,omitempty"`
			Scale int `json:"Scale,omitempty"`
		}
		BossCurrencyA2 int `json:"BossCurrencyA2,omitempty"`
		BossCurrencyB3 int `json:"BossCurrencyB3,omitempty"`
		FinalBossCurrencyA int `json:"FinalBossCurrencyA,omitempty"`
		InstanceContentTextDataObjectiveEndTargetID int `json:"InstanceContentTextDataObjectiveEndTargetID,omitempty"`
		Order int `json:"Order,omitempty"`
		SortKey int `json:"SortKey,omitempty"`
		BNpcBaseBossTarget string `json:"BNpcBaseBossTarget,omitempty"`
		BossCurrencyB1 int `json:"BossCurrencyB1,omitempty"`
		ID int `json:"ID,omitempty"`
		InstanceContentTextDataBossStart int `json:"InstanceContentTextDataBossStart,omitempty"`
		TerritoryTypeTarget string `json:"TerritoryTypeTarget,omitempty"`
	}
	ItemLevelSync int `json:"ItemLevelSync,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	ID int `json:"ID,omitempty"`
	IconID int `json:"IconID,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
}