// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"bytes"
	"encoding/json"
)

// THIS FILE IS GENERATED BY https://github.com/bix-digital/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// SubstanceSourceMaterial is documented here http://hl7.org/fhir/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterial struct {
	ID                   *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                                      `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SourceMaterialClass  *CodeableConcept                             `bson:"sourceMaterialClass,omitempty" json:"sourceMaterialClass,omitempty"`
	SourceMaterialType   *CodeableConcept                             `bson:"sourceMaterialType,omitempty" json:"sourceMaterialType,omitempty"`
	SourceMaterialState  *CodeableConcept                             `bson:"sourceMaterialState,omitempty" json:"sourceMaterialState,omitempty"`
	OrganismId           *Identifier                                  `bson:"organismId,omitempty" json:"organismId,omitempty"`
	OrganismName         *string                                      `bson:"organismName,omitempty" json:"organismName,omitempty"`
	ParentSubstanceId    []Identifier                                 `bson:"parentSubstanceId,omitempty" json:"parentSubstanceId,omitempty"`
	ParentSubstanceName  []string                                     `bson:"parentSubstanceName,omitempty" json:"parentSubstanceName,omitempty"`
	CountryOfOrigin      []CodeableConcept                            `bson:"countryOfOrigin,omitempty" json:"countryOfOrigin,omitempty"`
	GeographicalLocation []string                                     `bson:"geographicalLocation,omitempty" json:"geographicalLocation,omitempty"`
	DevelopmentStage     *CodeableConcept                             `bson:"developmentStage,omitempty" json:"developmentStage,omitempty"`
	FractionDescription  []SubstanceSourceMaterialFractionDescription `bson:"fractionDescription,omitempty" json:"fractionDescription,omitempty"`
	Organism             *SubstanceSourceMaterialOrganism             `bson:"organism,omitempty" json:"organism,omitempty"`
	PartDescription      []SubstanceSourceMaterialPartDescription     `bson:"partDescription,omitempty" json:"partDescription,omitempty"`
}
type SubstanceSourceMaterialFractionDescription struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Fraction          *string          `bson:"fraction,omitempty" json:"fraction,omitempty"`
	MaterialType      *CodeableConcept `bson:"materialType,omitempty" json:"materialType,omitempty"`
}
type SubstanceSourceMaterialOrganism struct {
	ID                       *string                                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []Extension                                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension                                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Family                   *CodeableConcept                                `bson:"family,omitempty" json:"family,omitempty"`
	Genus                    *CodeableConcept                                `bson:"genus,omitempty" json:"genus,omitempty"`
	Species                  *CodeableConcept                                `bson:"species,omitempty" json:"species,omitempty"`
	IntraspecificType        *CodeableConcept                                `bson:"intraspecificType,omitempty" json:"intraspecificType,omitempty"`
	IntraspecificDescription *string                                         `bson:"intraspecificDescription,omitempty" json:"intraspecificDescription,omitempty"`
	Author                   []SubstanceSourceMaterialOrganismAuthor         `bson:"author,omitempty" json:"author,omitempty"`
	Hybrid                   *SubstanceSourceMaterialOrganismHybrid          `bson:"hybrid,omitempty" json:"hybrid,omitempty"`
	OrganismGeneral          *SubstanceSourceMaterialOrganismOrganismGeneral `bson:"organismGeneral,omitempty" json:"organismGeneral,omitempty"`
}
type SubstanceSourceMaterialOrganismAuthor struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	AuthorType        *CodeableConcept `bson:"authorType,omitempty" json:"authorType,omitempty"`
	AuthorDescription *string          `bson:"authorDescription,omitempty" json:"authorDescription,omitempty"`
}
type SubstanceSourceMaterialOrganismHybrid struct {
	ID                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	MaternalOrganismId   *string          `bson:"maternalOrganismId,omitempty" json:"maternalOrganismId,omitempty"`
	MaternalOrganismName *string          `bson:"maternalOrganismName,omitempty" json:"maternalOrganismName,omitempty"`
	PaternalOrganismId   *string          `bson:"paternalOrganismId,omitempty" json:"paternalOrganismId,omitempty"`
	PaternalOrganismName *string          `bson:"paternalOrganismName,omitempty" json:"paternalOrganismName,omitempty"`
	HybridType           *CodeableConcept `bson:"hybridType,omitempty" json:"hybridType,omitempty"`
}
type SubstanceSourceMaterialOrganismOrganismGeneral struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Kingdom           *CodeableConcept `bson:"kingdom,omitempty" json:"kingdom,omitempty"`
	Phylum            *CodeableConcept `bson:"phylum,omitempty" json:"phylum,omitempty"`
	Class             *CodeableConcept `bson:"class,omitempty" json:"class,omitempty"`
	Order             *CodeableConcept `bson:"order,omitempty" json:"order,omitempty"`
}
type SubstanceSourceMaterialPartDescription struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Part              *CodeableConcept `bson:"part,omitempty" json:"part,omitempty"`
	PartLocation      *CodeableConcept `bson:"partLocation,omitempty" json:"partLocation,omitempty"`
}
type OtherSubstanceSourceMaterial SubstanceSourceMaterial

// MarshalJSON marshals the given SubstanceSourceMaterial as JSON into a byte slice
func (r SubstanceSourceMaterial) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherSubstanceSourceMaterial
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceSourceMaterial: OtherSubstanceSourceMaterial(r),
		ResourceType:                 "SubstanceSourceMaterial",
	})
	return buffer.Bytes(), err
}

// UnmarshalSubstanceSourceMaterial unmarshals a SubstanceSourceMaterial.
func UnmarshalSubstanceSourceMaterial(b []byte) (SubstanceSourceMaterial, error) {
	var substanceSourceMaterial SubstanceSourceMaterial
	if err := json.Unmarshal(b, &substanceSourceMaterial); err != nil {
		return substanceSourceMaterial, err
	}
	return substanceSourceMaterial, nil
}
