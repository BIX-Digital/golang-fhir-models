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
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/bix-digital/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// AllergyIntoleranceCategory is documented here http://hl7.org/fhir/ValueSet/allergy-intolerance-category
type AllergyIntoleranceCategory int

const (
	AllergyIntoleranceCategoryFood AllergyIntoleranceCategory = iota
	AllergyIntoleranceCategoryMedication
	AllergyIntoleranceCategoryEnvironment
	AllergyIntoleranceCategoryBiologic
)

func (code AllergyIntoleranceCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AllergyIntoleranceCategory) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "food":
		*code = AllergyIntoleranceCategoryFood
	case "medication":
		*code = AllergyIntoleranceCategoryMedication
	case "environment":
		*code = AllergyIntoleranceCategoryEnvironment
	case "biologic":
		*code = AllergyIntoleranceCategoryBiologic
	default:
		return fmt.Errorf("unknown AllergyIntoleranceCategory code `%s`", s)
	}
	return nil
}
func (code AllergyIntoleranceCategory) String() string {
	return code.Code()
}
func (code AllergyIntoleranceCategory) Code() string {
	switch code {
	case AllergyIntoleranceCategoryFood:
		return "food"
	case AllergyIntoleranceCategoryMedication:
		return "medication"
	case AllergyIntoleranceCategoryEnvironment:
		return "environment"
	case AllergyIntoleranceCategoryBiologic:
		return "biologic"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceCategory) Display() string {
	switch code {
	case AllergyIntoleranceCategoryFood:
		return "Food"
	case AllergyIntoleranceCategoryMedication:
		return "Medication"
	case AllergyIntoleranceCategoryEnvironment:
		return "Environment"
	case AllergyIntoleranceCategoryBiologic:
		return "Biologic"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceCategory) Definition() string {
	switch code {
	case AllergyIntoleranceCategoryFood:
		return "Any substance consumed to provide nutritional support for the body."
	case AllergyIntoleranceCategoryMedication:
		return "Substances administered to achieve a physiological effect."
	case AllergyIntoleranceCategoryEnvironment:
		return "Any substances that are encountered in the environment, including any substance not already classified as food, medication, or biologic."
	case AllergyIntoleranceCategoryBiologic:
		return "A preparation that is synthesized from living organisms or their products, especially a human or animal protein, such as a hormone or antitoxin, that is used as a diagnostic, preventive, or therapeutic agent. Examples of biologic medications include: vaccines; allergenic extracts, which are used for both diagnosis and treatment (for example, allergy shots); gene therapies; cellular therapies.  There are other biologic products, such as tissues, which are not typically associated with allergies."
	}
	return "<unknown>"
}
