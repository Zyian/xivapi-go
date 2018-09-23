// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type GatheringPoint struct {
	GatheringPointBonus0Target string `json:"GatheringPointBonus0Target,omitempty"`
	GatheringSubCategoryTarget string `json:"GatheringSubCategoryTarget,omitempty"`
	TerritoryTypeTarget string `json:"TerritoryTypeTarget,omitempty"`
	GameContentLinks struct {
		GatheringItemPoint struct {
			GatheringPoint []interface{} `json:"GatheringPoint,omitempty"`
		}
	}
	GatheringPointBaseTargetID int `json:"GatheringPointBaseTargetID,omitempty"`
	GatheringPointBonus0 struct {
		BonusType struct {
			ID int `json:"ID,omitempty"`
			Text string `json:"Text,omitempty"`
			Text_de string `json:"Text_de,omitempty"`
			Text_en string `json:"Text_en,omitempty"`
			Text_fr string `json:"Text_fr,omitempty"`
			Text_ja string `json:"Text_ja,omitempty"`
		}
		ConditionValue int `json:"ConditionValue,omitempty"`
		ConditionTarget string `json:"ConditionTarget,omitempty"`
		ConditionTargetID int `json:"ConditionTargetID,omitempty"`
		ID int `json:"ID,omitempty"`
		BonusTypeTarget string `json:"BonusTypeTarget,omitempty"`
		BonusTypeTargetID int `json:"BonusTypeTargetID,omitempty"`
		BonusValue int `json:"BonusValue,omitempty"`
		Condition struct {
			Text_fr string `json:"Text_fr,omitempty"`
			Text_ja string `json:"Text_ja,omitempty"`
			ID int `json:"ID,omitempty"`
			Text string `json:"Text,omitempty"`
			Text_de string `json:"Text_de,omitempty"`
			Text_en string `json:"Text_en,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	TerritoryTypeTargetID int `json:"TerritoryTypeTargetID,omitempty"`
	GatheringPointBonus0TargetID int `json:"GatheringPointBonus0TargetID,omitempty"`
	GatheringPointBonus1 int `json:"GatheringPointBonus1,omitempty"`
	GatheringSubCategoryTargetID int `json:"GatheringSubCategoryTargetID,omitempty"`
	PlaceName int `json:"PlaceName,omitempty"`
	TerritoryType struct {
		PlaceNameRegionTarget string `json:"PlaceNameRegionTarget,omitempty"`
		PlaceNameRegionTargetID int `json:"PlaceNameRegionTargetID,omitempty"`
		PlaceNameTarget string `json:"PlaceNameTarget,omitempty"`
		PlaceNameZone struct {
			Name_ja string `json:"Name_ja,omitempty"`
			Icon string `json:"Icon,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
		}
		ArrayEventHandlerTarget string `json:"ArrayEventHandlerTarget,omitempty"`
		Bg string `json:"Bg,omitempty"`
		Bg_en string `json:"Bg_en,omitempty"`
		MapTarget string `json:"MapTarget,omitempty"`
		PlaceNameZoneTarget string `json:"PlaceNameZoneTarget,omitempty"`
		PlaceNameZoneTargetID int `json:"PlaceNameZoneTargetID,omitempty"`
		Aetheryte struct {
			Level1 int `json:"Level1,omitempty"`
			AetherstreamY int `json:"AetherstreamY,omitempty"`
			IsAetheryte int `json:"IsAetheryte,omitempty"`
			AetherstreamX int `json:"AetherstreamX,omitempty"`
			ID int `json:"ID,omitempty"`
			Level0 int `json:"Level0,omitempty"`
			Level2 int `json:"Level2,omitempty"`
			Level3 int `json:"Level3,omitempty"`
			Map int `json:"Map,omitempty"`
			AethernetGroup int `json:"AethernetGroup,omitempty"`
			AethernetName int `json:"AethernetName,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			Territory int `json:"Territory,omitempty"`
		}
		ArrayEventHandlerTargetID int `json:"ArrayEventHandlerTargetID,omitempty"`
		PlaceName struct {
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			ID int `json:"ID,omitempty"`
		}
		PlaceNameRegion struct {
			Name_ja string `json:"Name_ja,omitempty"`
			ID int `json:"ID,omitempty"`
			Icon string `json:"Icon,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
		}
		PlaceNameTargetID int `json:"PlaceNameTargetID,omitempty"`
		TerritoryIntendedUse int `json:"TerritoryIntendedUse,omitempty"`
		WeatherRate int `json:"WeatherRate,omitempty"`
		AetheryteTarget string `json:"AetheryteTarget,omitempty"`
		ID int `json:"ID,omitempty"`
		Map struct {
			MapFilename string `json:"MapFilename,omitempty"`
			MapFilenameId string `json:"MapFilenameId,omitempty"`
			OffsetX int `json:"OffsetX,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			DiscoveryArrayByte int `json:"DiscoveryArrayByte,omitempty"`
			Hierarchy int `json:"Hierarchy,omitempty"`
			MapMarkerRange int `json:"MapMarkerRange,omitempty"`
			PlaceNameSub int `json:"PlaceNameSub,omitempty"`
			DiscoveryIndex int `json:"DiscoveryIndex,omitempty"`
			ID int `json:"ID,omitempty"`
			SizeFactor int `json:"SizeFactor,omitempty"`
			TerritoryType int `json:"TerritoryType,omitempty"`
			OffsetY int `json:"OffsetY,omitempty"`
			PlaceNameRegion int `json:"PlaceNameRegion,omitempty"`
		}
		MapTargetID int `json:"MapTargetID,omitempty"`
		AetheryteTargetID int `json:"AetheryteTargetID,omitempty"`
		ArrayEventHandler struct {
			MULTIREF1 int `json:"MULTIREF1,omitempty"`
			MULTIREF10 int `json:"MULTIREF10,omitempty"`
			MULTIREF2 int `json:"MULTIREF2,omitempty"`
			MULTIREF15 int `json:"MULTIREF15,omitempty"`
			MULTIREF4 int `json:"MULTIREF4,omitempty"`
			MULTIREF5 int `json:"MULTIREF5,omitempty"`
			MULTIREF8 int `json:"MULTIREF8,omitempty"`
			MULTIREF6 int `json:"MULTIREF6,omitempty"`
			ID int `json:"ID,omitempty"`
			MULTIREF12 int `json:"MULTIREF12,omitempty"`
			MULTIREF13 int `json:"MULTIREF13,omitempty"`
			MULTIREF3 int `json:"MULTIREF3,omitempty"`
			MULTIREF9 int `json:"MULTIREF9,omitempty"`
			MULTIREF0 int `json:"MULTIREF0,omitempty"`
			MULTIREF11 int `json:"MULTIREF11,omitempty"`
			MULTIREF14 int `json:"MULTIREF14,omitempty"`
			MULTIREF7 int `json:"MULTIREF7,omitempty"`
		}
		Name string `json:"Name,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
	}
	Url string `json:"Url,omitempty"`
	GatheringPointBase struct {
		Item0 int `json:"Item0,omitempty"`
		Item1 int `json:"Item1,omitempty"`
		Item7 int `json:"Item7,omitempty"`
		GatheringTypeTargetID int `json:"GatheringTypeTargetID,omitempty"`
		IsLimited int `json:"IsLimited,omitempty"`
		Item3 int `json:"Item3,omitempty"`
		Item5 int `json:"Item5,omitempty"`
		Item6 int `json:"Item6,omitempty"`
		GatheringType struct {
			ID int `json:"ID,omitempty"`
			IconMain string `json:"IconMain,omitempty"`
			IconOffID int `json:"IconOffID,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			IconMainID int `json:"IconMainID,omitempty"`
			IconOff string `json:"IconOff,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
		}
		ID int `json:"ID,omitempty"`
		Item2 int `json:"Item2,omitempty"`
		Item4 int `json:"Item4,omitempty"`
		GatheringLevel int `json:"GatheringLevel,omitempty"`
		GatheringTypeTarget string `json:"GatheringTypeTarget,omitempty"`
	}
	GatheringPointBaseTarget string `json:"GatheringPointBaseTarget,omitempty"`
	GatheringSubCategory struct {
		ItemTargetID int `json:"ItemTargetID,omitempty"`
		FolkloreBook_en string `json:"FolkloreBook_en,omitempty"`
		FolkloreBook_fr string `json:"FolkloreBook_fr,omitempty"`
		ID int `json:"ID,omitempty"`
		Item struct {
			EquipSlotCategory int `json:"EquipSlotCategory,omitempty"`
			ItemSpecialBonusParam int `json:"ItemSpecialBonusParam,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			EquipRestriction int `json:"EquipRestriction,omitempty"`
			StartsWithVowel int `json:"StartsWithVowel,omitempty"`
			IsCrestWorthy int `json:"IsCrestWorthy,omitempty"`
			BaseParamSpecial4 int `json:"BaseParamSpecial4,omitempty"`
			Description_de string `json:"Description_de,omitempty"`
			IsUntradable int `json:"IsUntradable,omitempty"`
			ItemRepair int `json:"ItemRepair,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			BaseParam3 int `json:"BaseParam3,omitempty"`
			Block int `json:"Block,omitempty"`
			DelayMs int `json:"DelayMs,omitempty"`
			LevelItem int `json:"LevelItem,omitempty"`
			StackSize int `json:"StackSize,omitempty"`
			BaseParamValue1 int `json:"BaseParamValue1,omitempty"`
			BaseParamValueSpecial1 int `json:"BaseParamValueSpecial1,omitempty"`
			BaseParamValueSpecial2 int `json:"BaseParamValueSpecial2,omitempty"`
			BaseParamValueSpecial5 int `json:"BaseParamValueSpecial5,omitempty"`
			CanBeHq int `json:"CanBeHq,omitempty"`
			ItemSearchCategory int `json:"ItemSearchCategory,omitempty"`
			PriceLow int `json:"PriceLow,omitempty"`
			BaseParam0 int `json:"BaseParam0,omitempty"`
			IsGlamourous int `json:"IsGlamourous,omitempty"`
			MaterializeType int `json:"MaterializeType,omitempty"`
			Plural_de string `json:"Plural_de,omitempty"`
			Salvage int `json:"Salvage,omitempty"`
			BaseParamSpecial2 int `json:"BaseParamSpecial2,omitempty"`
			DamageMag int `json:"DamageMag,omitempty"`
			DefensePhys int `json:"DefensePhys,omitempty"`
			Description_fr string `json:"Description_fr,omitempty"`
			ModelSub string `json:"ModelSub,omitempty"`
			BaseParam4 int `json:"BaseParam4,omitempty"`
			DamagePhys int `json:"DamagePhys,omitempty"`
			Rarity int `json:"Rarity,omitempty"`
			BaseParam5 int `json:"BaseParam5,omitempty"`
			BaseParamValueSpecial0 int `json:"BaseParamValueSpecial0,omitempty"`
			CooldownS int `json:"CooldownS,omitempty"`
			FilterGroup int `json:"FilterGroup,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Singular_de string `json:"Singular_de,omitempty"`
			Singular_en string `json:"Singular_en,omitempty"`
			BaseParam2 int `json:"BaseParam2,omitempty"`
			BaseParamValue3 int `json:"BaseParamValue3,omitempty"`
			BaseParamValueSpecial3 int `json:"BaseParamValueSpecial3,omitempty"`
			BlockRate int `json:"BlockRate,omitempty"`
			ID int `json:"ID,omitempty"`
			IsAdvancedMeldingPermitted int `json:"IsAdvancedMeldingPermitted,omitempty"`
			ItemGlamour int `json:"ItemGlamour,omitempty"`
			ItemUICategory int `json:"ItemUICategory,omitempty"`
			BaseParamModifier int `json:"BaseParamModifier,omitempty"`
			Singular string `json:"Singular,omitempty"`
			ModelMain string `json:"ModelMain,omitempty"`
			BaseParamSpecial5 int `json:"BaseParamSpecial5,omitempty"`
			DefenseMag int `json:"DefenseMag,omitempty"`
			IsUnique int `json:"IsUnique,omitempty"`
			ItemAction int `json:"ItemAction,omitempty"`
			PriceMid int `json:"PriceMid,omitempty"`
			Singular_fr string `json:"Singular_fr,omitempty"`
			BaseParamSpecial0 int `json:"BaseParamSpecial0,omitempty"`
			BaseParamSpecial1 int `json:"BaseParamSpecial1,omitempty"`
			BaseParamValueSpecial4 int `json:"BaseParamValueSpecial4,omitempty"`
			IsCollectable int `json:"IsCollectable,omitempty"`
			IsPvP int `json:"IsPvP,omitempty"`
			BaseParam1 int `json:"BaseParam1,omitempty"`
			BaseParamValue0 int `json:"BaseParamValue0,omitempty"`
			BaseParamValue2 int `json:"BaseParamValue2,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			ClassJobRepair int `json:"ClassJobRepair,omitempty"`
			IsDyeable int `json:"IsDyeable,omitempty"`
			LevelEquip int `json:"LevelEquip,omitempty"`
			AetherialReduce int `json:"AetherialReduce,omitempty"`
			ItemSeries int `json:"ItemSeries,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			Description string `json:"Description,omitempty"`
			BaseParamValue5 int `json:"BaseParamValue5,omitempty"`
			ClassJobUse int `json:"ClassJobUse,omitempty"`
			Description_ja string `json:"Description_ja,omitempty"`
			IconID int `json:"IconID,omitempty"`
			IsIndisposable int `json:"IsIndisposable,omitempty"`
			ItemSpecialBonus int `json:"ItemSpecialBonus,omitempty"`
			BaseParamValue4 int `json:"BaseParamValue4,omitempty"`
			Description_en string `json:"Description_en,omitempty"`
			GrandCompany int `json:"GrandCompany,omitempty"`
			MateriaSlotCount int `json:"MateriaSlotCount,omitempty"`
			Name string `json:"Name,omitempty"`
			Plural string `json:"Plural,omitempty"`
			Plural_en string `json:"Plural_en,omitempty"`
			Plural_fr string `json:"Plural_fr,omitempty"`
			BaseParamSpecial3 int `json:"BaseParamSpecial3,omitempty"`
			Singular_ja string `json:"Singular_ja,omitempty"`
			Stain int `json:"Stain,omitempty"`
			Plural_ja string `json:"Plural_ja,omitempty"`
		}
		ItemTarget string `json:"ItemTarget,omitempty"`
		FolkloreBook string `json:"FolkloreBook,omitempty"`
		FolkloreBook_de string `json:"FolkloreBook_de,omitempty"`
		FolkloreBook_ja string `json:"FolkloreBook_ja,omitempty"`
	}
}
