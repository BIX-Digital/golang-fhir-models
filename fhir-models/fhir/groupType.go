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
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// GroupType is documented here http://hl7.org/fhir/ValueSet/group-type
type GroupType int

const (
	GroupTypePerson GroupType = iota
	GroupTypeAnimal
	GroupTypePractitioner
	GroupTypeDevice
	GroupTypeMedication
	GroupTypeSubstance
)

func (code GroupType) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *GroupType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "person":
		*code = GroupTypePerson
	case "animal":
		*code = GroupTypeAnimal
	case "practitioner":
		*code = GroupTypePractitioner
	case "device":
		*code = GroupTypeDevice
	case "medication":
		*code = GroupTypeMedication
	case "substance":
		*code = GroupTypeSubstance
	default:
		return fmt.Errorf("unknown GroupType code `%s`", s)
	}
	return nil
}
func (code GroupType) String() string {
	return code.Code()
}
func (code GroupType) Code() string {
	switch code {
	case GroupTypePerson:
		return "person"
	case GroupTypeAnimal:
		return "animal"
	case GroupTypePractitioner:
		return "practitioner"
	case GroupTypeDevice:
		return "device"
	case GroupTypeMedication:
		return "medication"
	case GroupTypeSubstance:
		return "substance"
	}
	return "<unknown>"
}
func (code GroupType) Display() string {
	switch code {
	case GroupTypePerson:
		return "Person"
	case GroupTypeAnimal:
		return "Animal"
	case GroupTypePractitioner:
		return "Practitioner"
	case GroupTypeDevice:
		return "Device"
	case GroupTypeMedication:
		return "Medication"
	case GroupTypeSubstance:
		return "Substance"
	}
	return "<unknown>"
}
func (code GroupType) Definition() string {
	switch code {
	case GroupTypePerson:
		return "Group contains \"person\" Patient resources."
	case GroupTypeAnimal:
		return "Group contains \"animal\" Patient resources."
	case GroupTypePractitioner:
		return "Group contains healthcare practitioner resources (Practitioner or PractitionerRole)."
	case GroupTypeDevice:
		return "Group contains Device resources."
	case GroupTypeMedication:
		return "Group contains Medication resources."
	case GroupTypeSubstance:
		return "Group contains Substance resources."
	}
	return "<unknown>"
}
