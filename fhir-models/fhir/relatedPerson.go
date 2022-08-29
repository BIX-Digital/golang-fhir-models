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

// RelatedPerson is documented here http://hl7.org/fhir/StructureDefinition/RelatedPerson
type RelatedPerson struct {
	ID                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                        `bson:"active,omitempty" json:"active,omitempty"`
	Patient           Reference                    `bson:"patient" json:"patient"`
	Relationship      []CodeableConcept            `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name              []HumanName                  `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint               `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender            *AdministrativeGender        `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate         *string                      `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address           []Address                    `bson:"address,omitempty" json:"address,omitempty"`
	Photo             []Attachment                 `bson:"photo,omitempty" json:"photo,omitempty"`
	Period            *Period                      `bson:"period,omitempty" json:"period,omitempty"`
	Communication     []RelatedPersonCommunication `bson:"communication,omitempty" json:"communication,omitempty"`
}
type RelatedPersonCommunication struct {
	ID                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          CodeableConcept `bson:"language" json:"language"`
	Preferred         *bool           `bson:"preferred,omitempty" json:"preferred,omitempty"`
}
type OtherRelatedPerson RelatedPerson

// MarshalJSON marshals the given RelatedPerson as JSON into a byte slice
func (r RelatedPerson) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherRelatedPerson
		ResourceType string `json:"resourceType"`
	}{
		OtherRelatedPerson: OtherRelatedPerson(r),
		ResourceType:       "RelatedPerson",
	})
	return buffer.Bytes(), err
}

// UnmarshalRelatedPerson unmarshals a RelatedPerson.
func UnmarshalRelatedPerson(b []byte) (RelatedPerson, error) {
	var relatedPerson RelatedPerson
	if err := json.Unmarshal(b, &relatedPerson); err != nil {
		return relatedPerson, err
	}
	return relatedPerson, nil
}
