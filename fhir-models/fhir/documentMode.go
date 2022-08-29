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

// DocumentMode is documented here http://hl7.org/fhir/ValueSet/document-mode
type DocumentMode int

const (
	DocumentModeProducer DocumentMode = iota
	DocumentModeConsumer
)

func (code DocumentMode) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *DocumentMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "producer":
		*code = DocumentModeProducer
	case "consumer":
		*code = DocumentModeConsumer
	default:
		return fmt.Errorf("unknown DocumentMode code `%s`", s)
	}
	return nil
}
func (code DocumentMode) String() string {
	return code.Code()
}
func (code DocumentMode) Code() string {
	switch code {
	case DocumentModeProducer:
		return "producer"
	case DocumentModeConsumer:
		return "consumer"
	}
	return "<unknown>"
}
func (code DocumentMode) Display() string {
	switch code {
	case DocumentModeProducer:
		return "Producer"
	case DocumentModeConsumer:
		return "Consumer"
	}
	return "<unknown>"
}
func (code DocumentMode) Definition() string {
	switch code {
	case DocumentModeProducer:
		return "The application produces documents of the specified type."
	case DocumentModeConsumer:
		return "The application consumes documents of the specified type."
	}
	return "<unknown>"
}
