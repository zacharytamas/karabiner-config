package app

import (
	"encoding/json"
	"fmt"

	"github.com/zacharytamas/karabiner-config/internal/builders"
	"github.com/zacharytamas/karabiner-config/internal/karabiner"
	"github.com/zacharytamas/karabiner-config/internal/layers"
)

func Run() {

	hyperRule := builders.
		NewRule().
		Description("Caps Lock → Hyper").
		AddManipulators(builders.NewManipulator().
			From(builders.NewFromEvent().KeyCode("caps_lock").Modifiers("shift", "command")).
			Parameters(map[string]any{"basic.to_if_alone_timeout_milliseconds": 250}).
			To(builders.NewToEvent().
				KeyCode("left_command").
				Modifiers("option", "control", "shift")).
			ToIfAlone(builders.NewToEvent().KeyCode("escape")),
		).
		Build()
	hyperLayer := layers.NewHyperLayer("o", "hyper-o").Build()

	serialized, _ := json.MarshalIndent([]karabiner.Rule{hyperRule, hyperLayer}, "", "  ")

	fmt.Println(string(serialized))
}
