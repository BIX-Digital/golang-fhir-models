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

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Claim is documented here http://hl7.org/fhir/StructureDefinition/Claim
type Claim struct {
	Id                   *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               FinancialResourceStatusCodes `bson:"status" json:"status"`
	Type                 CodeableConcept              `bson:"type" json:"type"`
	SubType              *CodeableConcept             `bson:"subType,omitempty" json:"subType,omitempty"`
	Use                  Use                          `bson:"use" json:"use"`
	Patient              Reference                    `bson:"patient" json:"patient"`
	BillablePeriod       *Period                      `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Created              string                       `bson:"created" json:"created"`
	Enterer              *Reference                   `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Insurer              *Reference                   `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider             Reference                    `bson:"provider" json:"provider"`
	Priority             CodeableConcept              `bson:"priority" json:"priority"`
	FundsReserve         *CodeableConcept             `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	Related              []ClaimRelated               `bson:"related,omitempty" json:"related,omitempty"`
	Prescription         *Reference                   `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription *Reference                   `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                *ClaimPayee                  `bson:"payee,omitempty" json:"payee,omitempty"`
	Referral             *Reference                   `bson:"referral,omitempty" json:"referral,omitempty"`
	Facility             *Reference                   `bson:"facility,omitempty" json:"facility,omitempty"`
	CareTeam             []ClaimCareTeam              `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	SupportingInfo       []ClaimSupportingInfo        `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Diagnosis            []ClaimDiagnosis             `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure            []ClaimProcedure             `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Insurance            []ClaimInsurance             `bson:"insurance" json:"insurance"`
	Accident             *ClaimAccident               `bson:"accident,omitempty" json:"accident,omitempty"`
	Item                 []ClaimItem                  `bson:"item,omitempty" json:"item,omitempty"`
	Total                *Money                       `bson:"total,omitempty" json:"total,omitempty"`
}
type ClaimRelated struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Claim             *Reference       `bson:"claim,omitempty" json:"claim,omitempty"`
	Relationship      *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference         *Identifier      `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ClaimPayee struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	Party             *Reference      `bson:"party,omitempty" json:"party,omitempty"`
}
type ClaimCareTeam struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence" json:"sequence"`
	Provider          Reference        `bson:"provider" json:"provider"`
	Responsible       *bool            `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Qualification     *CodeableConcept `bson:"qualification,omitempty" json:"qualification,omitempty"`
}
type ClaimSupportingInfo struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence" json:"sequence"`
	Category          CodeableConcept  `bson:"category" json:"category"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	TimingDate        *string          `bson:"timingDate,omitempty" json:"timingDate,omitempty"`
	TimingPeriod      *Period          `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	ValueBoolean      *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueString       *string          `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueQuantity     *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueAttachment   *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueReference    *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}
type ClaimDiagnosis struct {
	Id                       *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                 int               `bson:"sequence" json:"sequence"`
	DiagnosisCodeableConcept CodeableConcept   `bson:"diagnosisCodeableConcept" json:"diagnosisCodeableConcept"`
	DiagnosisReference       Reference         `bson:"diagnosisReference" json:"diagnosisReference"`
	Type                     []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	OnAdmission              *CodeableConcept  `bson:"onAdmission,omitempty" json:"onAdmission,omitempty"`
	PackageCode              *CodeableConcept  `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}
type ClaimProcedure struct {
	Id                       *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                 int               `bson:"sequence" json:"sequence"`
	Type                     []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Date                     *string           `bson:"date,omitempty" json:"date,omitempty"`
	ProcedureCodeableConcept CodeableConcept   `bson:"procedureCodeableConcept" json:"procedureCodeableConcept"`
	ProcedureReference       Reference         `bson:"procedureReference" json:"procedureReference"`
	Udi                      []Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
}
type ClaimInsurance struct {
	Id                  *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence            int         `bson:"sequence" json:"sequence"`
	Focal               bool        `bson:"focal" json:"focal"`
	Identifier          *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Coverage            Reference   `bson:"coverage" json:"coverage"`
	BusinessArrangement *string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	PreAuthRef          []string    `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference  `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}
type ClaimAccident struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date              string           `bson:"date" json:"date"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	LocationAddress   *Address         `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
}
type ClaimItem struct {
	Id                      *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                int               `bson:"sequence" json:"sequence"`
	CareTeamSequence        []int             `bson:"careTeamSequence,omitempty" json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []int             `bson:"diagnosisSequence,omitempty" json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []int             `bson:"procedureSequence,omitempty" json:"procedureSequence,omitempty"`
	InformationSequence     []int             `bson:"informationSequence,omitempty" json:"informationSequence,omitempty"`
	Revenue                 *CodeableConcept  `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category                *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService        CodeableConcept   `bson:"productOrService" json:"productOrService"`
	Modifier                []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate            *string           `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period           `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept  `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address          `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference        `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice               *Money            `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                  *string           `bson:"factor,omitempty" json:"factor,omitempty"`
	Net                     *Money            `bson:"net,omitempty" json:"net,omitempty"`
	Udi                     []Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite                *CodeableConcept  `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Encounter               []Reference       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Detail                  []ClaimItemDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ClaimItemDetail struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                        `bson:"sequence" json:"sequence"`
	Revenue           *CodeableConcept           `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService  CodeableConcept            `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept          `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept          `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                  `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                     `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                     `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference                `bson:"udi,omitempty" json:"udi,omitempty"`
	SubDetail         []ClaimItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ClaimItemDetailSubDetail struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int               `bson:"sequence" json:"sequence"`
	Revenue           *CodeableConcept  `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService  CodeableConcept   `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money            `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string           `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money            `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
}
type OtherClaim Claim

// MarshalJSON marshals the given Claim as JSON into a byte slice
func (r Claim) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherClaim
		ResourceType string `json:"resourceType"`
	}{
		OtherClaim:   OtherClaim(r),
		ResourceType: "Claim",
	})
	return buffer.Bytes(), err
}

// UnmarshalClaim unmarshals a Claim.
func UnmarshalClaim(b []byte) (Claim, error) {
	var claim Claim
	if err := json.Unmarshal(b, &claim); err != nil {
		return claim, err
	}
	return claim, nil
}
