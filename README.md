# Golang FHIR Models

**This is a fork of https://github.com/samply/golang-fhir-models with support for polymorphic types and several fixes.**

This repository contains a FHIR® R4 models for Go. The models consist of Go structs for each resource and data type. The structs are suitable for JSON encoding/decoding. 

## Features

* resources implement the [Marshaler][1] interface
* unmarshal functions are provided for every resource
* enums are provided for every ValueSet used in a [required binding][2], has a computer friendly name and refers only to one CodeSystem
* enums implement `Code()`, `Display()` and `Definition()` methods

## Usage

In your project, import `github.com/bix-digital/golang-fhir-models/fhir-models/fhir` and you are done.

## Develop

This repository contains two Go modules, the generated models itself and the generator. Both modules use `go generate` to generate the FHIR models and are executed via `./gen-resources.sh` under `fhir-models-gen` and `fhir-models`.

## License

Copyright 2019 The Samply Development Community

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

[1]: <https://golang.org/pkg/encoding/json/#Marshaler>
[2]: <https://www.hl7.org/fhir/terminologies.html#strength>
