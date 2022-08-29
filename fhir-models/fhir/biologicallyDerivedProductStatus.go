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

// BiologicallyDerivedProductStatus is documented here http://hl7.org/fhir/ValueSet/product-status
type BiologicallyDerivedProductStatus int

const (
	BiologicallyDerivedProductStatusAvailable BiologicallyDerivedProductStatus = iota
	BiologicallyDerivedProductStatusUnavailable
)

func (code BiologicallyDerivedProductStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *BiologicallyDerivedProductStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "available":
		*code = BiologicallyDerivedProductStatusAvailable
	case "unavailable":
		*code = BiologicallyDerivedProductStatusUnavailable
	default:
		return fmt.Errorf("unknown BiologicallyDerivedProductStatus code `%s`", s)
	}
	return nil
}
func (code BiologicallyDerivedProductStatus) String() string {
	return code.Code()
}
func (code BiologicallyDerivedProductStatus) Code() string {
	switch code {
	case BiologicallyDerivedProductStatusAvailable:
		return "available"
	case BiologicallyDerivedProductStatusUnavailable:
		return "unavailable"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductStatus) Display() string {
	switch code {
	case BiologicallyDerivedProductStatusAvailable:
		return "Available"
	case BiologicallyDerivedProductStatusUnavailable:
		return "Unavailable"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductStatus) Definition() string {
	switch code {
	case BiologicallyDerivedProductStatusAvailable:
		return "Product is currently available for use."
	case BiologicallyDerivedProductStatusUnavailable:
		return "Product is not currently available for use."
	}
	return "<unknown>"
}
