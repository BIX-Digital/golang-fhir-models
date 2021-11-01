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

// PublicationStatus is documented here http://hl7.org/fhir/ValueSet/publication-status
type PublicationStatus int

const (
	PublicationStatusDraft PublicationStatus = iota
	PublicationStatusActive
	PublicationStatusRetired
	PublicationStatusUnknown
)

func (code PublicationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *PublicationStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "draft":
		*code = PublicationStatusDraft
	case "active":
		*code = PublicationStatusActive
	case "retired":
		*code = PublicationStatusRetired
	case "unknown":
		*code = PublicationStatusUnknown
	default:
		return fmt.Errorf("unknown PublicationStatus code `%s`", s)
	}
	return nil
}
func (code PublicationStatus) String() string {
	return code.Code()
}
func (code PublicationStatus) Code() string {
	switch code {
	case PublicationStatusDraft:
		return "draft"
	case PublicationStatusActive:
		return "active"
	case PublicationStatusRetired:
		return "retired"
	case PublicationStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code PublicationStatus) Display() string {
	switch code {
	case PublicationStatusDraft:
		return "Draft"
	case PublicationStatusActive:
		return "Active"
	case PublicationStatusRetired:
		return "Retired"
	case PublicationStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code PublicationStatus) Definition() string {
	switch code {
	case PublicationStatusDraft:
		return "This resource is still under development and is not yet considered to be ready for normal use."
	case PublicationStatusActive:
		return "This resource is ready for normal use."
	case PublicationStatusRetired:
		return "This resource has been withdrawn or superseded and should no longer be used."
	case PublicationStatusUnknown:
		return "The authoring system does not know which of the status values currently applies for this resource.  Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
