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

// EligibilityRequestPurpose is documented here http://hl7.org/fhir/ValueSet/eligibilityrequest-purpose
type EligibilityRequestPurpose int

const (
	EligibilityRequestPurposeAuthRequirements EligibilityRequestPurpose = iota
	EligibilityRequestPurposeBenefits
	EligibilityRequestPurposeDiscovery
	EligibilityRequestPurposeValidation
)

func (code EligibilityRequestPurpose) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *EligibilityRequestPurpose) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "auth-requirements":
		*code = EligibilityRequestPurposeAuthRequirements
	case "benefits":
		*code = EligibilityRequestPurposeBenefits
	case "discovery":
		*code = EligibilityRequestPurposeDiscovery
	case "validation":
		*code = EligibilityRequestPurposeValidation
	default:
		return fmt.Errorf("unknown EligibilityRequestPurpose code `%s`", s)
	}
	return nil
}
func (code EligibilityRequestPurpose) String() string {
	return code.Code()
}
func (code EligibilityRequestPurpose) Code() string {
	switch code {
	case EligibilityRequestPurposeAuthRequirements:
		return "auth-requirements"
	case EligibilityRequestPurposeBenefits:
		return "benefits"
	case EligibilityRequestPurposeDiscovery:
		return "discovery"
	case EligibilityRequestPurposeValidation:
		return "validation"
	}
	return "<unknown>"
}
func (code EligibilityRequestPurpose) Display() string {
	switch code {
	case EligibilityRequestPurposeAuthRequirements:
		return "Coverage auth-requirements"
	case EligibilityRequestPurposeBenefits:
		return "Coverage benefits"
	case EligibilityRequestPurposeDiscovery:
		return "Coverage Discovery"
	case EligibilityRequestPurposeValidation:
		return "Coverage Validation"
	}
	return "<unknown>"
}
func (code EligibilityRequestPurpose) Definition() string {
	switch code {
	case EligibilityRequestPurposeAuthRequirements:
		return "The prior authorization requirements for the listed, or discovered if specified, converages for the categories of service and/or specifed biling codes are requested."
	case EligibilityRequestPurposeBenefits:
		return "The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested."
	case EligibilityRequestPurposeDiscovery:
		return "The insurer is requested to report on any coverages which they are aware of in addition to any specifed."
	case EligibilityRequestPurposeValidation:
		return "A check that the specified coverages are in-force is requested."
	}
	return "<unknown>"
}
