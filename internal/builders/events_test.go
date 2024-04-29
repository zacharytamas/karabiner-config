package builders

import (
	"reflect"
	"testing"

	"github.com/zacharytamas/karabiner-config/internal/karabiner"
)

func TestFromEventBuilder(t *testing.T) {
	builder := NewFromEvent().KeyCode("a").Modifiers("shift", "command")

	built := builder.Build()
	modifiers := built.Modifiers

	expectedModifiers := &karabiner.FromModifiers{
		Mandatory: nil,
		Optional:  []karabiner.Modifier{"shift", "command"},
	}

	if built.KeyCode != "a" {
		t.Errorf("Expected key code to be 'a', got %s", built.KeyCode)
	}

	if !reflect.DeepEqual(modifiers, expectedModifiers) {
		t.Errorf("Expected modifiers to be %v, got %v", expectedModifiers, modifiers)
	}
}

func TestToEventBuilder(t *testing.T) {
	builder := NewToEvent().KeyCode("a").Modifiers("shift", "command")

	built := builder.Build()
	expectedModifiers := []karabiner.Modifier{"shift", "command"}

	if built.KeyCode != "a" {
		t.Errorf("Expected key code to be 'a', got %s", built.KeyCode)
	}

	if !reflect.DeepEqual(built.Modifiers, expectedModifiers) {
		t.Errorf("Expected modifiers to be %v, got %v", expectedModifiers, built.Modifiers)
	}

}
