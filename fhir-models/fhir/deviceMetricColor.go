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

// DeviceMetricColor is documented here http://hl7.org/fhir/ValueSet/metric-color
type DeviceMetricColor int

const (
	DeviceMetricColorBlack DeviceMetricColor = iota
	DeviceMetricColorRed
	DeviceMetricColorGreen
	DeviceMetricColorYellow
	DeviceMetricColorBlue
	DeviceMetricColorMagenta
	DeviceMetricColorCyan
	DeviceMetricColorWhite
)

func (code DeviceMetricColor) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *DeviceMetricColor) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "black":
		*code = DeviceMetricColorBlack
	case "red":
		*code = DeviceMetricColorRed
	case "green":
		*code = DeviceMetricColorGreen
	case "yellow":
		*code = DeviceMetricColorYellow
	case "blue":
		*code = DeviceMetricColorBlue
	case "magenta":
		*code = DeviceMetricColorMagenta
	case "cyan":
		*code = DeviceMetricColorCyan
	case "white":
		*code = DeviceMetricColorWhite
	default:
		return fmt.Errorf("unknown DeviceMetricColor code `%s`", s)
	}
	return nil
}
func (code DeviceMetricColor) String() string {
	return code.Code()
}
func (code DeviceMetricColor) Code() string {
	switch code {
	case DeviceMetricColorBlack:
		return "black"
	case DeviceMetricColorRed:
		return "red"
	case DeviceMetricColorGreen:
		return "green"
	case DeviceMetricColorYellow:
		return "yellow"
	case DeviceMetricColorBlue:
		return "blue"
	case DeviceMetricColorMagenta:
		return "magenta"
	case DeviceMetricColorCyan:
		return "cyan"
	case DeviceMetricColorWhite:
		return "white"
	}
	return "<unknown>"
}
func (code DeviceMetricColor) Display() string {
	switch code {
	case DeviceMetricColorBlack:
		return "Color Black"
	case DeviceMetricColorRed:
		return "Color Red"
	case DeviceMetricColorGreen:
		return "Color Green"
	case DeviceMetricColorYellow:
		return "Color Yellow"
	case DeviceMetricColorBlue:
		return "Color Blue"
	case DeviceMetricColorMagenta:
		return "Color Magenta"
	case DeviceMetricColorCyan:
		return "Color Cyan"
	case DeviceMetricColorWhite:
		return "Color White"
	}
	return "<unknown>"
}
func (code DeviceMetricColor) Definition() string {
	switch code {
	case DeviceMetricColorBlack:
		return "Color for representation - black."
	case DeviceMetricColorRed:
		return "Color for representation - red."
	case DeviceMetricColorGreen:
		return "Color for representation - green."
	case DeviceMetricColorYellow:
		return "Color for representation - yellow."
	case DeviceMetricColorBlue:
		return "Color for representation - blue."
	case DeviceMetricColorMagenta:
		return "Color for representation - magenta."
	case DeviceMetricColorCyan:
		return "Color for representation - cyan."
	case DeviceMetricColorWhite:
		return "Color for representation - white."
	}
	return "<unknown>"
}
