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

// QuantityComparator is documented here http://hl7.org/fhir/ValueSet/quantity-comparator
type QuantityComparator int

const (
	QuantityComparatorLessThan QuantityComparator = iota
	QuantityComparatorLessOrEquals
	QuantityComparatorGreaterOrEquals
	QuantityComparatorGreaterThan
)

func (code QuantityComparator) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *QuantityComparator) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "<":
		*code = QuantityComparatorLessThan
	case "<=":
		*code = QuantityComparatorLessOrEquals
	case ">=":
		*code = QuantityComparatorGreaterOrEquals
	case ">":
		*code = QuantityComparatorGreaterThan
	default:
		return fmt.Errorf("unknown QuantityComparator code `%s`", s)
	}
	return nil
}
func (code QuantityComparator) String() string {
	return code.Code()
}
func (code QuantityComparator) Code() string {
	switch code {
	case QuantityComparatorLessThan:
		return "<"
	case QuantityComparatorLessOrEquals:
		return "<="
	case QuantityComparatorGreaterOrEquals:
		return ">="
	case QuantityComparatorGreaterThan:
		return ">"
	}
	return "<unknown>"
}
func (code QuantityComparator) Display() string {
	switch code {
	case QuantityComparatorLessThan:
		return "Less than"
	case QuantityComparatorLessOrEquals:
		return "Less or Equal to"
	case QuantityComparatorGreaterOrEquals:
		return "Greater or Equal to"
	case QuantityComparatorGreaterThan:
		return "Greater than"
	}
	return "<unknown>"
}
func (code QuantityComparator) Definition() string {
	switch code {
	case QuantityComparatorLessThan:
		return "The actual value is less than the given value."
	case QuantityComparatorLessOrEquals:
		return "The actual value is less than or equal to the given value."
	case QuantityComparatorGreaterOrEquals:
		return "The actual value is greater than or equal to the given value."
	case QuantityComparatorGreaterThan:
		return "The actual value is greater than the given value."
	}
	return "<unknown>"
}
