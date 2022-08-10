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

// DocumentReferenceStatus is documented here http://hl7.org/fhir/ValueSet/document-reference-status
type DocumentReferenceStatus int

const (
	DocumentReferenceStatusCurrent DocumentReferenceStatus = iota
	DocumentReferenceStatusSuperseded
	DocumentReferenceStatusEnteredInError
)

func (code DocumentReferenceStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *DocumentReferenceStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "current":
		*code = DocumentReferenceStatusCurrent
	case "superseded":
		*code = DocumentReferenceStatusSuperseded
	case "entered-in-error":
		*code = DocumentReferenceStatusEnteredInError
	default:
		return fmt.Errorf("unknown DocumentReferenceStatus code `%s`", s)
	}
	return nil
}
func (code DocumentReferenceStatus) String() string {
	return code.Code()
}
func (code DocumentReferenceStatus) Code() string {
	switch code {
	case DocumentReferenceStatusCurrent:
		return "current"
	case DocumentReferenceStatusSuperseded:
		return "superseded"
	case DocumentReferenceStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code DocumentReferenceStatus) Display() string {
	switch code {
	case DocumentReferenceStatusCurrent:
		return "Current"
	case DocumentReferenceStatusSuperseded:
		return "Superseded"
	case DocumentReferenceStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code DocumentReferenceStatus) Definition() string {
	switch code {
	case DocumentReferenceStatusCurrent:
		return "This is the current reference for this document."
	case DocumentReferenceStatusSuperseded:
		return "This reference has been superseded by another reference."
	case DocumentReferenceStatusEnteredInError:
		return "This reference was created in error."
	}
	return "<unknown>"
}
