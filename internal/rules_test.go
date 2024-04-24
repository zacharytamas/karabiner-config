package rules

import (
	"testing"
)

func TestRule(t *testing.T) {
	r := Rule{
		Description: "Test rule",
		Manipulators: []Manipulator{
			Manipulator{},
		},
	}

	var _ = r

}
