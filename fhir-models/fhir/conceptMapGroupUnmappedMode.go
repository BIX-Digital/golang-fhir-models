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

// THIS FILE IS GENERATED BY https://github.com/bix-digital/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ConceptMapGroupUnmappedMode is documented here http://hl7.org/fhir/ValueSet/conceptmap-unmapped-mode
type ConceptMapGroupUnmappedMode int

const (
	ConceptMapGroupUnmappedModeProvided ConceptMapGroupUnmappedMode = iota
	ConceptMapGroupUnmappedModeFixed
	ConceptMapGroupUnmappedModeOtherMap
)

func (code ConceptMapGroupUnmappedMode) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *ConceptMapGroupUnmappedMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "provided":
		*code = ConceptMapGroupUnmappedModeProvided
	case "fixed":
		*code = ConceptMapGroupUnmappedModeFixed
	case "other-map":
		*code = ConceptMapGroupUnmappedModeOtherMap
	default:
		return fmt.Errorf("unknown ConceptMapGroupUnmappedMode code `%s`", s)
	}
	return nil
}
func (code ConceptMapGroupUnmappedMode) String() string {
	return code.Code()
}
func (code ConceptMapGroupUnmappedMode) Code() string {
	switch code {
	case ConceptMapGroupUnmappedModeProvided:
		return "provided"
	case ConceptMapGroupUnmappedModeFixed:
		return "fixed"
	case ConceptMapGroupUnmappedModeOtherMap:
		return "other-map"
	}
	return "<unknown>"
}
func (code ConceptMapGroupUnmappedMode) Display() string {
	switch code {
	case ConceptMapGroupUnmappedModeProvided:
		return "Provided Code"
	case ConceptMapGroupUnmappedModeFixed:
		return "Fixed Code"
	case ConceptMapGroupUnmappedModeOtherMap:
		return "Other Map"
	}
	return "<unknown>"
}
func (code ConceptMapGroupUnmappedMode) Definition() string {
	switch code {
	case ConceptMapGroupUnmappedModeProvided:
		return "Use the code as provided in the $translate request."
	case ConceptMapGroupUnmappedModeFixed:
		return "Use the code explicitly provided in the group.unmapped."
	case ConceptMapGroupUnmappedModeOtherMap:
		return "Use the map identified by the canonical URL in the url element."
	}
	return "<unknown>"
}
