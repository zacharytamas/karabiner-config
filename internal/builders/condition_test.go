package builders

import (
	"testing"
)

func TestConditionBuilder(t *testing.T) {
	builder := NewCondition().Name("test").Type("variable_if").Value(1)

	built := builder.Build()

	if built.Name != "test" {
		t.Errorf("Expected condition name to be 'test', got %s", built.Name)
	}
	if built.Type != "variable_if" {
		t.Errorf("Expected condition type to be 'variable_if', got %s", built.Type)
	}
	if built.Value != 1 {
		t.Errorf("Expected condition value to be 1, got %v", built.Value)
	}
}
