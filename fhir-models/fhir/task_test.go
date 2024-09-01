package fhir

import (
	"encoding/json"
	"testing"
)

func TestTaskInput_ValuePolymorphism(t *testing.T) {
	// Expected JSON:
	// {
	//	 "type": {},
	//	 "valueInteger": 5
	// }

	inputValue := 5
	taskInput := TaskInput{ValueInteger: &inputValue}
	data, _ := json.MarshalIndent(taskInput, "", "  ")
	asMap := map[string]interface{}{}
	_ = json.Unmarshal(data, &asMap)
	if len(asMap) != 2 {
		t.Errorf("expected 2 fields, got %d", len(asMap))
	}
	if asMap["valueInteger"] != inputValue {
		t.Errorf("expected %d, got %v", inputValue, asMap["valueInteger"])
	}
}
