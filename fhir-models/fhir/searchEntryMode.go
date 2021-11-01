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

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// SearchEntryMode is documented here http://hl7.org/fhir/ValueSet/search-entry-mode
type SearchEntryMode int

const (
	SearchEntryModeMatch SearchEntryMode = iota
	SearchEntryModeInclude
	SearchEntryModeOutcome
)

func (code SearchEntryMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SearchEntryMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "match":
		*code = SearchEntryModeMatch
	case "include":
		*code = SearchEntryModeInclude
	case "outcome":
		*code = SearchEntryModeOutcome
	default:
		return fmt.Errorf("unknown SearchEntryMode code `%s`", s)
	}
	return nil
}
func (code SearchEntryMode) String() string {
	return code.Code()
}
func (code SearchEntryMode) Code() string {
	switch code {
	case SearchEntryModeMatch:
		return "match"
	case SearchEntryModeInclude:
		return "include"
	case SearchEntryModeOutcome:
		return "outcome"
	}
	return "<unknown>"
}
func (code SearchEntryMode) Display() string {
	switch code {
	case SearchEntryModeMatch:
		return "Match"
	case SearchEntryModeInclude:
		return "Include"
	case SearchEntryModeOutcome:
		return "Outcome"
	}
	return "<unknown>"
}
func (code SearchEntryMode) Definition() string {
	switch code {
	case SearchEntryModeMatch:
		return "This resource matched the search specification."
	case SearchEntryModeInclude:
		return "This resource is returned because it is referred to from another resource in the search set."
	case SearchEntryModeOutcome:
		return "An OperationOutcome that provides additional information about the processing of a search."
	}
	return "<unknown>"
}
