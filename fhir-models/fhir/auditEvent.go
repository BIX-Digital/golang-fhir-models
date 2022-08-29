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
)

// THIS FILE IS GENERATED BY https://github.com/bix-digital/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// AuditEvent is documented here http://hl7.org/fhir/StructureDefinition/AuditEvent
type AuditEvent struct {
	ID                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              Coding             `bson:"type" json:"type"`
	Subtype           []Coding           `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Action            *AuditEventAction  `bson:"action,omitempty" json:"action,omitempty"`
	Period            *Period            `bson:"period,omitempty" json:"period,omitempty"`
	Recorded          string             `bson:"recorded" json:"recorded"`
	Outcome           *AuditEventOutcome `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OutcomeDesc       *string            `bson:"outcomeDesc,omitempty" json:"outcomeDesc,omitempty"`
	PurposeOfEvent    []CodeableConcept  `bson:"purposeOfEvent,omitempty" json:"purposeOfEvent,omitempty"`
	Agent             []AuditEventAgent  `bson:"agent" json:"agent"`
	Source            AuditEventSource   `bson:"source" json:"source"`
	Entity            []AuditEventEntity `bson:"entity,omitempty" json:"entity,omitempty"`
}
type AuditEventAgent struct {
	ID                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept        `bson:"type,omitempty" json:"type,omitempty"`
	Role              []CodeableConcept       `bson:"role,omitempty" json:"role,omitempty"`
	Who               *Reference              `bson:"who,omitempty" json:"who,omitempty"`
	AltId             *string                 `bson:"altId,omitempty" json:"altId,omitempty"`
	Name              *string                 `bson:"name,omitempty" json:"name,omitempty"`
	Requestor         bool                    `bson:"requestor" json:"requestor"`
	Location          *Reference              `bson:"location,omitempty" json:"location,omitempty"`
	Policy            []string                `bson:"policy,omitempty" json:"policy,omitempty"`
	Media             *Coding                 `bson:"media,omitempty" json:"media,omitempty"`
	Network           *AuditEventAgentNetwork `bson:"network,omitempty" json:"network,omitempty"`
	PurposeOfUse      []CodeableConcept       `bson:"purposeOfUse,omitempty" json:"purposeOfUse,omitempty"`
}
type AuditEventAgentNetwork struct {
	ID                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Address           *string                     `bson:"address,omitempty" json:"address,omitempty"`
	Type              *AuditEventAgentNetworkType `bson:"type,omitempty" json:"type,omitempty"`
}
type AuditEventSource struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Site              *string     `bson:"site,omitempty" json:"site,omitempty"`
	Observer          Reference   `bson:"observer" json:"observer"`
	Type              []Coding    `bson:"type,omitempty" json:"type,omitempty"`
}
type AuditEventEntity struct {
	ID                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	What              *Reference               `bson:"what,omitempty" json:"what,omitempty"`
	Type              *Coding                  `bson:"type,omitempty" json:"type,omitempty"`
	Role              *Coding                  `bson:"role,omitempty" json:"role,omitempty"`
	Lifecycle         *Coding                  `bson:"lifecycle,omitempty" json:"lifecycle,omitempty"`
	SecurityLabel     []Coding                 `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Name              *string                  `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                  `bson:"description,omitempty" json:"description,omitempty"`
	Query             *string                  `bson:"query,omitempty" json:"query,omitempty"`
	Detail            []AuditEventEntityDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}
type AuditEventEntityDetail struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string      `bson:"type" json:"type"`
	ValueString       string      `bson:"valueString" json:"valueString"`
	ValueBase64Binary string      `bson:"valueBase64Binary" json:"valueBase64Binary"`
}
type OtherAuditEvent AuditEvent

// MarshalJSON marshals the given AuditEvent as JSON into a byte slice
func (r AuditEvent) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherAuditEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAuditEvent: OtherAuditEvent(r),
		ResourceType:    "AuditEvent",
	})
	return buffer.Bytes(), err
}

// UnmarshalAuditEvent unmarshals a AuditEvent.
func UnmarshalAuditEvent(b []byte) (AuditEvent, error) {
	var auditEvent AuditEvent
	if err := json.Unmarshal(b, &auditEvent); err != nil {
		return auditEvent, err
	}
	return auditEvent, nil
}
