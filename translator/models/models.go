package models

type SRDAbilityScores struct {
	Desc      []string `json:"desc"`
	Full_name string   `json:"full_name"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Skills    []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"skills"`
	Url string `json:"url"`
}

type SRDAlignments struct {
	Abbreviation string `json:"abbreviation"`
	Desc         string `json:"desc"`
	Index        string `json:"index"`
	Name         string `json:"name"`
	Url          string `json:"url"`
}

type SRDBackgrounds struct {
	Bonds struct {
		Choose int `json:"choose"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Option_type string `json:"option_type"`
				String      string `json:"string"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"bonds"`
	Feature struct {
		Desc []string `json:"desc"`
		Name string   `json:"name"`
	} `json:"feature"`
	Flaws struct {
		Choose int `json:"choose"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Option_type string `json:"option_type"`
				String      string `json:"string"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"flaws"`
	Ideals struct {
		Choose int `json:"choose"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Alignments []struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"alignments"`
				Desc        string `json:"desc"`
				Option_type string `json:"option_type"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"ideals"`
	Index            string `json:"index"`
	Language_options struct {
		Choose int `json:"choose"`
		From   struct {
			Option_set_type   string `json:"option_set_type"`
			Resource_list_url string `json:"resource_list_url"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"language_options"`
	Name               string `json:"name"`
	Personality_traits struct {
		Choose int `json:"choose"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Option_type string `json:"option_type"`
				String      string `json:"string"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"personality_traits"`
	Starting_equipment []struct {
		Equipment struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"equipment"`
		Quantity int `json:"quantity"`
	} `json:"starting_equipment"`
	Starting_equipment_options []struct {
		Choose int `json:"choose"`
		From   struct {
			Equipment_category struct {
				Index string `json:"index"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"equipment_category"`
			Option_set_type string `json:"option_set_type"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"starting_equipment_options"`
	Starting_proficiencies []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"starting_proficiencies"`
	Url string `json:"url"`
}

type SRDClasses struct {
	Class_levels   string `json:"class_levels"`
	Hit_die        int    `json:"hit_die"`
	Index          string `json:"index"`
	Multi_classing struct {
		Prerequisite_options struct {
			Choose int `json:"choose"`
			From   struct {
				Option_set_type string `json:"option_set_type"`
				Options         []struct {
					Ability_score struct {
						Index string `json:"index"`
						Name  string `json:"name"`
						Url   string `json:"url"`
					} `json:"ability_score"`
					Minimum_score int    `json:"minimum_score"`
					Option_type   string `json:"option_type"`
				} `json:"options"`
			} `json:"from"`
			Type string `json:"type"`
		} `json:"prerequisite_options"`
		Prerequisites []struct {
			Ability_score struct {
				Index string `json:"index"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"ability_score"`
			Minimum_score int `json:"minimum_score"`
		} `json:"prerequisites"`
		Proficiencies []struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"proficiencies"`
		Proficiency_choices []struct {
			Choose int    `json:"choose"`
			Desc   string `json:"desc"`
			From   struct {
				Option_set_type string `json:"option_set_type"`
				Options         []struct {
					Item struct {
						Index string `json:"index"`
						Name  string `json:"name"`
						Url   string `json:"url"`
					} `json:"item"`
					Option_type string `json:"option_type"`
				} `json:"options"`
			} `json:"from"`
			Type string `json:"type"`
		} `json:"proficiency_choices"`
	} `json:"multi_classing"`
	Name          string `json:"name"`
	Proficiencies []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"proficiencies"`
	Proficiency_choices []struct {
		Choose int    `json:"choose"`
		Desc   string `json:"desc"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Item struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"item"`
				Option_type string `json:"option_type"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"proficiency_choices"`
	Saving_throws []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"saving_throws"`
	Spellcasting struct {
		Info []struct {
			Desc []string `json:"desc"`
			Name string   `json:"name"`
		} `json:"info"`
		Level                int `json:"level"`
		Spellcasting_ability struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"spellcasting_ability"`
	} `json:"spellcasting"`
	Spells             string `json:"spells"`
	Starting_equipment []struct {
		Equipment struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"equipment"`
		Quantity int `json:"quantity"`
	} `json:"starting_equipment"`
	Starting_equipment_options []struct {
		Choose int    `json:"choose"`
		Desc   string `json:"desc"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Choice struct {
					Choose int    `json:"choose"`
					Desc   string `json:"desc"`
					From   struct {
						Equipment_category struct {
							Index string `json:"index"`
							Name  string `json:"name"`
							Url   string `json:"url"`
						} `json:"equipment_category"`
						Option_set_type string `json:"option_set_type"`
					} `json:"from"`
					Type string `json:"type"`
				} `json:"choice"`
				Count int `json:"count"`
				Of    struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"of"`
				Option_type string `json:"option_type"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"starting_equipment_options"`
	Subclasses []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"subclasses"`
	Url string `json:"url"`
}

type SRDConditions struct {
	Desc  []string `json:"desc"`
	Index string   `json:"index"`
	Name  string   `json:"name"`
	Url   string   `json:"url"`
}

type SRDDamageTypes struct {
	Desc  []string `json:"desc"`
	Index string   `json:"index"`
	Name  string   `json:"name"`
	Url   string   `json:"url"`
}

type SRDEquipmentCategories struct {
	Equipment []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"equipment"`
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

type SRDEquipment struct {
	Armor_category string `json:"armor_category"`
	Armor_class    struct {
		Base      int  `json:"base"`
		Dex_bonus bool `json:"dex_bonus"`
		Max_bonus int  `json:"max_bonus"`
	} `json:"armor_class"`
	Capacity       string `json:"capacity"`
	Category_range string `json:"category_range"`
	Contents       []struct {
		Item struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"item"`
		Quantity int `json:"quantity"`
	} `json:"contents"`
	Cost struct {
		Quantity int    `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"cost"`
	Damage struct {
		Damage_dice string `json:"damage_dice"`
		Damage_type struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"damage_type"`
	} `json:"damage"`
	Desc               []string `json:"desc"`
	Equipment_category struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"equipment_category"`
	Gear_category struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"gear_category"`
	Index      string `json:"index"`
	Name       string `json:"name"`
	Properties []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"properties"`
	Quantity int `json:"quantity"`
	Range    struct {
		Long   int `json:"long"`
		Normal int `json:"normal"`
	} `json:"range"`
	Special []string `json:"special"`
	Speed   struct {
		Quantity float64 `json:"quantity"`
		Unit     string  `json:"unit"`
	} `json:"speed"`
	Stealth_disadvantage bool `json:"stealth_disadvantage"`
	Str_minimum          int  `json:"str_minimum"`
	Throw_range          struct {
		Long   int `json:"long"`
		Normal int `json:"normal"`
	} `json:"throw_range"`
	Tool_category     string `json:"tool_category"`
	Two_handed_damage struct {
		Damage_dice string `json:"damage_dice"`
		Damage_type struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"damage_type"`
	} `json:"two_handed_damage"`
	Url              string  `json:"url"`
	Vehicle_category string  `json:"vehicle_category"`
	Weapon_category  string  `json:"weapon_category"`
	Weapon_range     string  `json:"weapon_range"`
	Weight           float64 `json:"weight"`
}

type SRDFeats struct {
	Desc          []string `json:"desc"`
	Index         string   `json:"index"`
	Name          string   `json:"name"`
	Prerequisites []struct {
		Ability_score struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"ability_score"`
		Minimum_score int `json:"minimum_score"`
	} `json:"prerequisites"`
	Url string `json:"url"`
}

type SRDFeatures struct {
	Class struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"class"`
	Desc             []string `json:"desc"`
	Feature_specific struct {
		Expertise_options struct {
			Choose int `json:"choose"`
			From   struct {
				Option_set_type string `json:"option_set_type"`
				Options         []struct {
					Item struct {
						Index string `json:"index"`
						Name  string `json:"name"`
						Url   string `json:"url"`
					} `json:"item"`
					Option_type string `json:"option_type"`
				} `json:"options"`
			} `json:"from"`
			Type string `json:"type"`
		} `json:"expertise_options"`
		Invocations []struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"invocations"`
		Subfeature_options struct {
			Choose int `json:"choose"`
			From   struct {
				Option_set_type string `json:"option_set_type"`
				Options         []struct {
					Item struct {
						Index string `json:"index"`
						Name  string `json:"name"`
						Url   string `json:"url"`
					} `json:"item"`
					Option_type string `json:"option_type"`
				} `json:"options"`
			} `json:"from"`
			Type string `json:"type"`
		} `json:"subfeature_options"`
	} `json:"feature_specific"`
	Index  string `json:"index"`
	Level  int    `json:"level"`
	Name   string `json:"name"`
	Parent struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"parent"`
	Prerequisites []interface{} `json:"prerequisites"`
	Reference     string        `json:"reference"`
	Subclass      struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"subclass"`
	Url string `json:"url"`
}

type SRDLanguages struct {
	Desc             string   `json:"desc"`
	Index            string   `json:"index"`
	Name             string   `json:"name"`
	Script           string   `json:"script"`
	Type             string   `json:"type"`
	Typical_speakers []string `json:"typical_speakers"`
	Url              string   `json:"url"`
}

type SRDLevels struct {
	Ability_score_bonuses int `json:"ability_score_bonuses"`
	Class                 struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"class"`
	Class_specific struct {
		Action_surges            int           `json:"action_surges"`
		Arcane_recovery_levels   int           `json:"arcane_recovery_levels"`
		Aura_range               int           `json:"aura_range"`
		Bardic_inspiration_die   int           `json:"bardic_inspiration_die"`
		Brutal_critical_dice     int           `json:"brutal_critical_dice"`
		Channel_divinity_charges int           `json:"channel_divinity_charges"`
		Creating_spell_slots     []interface{} `json:"creating_spell_slots"`
		Destroy_undead_cr        float64       `json:"destroy_undead_cr"`
		Extra_attacks            int           `json:"extra_attacks"`
		Favored_enemies          int           `json:"favored_enemies"`
		Favored_terrain          int           `json:"favored_terrain"`
		Indomitable_uses         int           `json:"indomitable_uses"`
		Invocations_known        int           `json:"invocations_known"`
		Ki_points                int           `json:"ki_points"`
		Magical_secrets_max_5    int           `json:"magical_secrets_max_5"`
		Magical_secrets_max_7    int           `json:"magical_secrets_max_7"`
		Magical_secrets_max_9    int           `json:"magical_secrets_max_9"`
		Martial_arts             struct {
			Dice_count int `json:"dice_count"`
			Dice_value int `json:"dice_value"`
		} `json:"martial_arts"`
		Metamagic_known        int `json:"metamagic_known"`
		Mystic_arcanum_level_6 int `json:"mystic_arcanum_level_6"`
		Mystic_arcanum_level_7 int `json:"mystic_arcanum_level_7"`
		Mystic_arcanum_level_8 int `json:"mystic_arcanum_level_8"`
		Mystic_arcanum_level_9 int `json:"mystic_arcanum_level_9"`
		Rage_count             int `json:"rage_count"`
		Rage_damage_bonus      int `json:"rage_damage_bonus"`
		Sneak_attack           struct {
			Dice_count int `json:"dice_count"`
			Dice_value int `json:"dice_value"`
		} `json:"sneak_attack"`
		Song_of_rest_die   int     `json:"song_of_rest_die"`
		Sorcery_points     int     `json:"sorcery_points"`
		Unarmored_movement int     `json:"unarmored_movement"`
		Wild_shape_fly     bool    `json:"wild_shape_fly"`
		Wild_shape_max_cr  float64 `json:"wild_shape_max_cr"`
		Wild_shape_swim    bool    `json:"wild_shape_swim"`
	} `json:"class_specific"`
	Features []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"features"`
	Index        string `json:"index"`
	Level        int    `json:"level"`
	Prof_bonus   int    `json:"prof_bonus"`
	Spellcasting struct {
		Cantrips_known      int `json:"cantrips_known"`
		Spell_slots_level_1 int `json:"spell_slots_level_1"`
		Spell_slots_level_2 int `json:"spell_slots_level_2"`
		Spell_slots_level_3 int `json:"spell_slots_level_3"`
		Spell_slots_level_4 int `json:"spell_slots_level_4"`
		Spell_slots_level_5 int `json:"spell_slots_level_5"`
		Spell_slots_level_6 int `json:"spell_slots_level_6"`
		Spell_slots_level_7 int `json:"spell_slots_level_7"`
		Spell_slots_level_8 int `json:"spell_slots_level_8"`
		Spell_slots_level_9 int `json:"spell_slots_level_9"`
		Spells_known        int `json:"spells_known"`
	} `json:"spellcasting"`
	Subclass struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"subclass"`
	Subclass_specific struct {
		Additional_magical_secrets_max_lvl int `json:"additional_magical_secrets_max_lvl"`
		Aura_range                         int `json:"aura_range"`
	} `json:"subclass_specific"`
	Url string `json:"url"`
}

type SRDMagicItems struct {
	Desc               []string `json:"desc"`
	Equipment_category struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"equipment_category"`
	Index  string `json:"index"`
	Name   string `json:"name"`
	Rarity struct {
		Name string `json:"name"`
	} `json:"rarity"`
	Url      string `json:"url"`
	Variant  bool   `json:"variant"`
	Variants []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"variants"`
}

type SRDMagicSchools struct {
	Desc  string `json:"desc"`
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

type SRDMonsters struct {
	Actions []struct {
		Actions []struct {
			Action_name string `json:"action_name"`
			Count       int    `json:"count"`
			Type        string `json:"type"`
		} `json:"actions"`
		Attack_bonus int `json:"attack_bonus"`
		Damage       []struct {
			Damage_dice string `json:"damage_dice"`
			Damage_type struct {
				Index string `json:"index"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"damage_type"`
		} `json:"damage"`
		Dc struct {
			Dc_type struct {
				Index string `json:"index"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"dc_type"`
			Dc_value     int    `json:"dc_value"`
			Success_type string `json:"success_type"`
		} `json:"dc"`
		Desc             string `json:"desc"`
		Multiattack_type string `json:"multiattack_type"`
		Name             string `json:"name"`
		Usage            struct {
			Times int    `json:"times"`
			Type  string `json:"type"`
		} `json:"usage"`
	} `json:"actions"`
	Alignment   string `json:"alignment"`
	Armor_class []struct {
		Type  string `json:"type"`
		Value int    `json:"value"`
	} `json:"armor_class"`
	Challenge_rating     float64 `json:"challenge_rating"`
	Charisma             int     `json:"charisma"`
	Condition_immunities []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"condition_immunities"`
	Constitution           int      `json:"constitution"`
	Damage_immunities      []string `json:"damage_immunities"`
	Damage_resistances     []string `json:"damage_resistances"`
	Damage_vulnerabilities []string `json:"damage_vulnerabilities"`
	Desc                   string   `json:"desc"`
	Dexterity              int      `json:"dexterity"`
	Forms                  []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"forms"`
	Hit_dice          string `json:"hit_dice"`
	Hit_points        int    `json:"hit_points"`
	Hit_points_roll   string `json:"hit_points_roll"`
	Image             string `json:"image"`
	Index             string `json:"index"`
	Intelligence      int    `json:"intelligence"`
	Languages         string `json:"languages"`
	Legendary_actions []struct {
		Attack_bonus int `json:"attack_bonus"`
		Damage       []struct {
			Damage_dice string `json:"damage_dice"`
			Damage_type struct {
				Index string `json:"index"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"damage_type"`
		} `json:"damage"`
		Desc string `json:"desc"`
		Name string `json:"name"`
	} `json:"legendary_actions"`
	Name          string `json:"name"`
	Proficiencies []struct {
		Proficiency struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"proficiency"`
		Value int `json:"value"`
	} `json:"proficiencies"`
	Proficiency_bonus int `json:"proficiency_bonus"`
	Reactions         []struct {
		Desc string `json:"desc"`
		Name string `json:"name"`
	} `json:"reactions"`
	Senses struct {
		Blindsight         string  `json:"blindsight"`
		Darkvision         string  `json:"darkvision"`
		Passive_perception float64 `json:"passive_perception"`
		Tremorsense        string  `json:"tremorsense"`
		Truesight          string  `json:"truesight"`
	} `json:"senses"`
	Size              string `json:"size"`
	Special_abilities []struct {
		Dc struct {
			Dc_type struct {
				Index string `json:"index"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"dc_type"`
			Dc_value     int    `json:"dc_value"`
			Success_type string `json:"success_type"`
		} `json:"dc"`
		Desc string `json:"desc"`
		Name string `json:"name"`
	} `json:"special_abilities"`
	Speed struct {
		Burrow string `json:"burrow"`
		Climb  string `json:"climb"`
		Fly    string `json:"fly"`
		Hover  bool   `json:"hover"`
		Swim   string `json:"swim"`
		Walk   string `json:"walk"`
	} `json:"speed"`
	Strength int    `json:"strength"`
	Subtype  string `json:"subtype"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Wisdom   int    `json:"wisdom"`
	Xp       int    `json:"xp"`
}

type SRDProficiencies struct {
	Classes []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"classes"`
	Index string `json:"index"`
	Name  string `json:"name"`
	Races []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"races"`
	Reference struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"reference"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

type SRDRaces struct {
	Ability_bonus_options struct {
		Choose int `json:"choose"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Ability_score struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"ability_score"`
				Bonus       int    `json:"bonus"`
				Option_type string `json:"option_type"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"ability_bonus_options"`
	Ability_bonuses []struct {
		Ability_score struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"ability_score"`
		Bonus int `json:"bonus"`
	} `json:"ability_bonuses"`
	Age              string `json:"age"`
	Alignment        string `json:"alignment"`
	Index            string `json:"index"`
	Language_desc    string `json:"language_desc"`
	Language_options struct {
		Choose int `json:"choose"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Item struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"item"`
				Option_type string `json:"option_type"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"language_options"`
	Languages []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"languages"`
	Name                   string `json:"name"`
	Size                   string `json:"size"`
	Size_description       string `json:"size_description"`
	Speed                  int    `json:"speed"`
	Starting_proficiencies []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"starting_proficiencies"`
	Starting_proficiency_options struct {
		Choose int    `json:"choose"`
		Desc   string `json:"desc"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Item struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"item"`
				Option_type string `json:"option_type"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"starting_proficiency_options"`
	Subraces []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"subraces"`
	Traits []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"traits"`
	Url string `json:"url"`
}

type SRDRuleSections struct {
	Desc  string `json:"desc"`
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

type SRDRules struct {
	Desc        string `json:"desc"`
	Index       string `json:"index"`
	Name        string `json:"name"`
	Subsections []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"subsections"`
	Url string `json:"url"`
}

type SRDSkills struct {
	Ability_score struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"ability_score"`
	Desc  []string `json:"desc"`
	Index string   `json:"index"`
	Name  string   `json:"name"`
	Url   string   `json:"url"`
}

type SRDSpells struct {
	Area_of_effect struct {
		Size int    `json:"size"`
		Type string `json:"type"`
	} `json:"area_of_effect"`
	Attack_type  string `json:"attack_type"`
	Casting_time string `json:"casting_time"`
	Classes      []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"classes"`
	Components    []string `json:"components"`
	Concentration bool     `json:"concentration"`
	Damage        struct {
		Damage_at_character_level map[string]string `json:"damage_at_character_level"`
		Damage_at_slot_level      map[string]string `json:"damage_at_slot_level"`
		Damage_type               struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"damage_type"`
	} `json:"damage"`
	Dc struct {
		Dc_success string `json:"dc_success"`
		Dc_type    struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"dc_type"`
		Desc string `json:"desc"`
	} `json:"dc"`
	Desc               []string          `json:"desc"`
	Duration           string            `json:"duration"`
	Heal_at_slot_level map[string]string `json:"heal_at_slot_level"`
	Higher_level       []string          `json:"higher_level"`
	Index              string            `json:"index"`
	Level              int               `json:"level"`
	Material           string            `json:"material"`
	Name               string            `json:"name"`
	Range              string            `json:"range"`
	Ritual             bool              `json:"ritual"`
	School             struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"school"`
	Subclasses []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"subclasses"`
	Url string `json:"url"`
}

type SRDSubclasses struct {
	Class struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"class"`
	Desc   []string `json:"desc"`
	Index  string   `json:"index"`
	Name   string   `json:"name"`
	Spells []struct {
		Prerequisites []struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Type  string `json:"type"`
			Url   string `json:"url"`
		} `json:"prerequisites"`
		Spell struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"spell"`
	} `json:"spells"`
	Subclass_flavor string `json:"subclass_flavor"`
	Subclass_levels string `json:"subclass_levels"`
	Url             string `json:"url"`
}

type SRDSubraces struct {
	Ability_bonuses []struct {
		Ability_score struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"ability_score"`
		Bonus int `json:"bonus"`
	} `json:"ability_bonuses"`
	Desc             string `json:"desc"`
	Index            string `json:"index"`
	Language_options struct {
		Choose int `json:"choose"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Item struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"item"`
				Option_type string `json:"option_type"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"language_options"`
	Languages []interface{} `json:"languages"`
	Name      string        `json:"name"`
	Race      struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"race"`
	Racial_traits []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"racial_traits"`
	Starting_proficiencies []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"starting_proficiencies"`
	Url string `json:"url"`
}

type SRDTraits struct {
	Desc             []string `json:"desc"`
	Index            string   `json:"index"`
	Language_options struct {
		Choose int `json:"choose"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Item struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"item"`
				Option_type string `json:"option_type"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"language_options"`
	Name   string `json:"name"`
	Parent struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"parent"`
	Proficiencies []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"proficiencies"`
	Proficiency_choices struct {
		Choose int `json:"choose"`
		From   struct {
			Option_set_type string `json:"option_set_type"`
			Options         []struct {
				Item struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"item"`
				Option_type string `json:"option_type"`
			} `json:"options"`
		} `json:"from"`
		Type string `json:"type"`
	} `json:"proficiency_choices"`
	Races []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"races"`
	Subraces []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"subraces"`
	Trait_specific struct {
		Breath_weapon struct {
			Area_of_effect struct {
				Size int    `json:"size"`
				Type string `json:"type"`
			} `json:"area_of_effect"`
			Damage []struct {
				Damage_at_character_level map[string]string `json:"damage_at_character_level"`
				Damage_type               struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"damage_type"`
			} `json:"damage"`
			Dc struct {
				Dc_type struct {
					Index string `json:"index"`
					Name  string `json:"name"`
					Url   string `json:"url"`
				} `json:"dc_type"`
				Success_type string `json:"success_type"`
			} `json:"dc"`
			Desc  string `json:"desc"`
			Name  string `json:"name"`
			Usage struct {
				Times int    `json:"times"`
				Type  string `json:"type"`
			} `json:"usage"`
		} `json:"breath_weapon"`
		Damage_type struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"damage_type"`
		Spell_options struct {
			Choose int `json:"choose"`
			From   struct {
				Option_set_type string `json:"option_set_type"`
				Options         []struct {
					Item struct {
						Index string `json:"index"`
						Name  string `json:"name"`
						Url   string `json:"url"`
					} `json:"item"`
					Option_type string `json:"option_type"`
				} `json:"options"`
			} `json:"from"`
			Type string `json:"type"`
		} `json:"spell_options"`
		Subtrait_options struct {
			Choose int `json:"choose"`
			From   struct {
				Option_set_type string `json:"option_set_type"`
				Options         []struct {
					Item struct {
						Index string `json:"index"`
						Name  string `json:"name"`
						Url   string `json:"url"`
					} `json:"item"`
					Option_type string `json:"option_type"`
				} `json:"options"`
			} `json:"from"`
			Type string `json:"type"`
		} `json:"subtrait_options"`
	} `json:"trait_specific"`
	Url string `json:"url"`
}

type SRDWeaponProperties struct {
	Desc  []string `json:"desc"`
	Index string   `json:"index"`
	Name  string   `json:"name"`
	Url   string   `json:"url"`
}
