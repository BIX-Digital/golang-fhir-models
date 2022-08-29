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

// VisionPrescription is documented here http://hl7.org/fhir/StructureDefinition/VisionPrescription
type VisionPrescription struct {
	ID                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes          `bson:"status" json:"status"`
	Created           string                                `bson:"created" json:"created"`
	Patient           Reference                             `bson:"patient" json:"patient"`
	Encounter         *Reference                            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateWritten       string                                `bson:"dateWritten" json:"dateWritten"`
	Prescriber        Reference                             `bson:"prescriber" json:"prescriber"`
	LensSpecification []VisionPrescriptionLensSpecification `bson:"lensSpecification" json:"lensSpecification"`
}
type VisionPrescriptionLensSpecification struct {
	ID                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Product           CodeableConcept                            `bson:"product" json:"product"`
	Eye               VisionEyes                                 `bson:"eye" json:"eye"`
	Sphere            *float64                                   `bson:"sphere,omitempty" json:"sphere,omitempty"`
	Cylinder          *float64                                   `bson:"cylinder,omitempty" json:"cylinder,omitempty"`
	Axis              *int                                       `bson:"axis,omitempty" json:"axis,omitempty"`
	Prism             []VisionPrescriptionLensSpecificationPrism `bson:"prism,omitempty" json:"prism,omitempty"`
	Add               *float64                                   `bson:"add,omitempty" json:"add,omitempty"`
	Power             *float64                                   `bson:"power,omitempty" json:"power,omitempty"`
	BackCurve         *float64                                   `bson:"backCurve,omitempty" json:"backCurve,omitempty"`
	Diameter          *float64                                   `bson:"diameter,omitempty" json:"diameter,omitempty"`
	Duration          *Quantity                                  `bson:"duration,omitempty" json:"duration,omitempty"`
	Color             *string                                    `bson:"color,omitempty" json:"color,omitempty"`
	Brand             *string                                    `bson:"brand,omitempty" json:"brand,omitempty"`
	Note              []Annotation                               `bson:"note,omitempty" json:"note,omitempty"`
}
type VisionPrescriptionLensSpecificationPrism struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Amount            float64     `bson:"amount" json:"amount"`
	Base              VisionBase  `bson:"base" json:"base"`
}
type OtherVisionPrescription VisionPrescription

// MarshalJSON marshals the given VisionPrescription as JSON into a byte slice
func (r VisionPrescription) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherVisionPrescription
		ResourceType string `json:"resourceType"`
	}{
		OtherVisionPrescription: OtherVisionPrescription(r),
		ResourceType:            "VisionPrescription",
	})
	return buffer.Bytes(), err
}

// UnmarshalVisionPrescription unmarshals a VisionPrescription.
func UnmarshalVisionPrescription(b []byte) (VisionPrescription, error) {
	var visionPrescription VisionPrescription
	if err := json.Unmarshal(b, &visionPrescription); err != nil {
		return visionPrescription, err
	}
	return visionPrescription, nil
}
