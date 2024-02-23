package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"translator/models"

	"github.com/joho/godotenv"
)

type translateFunc func([]string) []string

type modelSelectorFunc[T any] func(v []T, f translateFunc)

type handler[T any] struct {
	selector modelSelectorFunc[T]
}

func (d *handler[T]) handle(data []byte, translator *translator) ([]byte, error) {
	v := make([]T, 0)
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}

	d.selector(v, translator.Translate)
	return json.MarshalIndent(v, "", "  ")
}

type modelDescriptor interface {
	handle([]byte, *translator) ([]byte, error)
}

func singleTranslateWrap(f translateFunc, v string) string {
	return f([]string{v})[0]
}

var orm = map[string]modelDescriptor{
	"5e-SRD-Ability-Scores.json": &handler[*models.SRDAbilityScores]{
		selector: func(v []*models.SRDAbilityScores, f translateFunc) {
			for i := range v {
				v[i].Desc = f(v[i].Desc)
				v[i].Full_name = singleTranslateWrap(f, v[i].Full_name)
				for j := range v[i].Skills {
					v[i].Skills[j].Name = singleTranslateWrap(f, v[i].Skills[j].Name)
				}
			}
		},
	},
	"5e-SRD-Alignments.json": &handler[*models.SRDAlignments]{
		selector: func(v []*models.SRDAlignments, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = singleTranslateWrap(f, v[i].Desc)
			}
		},
	},
	"5e-SRD-Backgrounds.json": &handler[*models.SRDBackgrounds]{
		selector: func(v []*models.SRDBackgrounds, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Feature.Name = singleTranslateWrap(f,
					v[i].Feature.Name)
				v[i].Feature.Desc = f(v[i].Feature.Desc)
				for j := range v[i].Starting_equipment {
					v[i].Starting_equipment[j].Equipment.Name = singleTranslateWrap(f,
						v[i].Starting_equipment[j].Equipment.Name)
				}
				for j := range v[i].Starting_equipment_options {
					v[i].Starting_equipment_options[j].From.Equipment_category.Name = singleTranslateWrap(f,
						v[i].Starting_equipment_options[j].From.Equipment_category.Name)
				}
				for j := range v[i].Personality_traits.From.Options {
					v[i].Personality_traits.From.Options[j].String = singleTranslateWrap(f,
						v[i].Personality_traits.From.Options[j].String)
				}
				for j := range v[i].Bonds.From.Options {
					v[i].Bonds.From.Options[j].String = singleTranslateWrap(f,
						v[i].Bonds.From.Options[j].String)
				}
				for j := range v[i].Flaws.From.Options {
					v[i].Flaws.From.Options[j].String = singleTranslateWrap(f,
						v[i].Flaws.From.Options[j].String)
				}
				for j := range v[i].Ideals.From.Options {
					v[i].Ideals.From.Options[j].Desc = singleTranslateWrap(f,
						v[i].Ideals.From.Options[j].Desc)
					for k := range v[i].Ideals.From.Options[j].Alignments {
						v[i].Ideals.From.Options[j].Alignments[k].Name = singleTranslateWrap(f,
							v[i].Ideals.From.Options[j].Alignments[k].Name)
					}
				}
				for j := range v[i].Starting_proficiencies {
					v[i].Starting_proficiencies[j].Name = singleTranslateWrap(f,
						v[i].Starting_proficiencies[j].Name)
				}
			}
		},
	},
	"5e-SRD-Classes.json": &handler[*models.SRDClasses]{
		selector: func(v []*models.SRDClasses, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				for j := range v[i].Multi_classing.Prerequisite_options.From.Options {
					v[i].Multi_classing.Prerequisite_options.From.Options[j].Ability_score.Name = singleTranslateWrap(f,
						v[i].Multi_classing.Prerequisite_options.From.Options[j].Ability_score.Name)
				}
				for j := range v[i].Multi_classing.Prerequisites {
					v[i].Multi_classing.Prerequisites[j].Ability_score.Name = singleTranslateWrap(f,
						v[i].Multi_classing.Prerequisites[j].Ability_score.Name)
				}
				for j := range v[i].Multi_classing.Proficiencies {
					v[i].Multi_classing.Proficiencies[j].Name = singleTranslateWrap(f,
						v[i].Multi_classing.Proficiencies[j].Name)
				}
				for j := range v[i].Multi_classing.Proficiency_choices {
					v[i].Multi_classing.Proficiency_choices[j].Desc = singleTranslateWrap(f,
						v[i].Multi_classing.Proficiency_choices[j].Desc)
					for k := range v[i].Multi_classing.Proficiency_choices[j].From.Options {
						v[i].Multi_classing.Proficiency_choices[j].From.Options[k].Item.Name = singleTranslateWrap(f,
							v[i].Multi_classing.Proficiency_choices[j].From.Options[k].Item.Name)
					}
				}
				for j := range v[i].Proficiencies {
					v[i].Proficiencies[j].Name = singleTranslateWrap(f,
						v[i].Proficiencies[j].Name)
				}
				for j := range v[i].Proficiency_choices {
					v[i].Proficiency_choices[j].Desc = singleTranslateWrap(f,
						v[i].Proficiency_choices[j].Desc)
					for k := range v[i].Proficiency_choices[j].From.Options {
						v[i].Proficiency_choices[j].From.Options[k].Item.Name = singleTranslateWrap(f,
							v[i].Proficiency_choices[j].From.Options[k].Item.Name)
					}
				}
				for j := range v[i].Saving_throws {
					v[i].Saving_throws[j].Name = singleTranslateWrap(f,
						v[i].Saving_throws[j].Name)
				}
				for j := range v[i].Spellcasting.Info {
					v[i].Spellcasting.Info[j].Desc = f(v[i].Spellcasting.Info[j].Desc)
					v[i].Spellcasting.Info[j].Name = singleTranslateWrap(f,
						v[i].Spellcasting.Info[j].Name)
				}
				v[i].Spellcasting.Spellcasting_ability.Name = singleTranslateWrap(f,
					v[i].Spellcasting.Spellcasting_ability.Name)
				for j := range v[i].Starting_equipment {
					v[i].Starting_equipment[j].Equipment.Name = singleTranslateWrap(f,
						v[i].Starting_equipment[j].Equipment.Name)
				}
				for j := range v[i].Starting_equipment_options {
					v[i].Starting_equipment_options[j].Desc = singleTranslateWrap(f,
						v[i].Starting_equipment_options[j].Desc)
					for k := range v[i].Starting_equipment_options[j].From.Options {
						v[i].Starting_equipment_options[j].From.Options[k].Choice.Desc = singleTranslateWrap(f,
							v[i].Starting_equipment_options[j].From.Options[k].Choice.Desc)
						v[i].Starting_equipment_options[j].From.Options[k].Choice.From.Equipment_category.Name = singleTranslateWrap(f,
							v[i].Starting_equipment_options[j].From.Options[k].Choice.From.Equipment_category.Name)
						v[i].Starting_equipment_options[j].From.Options[k].Of.Name = singleTranslateWrap(f,
							v[i].Starting_equipment_options[j].From.Options[k].Of.Name)
					}
				}
				for j := range v[i].Subclasses {
					v[i].Subclasses[j].Name = singleTranslateWrap(f,
						v[i].Subclasses[j].Name)
				}
			}
		},
	},
	"5e-SRD-Conditions.json": &handler[*models.SRDConditions]{
		selector: func(v []*models.SRDConditions, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = f(v[i].Desc)
			}
		},
	},
	"5e-SRD-Damage-Types.json": &handler[*models.SRDDamageTypes]{
		selector: func(v []*models.SRDDamageTypes, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = f(v[i].Desc)
			}
		},
	},
	"5e-SRD-Equipment-Categories.json": &handler[*models.SRDEquipmentCategories]{
		selector: func(v []*models.SRDEquipmentCategories, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				for j := range v[i].Equipment {
					v[i].Equipment[j].Name = singleTranslateWrap(f, v[i].Equipment[j].Name)
				}
			}
		},
	},
	"5e-SRD-Equipment.json": &handler[*models.SRDEquipment]{
		selector: func(v []*models.SRDEquipment, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Armor_category = singleTranslateWrap(f, v[i].Armor_category)
				v[i].Capacity = singleTranslateWrap(f, v[i].Capacity)
				v[i].Category_range = singleTranslateWrap(f, v[i].Category_range)
				for j := range v[i].Contents {
					v[i].Contents[j].Item.Name = singleTranslateWrap(f,
						v[i].Contents[j].Item.Name)
				}
				v[i].Damage.Damage_type.Name = singleTranslateWrap(f,
					v[i].Damage.Damage_type.Name)
				v[i].Desc = f(v[i].Desc)
				v[i].Equipment_category.Name = singleTranslateWrap(f,
					v[i].Equipment_category.Name)
				v[i].Gear_category.Name = singleTranslateWrap(f,
					v[i].Gear_category.Name)
				for j := range v[i].Properties {
					v[i].Properties[j].Name = singleTranslateWrap(f,
						v[i].Properties[j].Name)
				}
				v[i].Special = f(v[i].Special)
				v[i].Tool_category = singleTranslateWrap(f, v[i].Tool_category)
				v[i].Two_handed_damage.Damage_type.Name = singleTranslateWrap(f, v[i].Two_handed_damage.Damage_type.Name)
				v[i].Vehicle_category = singleTranslateWrap(f, v[i].Vehicle_category)
				v[i].Weapon_category = singleTranslateWrap(f, v[i].Weapon_category)
				v[i].Weapon_range = singleTranslateWrap(f, v[i].Weapon_range)
			}
		},
	},
	"5e-SRD-Feats.json": &handler[*models.SRDFeats]{
		selector: func(v []*models.SRDFeats, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = f(v[i].Desc)
				for j := range v[i].Prerequisites {
					v[i].Prerequisites[j].Ability_score.Name = singleTranslateWrap(f,
						v[i].Prerequisites[j].Ability_score.Name)
				}
			}
		},
	},
	"5e-SRD-Features.json": &handler[*models.SRDFeatures]{
		selector: func(v []*models.SRDFeatures, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = f(v[i].Desc)
				v[i].Class.Name = singleTranslateWrap(f, v[i].Class.Name)

				for j := range v[i].Feature_specific.Expertise_options.From.Options {
					v[i].Feature_specific.Expertise_options.From.Options[j].Item.Name = singleTranslateWrap(f,
						v[i].Feature_specific.Expertise_options.From.Options[j].Item.Name)
				}
				for j := range v[i].Feature_specific.Invocations {
					v[i].Feature_specific.Invocations[j].Name = singleTranslateWrap(f,
						v[i].Feature_specific.Invocations[j].Name)
				}
				for j := range v[i].Feature_specific.Subfeature_options.From.Options {
					v[i].Feature_specific.Subfeature_options.From.Options[j].Item.Name = singleTranslateWrap(f,
						v[i].Feature_specific.Subfeature_options.From.Options[j].Item.Name)
				}
				v[i].Parent.Name = singleTranslateWrap(f, v[i].Parent.Name)
				v[i].Subclass.Name = singleTranslateWrap(f, v[i].Subclass.Name)
			}
		},
	},
	"5e-SRD-Languages.json": &handler[*models.SRDLanguages]{
		selector: func(v []*models.SRDLanguages, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = singleTranslateWrap(f, v[i].Desc)
				v[i].Script = singleTranslateWrap(f, v[i].Script)
				v[i].Type = singleTranslateWrap(f, v[i].Type)
				v[i].Typical_speakers = f(v[i].Typical_speakers)
			}
		},
	},
	"5e-SRD-Levels.json": &handler[*models.SRDLevels]{
		selector: func(v []*models.SRDLevels, f translateFunc) {
			for i := range v {
				v[i].Class.Name = singleTranslateWrap(f, v[i].Class.Name)
				for j := range v[i].Features {
					v[i].Features[j].Name = singleTranslateWrap(f, v[i].Features[j].Name)
				}
				v[i].Subclass.Name = singleTranslateWrap(f, v[i].Subclass.Name)
			}
		},
	},
	"5e-SRD-Magic-Items.json": &handler[*models.SRDMagicItems]{
		selector: func(v []*models.SRDMagicItems, f translateFunc) {
			for i := range v {
				v[i].Desc = f(v[i].Desc)
				v[i].Equipment_category.Name = singleTranslateWrap(f, v[i].Equipment_category.Name)
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Rarity.Name = singleTranslateWrap(f, v[i].Rarity.Name)
				for j := range v[i].Variants {
					v[i].Variants[j].Name = singleTranslateWrap(f, v[i].Variants[j].Name)
				}
			}
		},
	},
	"5e-SRD-Magic-Schools.json": &handler[*models.SRDMagicSchools]{
		selector: func(v []*models.SRDMagicSchools, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = singleTranslateWrap(f, v[i].Desc)
			}
		},
	},
	"5e-SRD-Monsters.json": &handler[*models.SRDMonsters]{
		selector: func(v []*models.SRDMonsters, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = singleTranslateWrap(f, v[i].Desc)
				for j := range v[i].Actions {
					v[i].Actions[j].Desc = singleTranslateWrap(f, v[i].Actions[j].Desc)

					for k := range v[i].Actions[j].Actions {
						v[i].Actions[j].Actions[k].Action_name = singleTranslateWrap(f,
							v[i].Actions[j].Actions[k].Action_name)
						v[i].Actions[j].Actions[k].Type = singleTranslateWrap(f,
							v[i].Actions[j].Actions[k].Type)
					}

					for k := range v[i].Actions[j].Damage {
						v[i].Actions[j].Damage[k].Damage_type.Name = singleTranslateWrap(f,
							v[i].Actions[j].Damage[k].Damage_type.Name)
					}

					v[i].Actions[j].Dc.Dc_type.Name = singleTranslateWrap(f,
						v[i].Actions[j].Dc.Dc_type.Name)

					v[i].Actions[j].Usage.Type = singleTranslateWrap(f,
						v[i].Actions[j].Usage.Type)
				}
				v[i].Alignment = singleTranslateWrap(f,
					v[i].Alignment)

				for j := range v[i].Condition_immunities {
					v[i].Condition_immunities[j].Name = singleTranslateWrap(f,
						v[i].Condition_immunities[j].Name)
				}

				v[i].Damage_immunities = f(v[i].Damage_immunities)
				v[i].Damage_resistances = f(v[i].Damage_resistances)
				v[i].Damage_vulnerabilities = f(v[i].Damage_vulnerabilities)
				for j := range v[i].Forms {
					v[i].Forms[j].Name = singleTranslateWrap(f,
						v[i].Forms[j].Name)
				}

				v[i].Languages = singleTranslateWrap(f,
					v[i].Languages)
				for j := range v[i].Legendary_actions {
					v[i].Legendary_actions[j].Name = singleTranslateWrap(f,
						v[i].Legendary_actions[j].Name)
					v[i].Legendary_actions[j].Desc = singleTranslateWrap(f,
						v[i].Legendary_actions[j].Desc)
					for k := range v[i].Legendary_actions[j].Damage {
						v[i].Legendary_actions[j].Damage[k].Damage_type.Name = singleTranslateWrap(f,
							v[i].Legendary_actions[j].Damage[k].Damage_type.Name)
					}
				}

				for j := range v[i].Proficiencies {
					v[i].Proficiencies[j].Proficiency.Name = singleTranslateWrap(f,
						v[i].Proficiencies[j].Proficiency.Name)
				}

				for j := range v[i].Reactions {
					v[i].Reactions[j].Desc = singleTranslateWrap(f,
						v[i].Reactions[j].Desc)
					v[i].Reactions[j].Name = singleTranslateWrap(f,
						v[i].Reactions[j].Name)
				}

				v[i].Senses.Blindsight = singleTranslateWrap(f,
					v[i].Senses.Blindsight)
				v[i].Senses.Darkvision = singleTranslateWrap(f,
					v[i].Senses.Darkvision)
				v[i].Senses.Tremorsense = singleTranslateWrap(f,
					v[i].Senses.Tremorsense)
				v[i].Senses.Truesight = singleTranslateWrap(f,
					v[i].Senses.Truesight)

				v[i].Size = singleTranslateWrap(f,
					v[i].Size)

				for j := range v[i].Special_abilities {
					v[i].Special_abilities[j].Dc.Dc_type.Name = singleTranslateWrap(f,
						v[i].Special_abilities[j].Dc.Dc_type.Name)
					v[i].Special_abilities[j].Name = singleTranslateWrap(f,
						v[i].Special_abilities[j].Name)
					v[i].Special_abilities[j].Desc = singleTranslateWrap(f,
						v[i].Special_abilities[j].Desc)
				}

				v[i].Speed.Burrow = singleTranslateWrap(f,
					v[i].Speed.Burrow)
				v[i].Speed.Climb = singleTranslateWrap(f,
					v[i].Speed.Climb)
				v[i].Speed.Fly = singleTranslateWrap(f,
					v[i].Speed.Fly)
				v[i].Speed.Swim = singleTranslateWrap(f,
					v[i].Speed.Swim)
				v[i].Speed.Walk = singleTranslateWrap(f,
					v[i].Speed.Walk)

				v[i].Subtype = singleTranslateWrap(f,
					v[i].Subtype)
				v[i].Type = singleTranslateWrap(f,
					v[i].Type)
			}
		},
	},
	"5e-SRD-Proficiencies.json": &handler[*models.SRDProficiencies]{
		selector: func(v []*models.SRDProficiencies, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				for j := range v[i].Classes {
					v[i].Classes[j].Name = singleTranslateWrap(f,
						v[i].Classes[j].Name)
				}
				for j := range v[i].Races {
					v[i].Races[j].Name = singleTranslateWrap(f,
						v[i].Races[j].Name)
				}

				v[i].Reference.Name = singleTranslateWrap(f, v[i].Reference.Name)
				v[i].Type = singleTranslateWrap(f, v[i].Type)
			}
		},
	},
	"5e-SRD-Races.json": &handler[*models.SRDRaces]{
		selector: func(v []*models.SRDRaces, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				for j := range v[i].Ability_bonus_options.From.Options {
					v[i].Ability_bonus_options.From.Options[j].Ability_score.Name = singleTranslateWrap(f,
						v[i].Ability_bonus_options.From.Options[j].Ability_score.Name)
				}
				for j := range v[i].Ability_bonuses {
					v[i].Ability_bonuses[j].Ability_score.Name = singleTranslateWrap(f,
						v[i].Ability_bonuses[j].Ability_score.Name)
				}

				v[i].Age = singleTranslateWrap(f, v[i].Age)
				v[i].Alignment = singleTranslateWrap(f, v[i].Alignment)
				v[i].Language_desc = singleTranslateWrap(f, v[i].Language_desc)

				for j := range v[i].Language_options.From.Options {
					v[i].Language_options.From.Options[j].Item.Name = singleTranslateWrap(f,
						v[i].Language_options.From.Options[j].Item.Name)
				}

				for j := range v[i].Languages {
					v[i].Languages[j].Name = singleTranslateWrap(f,
						v[i].Languages[j].Name)
				}

				v[i].Size = singleTranslateWrap(f, v[i].Size)
				v[i].Size_description = singleTranslateWrap(f, v[i].Size_description)

				for j := range v[i].Starting_proficiencies {
					v[i].Starting_proficiencies[j].Name = singleTranslateWrap(f,
						v[i].Starting_proficiencies[j].Name)
				}

				v[i].Starting_proficiency_options.Desc = singleTranslateWrap(f,
					v[i].Starting_proficiency_options.Desc)
				for j := range v[i].Starting_proficiency_options.From.Options {
					v[i].Starting_proficiency_options.From.Options[j].Item.Name = singleTranslateWrap(f,
						v[i].Starting_proficiency_options.From.Options[j].Item.Name)
				}

				for j := range v[i].Subraces {
					v[i].Subraces[j].Name = singleTranslateWrap(f,
						v[i].Subraces[j].Name)
				}

				for j := range v[i].Traits {
					v[i].Traits[j].Name = singleTranslateWrap(f,
						v[i].Traits[j].Name)
				}
			}
		},
	},
	"5e-SRD-Rule-Sections.json": &handler[*models.SRDRuleSections]{
		selector: func(v []*models.SRDRuleSections, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				descs := strings.Split(v[i].Desc, ".")
				for j := range descs {
					descs[j] = singleTranslateWrap(f, descs[j])
				}

				v[i].Desc = strings.Join(descs, ".")
			}
		},
	},
	"5e-SRD-Rules.json": &handler[*models.SRDRules]{
		selector: func(v []*models.SRDRules, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = singleTranslateWrap(f, v[i].Desc)
				for j := range v[i].Subsections {
					v[i].Subsections[j].Name = singleTranslateWrap(f,
						v[i].Subsections[j].Name)
				}
			}
		},
	},
	"5e-SRD-Skills.json": &handler[*models.SRDSkills]{
		selector: func(v []*models.SRDSkills, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = f(v[i].Desc)
				v[i].Ability_score.Name = singleTranslateWrap(f, v[i].Ability_score.Name)
			}
		},
	},
	"5e-SRD-Spells.json": &handler[*models.SRDSpells]{
		selector: func(v []*models.SRDSpells, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = f(v[i].Desc)

				v[i].Area_of_effect.Type = singleTranslateWrap(f, v[i].Area_of_effect.Type)
				v[i].Attack_type = singleTranslateWrap(f, v[i].Attack_type)
				v[i].Casting_time = singleTranslateWrap(f, v[i].Casting_time)

				for j := range v[i].Classes {
					v[i].Classes[j].Name = singleTranslateWrap(f,
						v[i].Classes[j].Name)
				}

				v[i].Components = f(v[i].Components)
				v[i].Damage.Damage_type.Name = singleTranslateWrap(f,
					v[i].Damage.Damage_type.Name)
				v[i].Dc.Dc_type.Name = singleTranslateWrap(f,
					v[i].Dc.Dc_type.Name)
				v[i].Dc.Desc = singleTranslateWrap(f,
					v[i].Dc.Desc)

				v[i].Duration = singleTranslateWrap(f,
					v[i].Duration)

				v[i].Higher_level = f(v[i].Higher_level)
				v[i].Material = singleTranslateWrap(f,
					v[i].Material)
				v[i].Range = singleTranslateWrap(f,
					v[i].Range)
				v[i].School.Name = singleTranslateWrap(f,
					v[i].School.Name)

				for j := range v[i].Subclasses {
					v[i].Subclasses[j].Name = singleTranslateWrap(f,
						v[i].Subclasses[j].Name)
				}
			}
		},
	},
	"5e-SRD-Subclasses.json": &handler[*models.SRDSubclasses]{
		selector: func(v []*models.SRDSubclasses, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = f(v[i].Desc)

				v[i].Class.Name = singleTranslateWrap(f, v[i].Class.Name)
				for j := range v[i].Spells {
					v[i].Spells[j].Spell.Name = singleTranslateWrap(f,
						v[i].Spells[j].Spell.Name)
					for k := range v[i].Spells[j].Prerequisites {
						v[i].Spells[j].Prerequisites[k].Name = singleTranslateWrap(f,
							v[i].Spells[j].Prerequisites[k].Name)
					}
				}
				v[i].Subclass_flavor = singleTranslateWrap(f, v[i].Subclass_flavor)
			}
		},
	},
	"5e-SRD-Subraces.json": &handler[*models.SRDSubraces]{
		selector: func(v []*models.SRDSubraces, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = singleTranslateWrap(f, v[i].Desc)

				for j := range v[i].Ability_bonuses {
					v[i].Ability_bonuses[j].Ability_score.Name = singleTranslateWrap(f,
						v[i].Ability_bonuses[j].Ability_score.Name)
				}

				for j := range v[i].Language_options.From.Options {
					v[i].Language_options.From.Options[j].Item.Name = singleTranslateWrap(f,
						v[i].Language_options.From.Options[j].Item.Name)
				}

				v[i].Race.Name = singleTranslateWrap(f, v[i].Race.Name)
				for j := range v[i].Racial_traits {
					v[i].Racial_traits[j].Name = singleTranslateWrap(f, v[i].Racial_traits[j].Name)
				}

				for j := range v[i].Starting_proficiencies {
					v[i].Starting_proficiencies[j].Name = singleTranslateWrap(f, v[i].Starting_proficiencies[j].Name)
				}
			}
		},
	},
	"5e-SRD-Traits.json": &handler[*models.SRDTraits]{
		selector: func(v []*models.SRDTraits, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = f(v[i].Desc)

				for j := range v[i].Language_options.From.Options {
					v[i].Language_options.From.Options[j].Item.Name = singleTranslateWrap(f,
						v[i].Language_options.From.Options[j].Item.Name)
				}

				v[i].Parent.Name = singleTranslateWrap(f, v[i].Parent.Name)
				for j := range v[i].Proficiencies {
					v[i].Proficiencies[j].Name = singleTranslateWrap(f,
						v[i].Proficiencies[j].Name)
				}

				for j := range v[i].Proficiency_choices.From.Options {
					v[i].Proficiency_choices.From.Options[j].Item.Name = singleTranslateWrap(f,
						v[i].Proficiency_choices.From.Options[j].Item.Name)
				}

				for j := range v[i].Races {
					v[i].Races[j].Name = singleTranslateWrap(f,
						v[i].Races[j].Name)
				}

				for j := range v[i].Subraces {
					v[i].Subraces[j].Name = singleTranslateWrap(f,
						v[i].Subraces[j].Name)
				}

				for j := range v[i].Trait_specific.Breath_weapon.Damage {
					v[i].Trait_specific.Breath_weapon.Damage[j].Damage_type.Name = singleTranslateWrap(f,
						v[i].Trait_specific.Breath_weapon.Damage[j].Damage_type.Name)
				}

				v[i].Trait_specific.Breath_weapon.Dc.Dc_type.Name = singleTranslateWrap(f,
					v[i].Trait_specific.Breath_weapon.Dc.Dc_type.Name)
				v[i].Trait_specific.Breath_weapon.Desc = singleTranslateWrap(f,
					v[i].Trait_specific.Breath_weapon.Desc)
				v[i].Trait_specific.Breath_weapon.Name = singleTranslateWrap(f,
					v[i].Trait_specific.Breath_weapon.Name)
				v[i].Trait_specific.Breath_weapon.Usage.Type = singleTranslateWrap(f,
					v[i].Trait_specific.Breath_weapon.Usage.Type)

				v[i].Trait_specific.Damage_type.Name = singleTranslateWrap(f,
					v[i].Trait_specific.Damage_type.Name)

				for j := range v[i].Trait_specific.Spell_options.From.Options {
					v[i].Trait_specific.Spell_options.From.Options[j].Item.Name = singleTranslateWrap(f,
						v[i].Trait_specific.Spell_options.From.Options[j].Item.Name)
				}

				for j := range v[i].Trait_specific.Subtrait_options.From.Options {
					v[i].Trait_specific.Subtrait_options.From.Options[j].Item.Name = singleTranslateWrap(f,
						v[i].Trait_specific.Subtrait_options.From.Options[j].Item.Name)
				}
			}
		},
	},
	"5e-SRD-Weapon-Properties.json": &handler[*models.SRDWeaponProperties]{
		selector: func(v []*models.SRDWeaponProperties, f translateFunc) {
			for i := range v {
				v[i].Name = singleTranslateWrap(f, v[i].Name)
				v[i].Desc = f(v[i].Desc)
			}
		},
	},
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	srcPath := filepath.Join(wd, "..", "src")
	destPath := filepath.Join(wd, "..", "ru")
	t := newTranslator(os.Getenv("APIKEY"))
	defer t.close()

	for file, m := range orm {
		filePath := filepath.Join(srcPath, file)

		buff, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("failed to read file %s: %v\n", filePath, err)
			return
		}

		data, err := m.handle(buff, t)
		if err != nil {
			fmt.Printf("failed to handle: %v\n", err)
			return
		}

		destFilePath := filepath.Join(destPath, file)
		if err := os.WriteFile(destFilePath, data, 0644); err != nil {
			fmt.Printf("failed to write translated: %v\n", err)
			return
		}
	}
}
