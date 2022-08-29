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

// ListStatus is documented here http://hl7.org/fhir/ValueSet/list-status
type ListStatus int

const (
	ListStatusCurrent ListStatus = iota
	ListStatusRetired
	ListStatusEnteredInError
)

func (code ListStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *ListStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "current":
		*code = ListStatusCurrent
	case "retired":
		*code = ListStatusRetired
	case "entered-in-error":
		*code = ListStatusEnteredInError
	default:
		return fmt.Errorf("unknown ListStatus code `%s`", s)
	}
	return nil
}
func (code ListStatus) String() string {
	return code.Code()
}
func (code ListStatus) Code() string {
	switch code {
	case ListStatusCurrent:
		return "current"
	case ListStatusRetired:
		return "retired"
	case ListStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code ListStatus) Display() string {
	switch code {
	case ListStatusCurrent:
		return "Current"
	case ListStatusRetired:
		return "Retired"
	case ListStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code ListStatus) Definition() string {
	switch code {
	case ListStatusCurrent:
		return "The list is considered to be an active part of the patient's record."
	case ListStatusRetired:
		return "The list is \"old\" and should no longer be considered accurate or relevant."
	case ListStatusEnteredInError:
		return "The list was never accurate.  It is retained for medico-legal purposes only."
	}
	return "<unknown>"
}
