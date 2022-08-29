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

// CodeSearchSupport is documented here http://hl7.org/fhir/ValueSet/code-search-support
type CodeSearchSupport int

const (
	CodeSearchSupportExplicit CodeSearchSupport = iota
	CodeSearchSupportAll
)

func (code CodeSearchSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CodeSearchSupport) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "explicit":
		*code = CodeSearchSupportExplicit
	case "all":
		*code = CodeSearchSupportAll
	default:
		return fmt.Errorf("unknown CodeSearchSupport code `%s`", s)
	}
	return nil
}
func (code CodeSearchSupport) String() string {
	return code.Code()
}
func (code CodeSearchSupport) Code() string {
	switch code {
	case CodeSearchSupportExplicit:
		return "explicit"
	case CodeSearchSupportAll:
		return "all"
	}
	return "<unknown>"
}
func (code CodeSearchSupport) Display() string {
	switch code {
	case CodeSearchSupportExplicit:
		return "Explicit Codes"
	case CodeSearchSupportAll:
		return "Implicit Codes"
	}
	return "<unknown>"
}
func (code CodeSearchSupport) Definition() string {
	switch code {
	case CodeSearchSupportExplicit:
		return "The search for code on ValueSet only includes codes explicitly detailed on includes or expansions."
	case CodeSearchSupportAll:
		return "The search for code on ValueSet only includes all codes based on the expansion of the value set."
	}
	return "<unknown>"
}
