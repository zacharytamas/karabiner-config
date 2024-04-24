package app

import (
	"encoding/json"
	"fmt"

	rules "github.com/zacharytamas/karabiner-config/internal"
)

func Run() {

	hyperRule := rules.
		NewBuilder("Caps Lock → Hyper").
		Manipulators(rules.Map("caps_lock")).
		Build()

	serialized, _ := json.MarshalIndent([]rules.Rule{hyperRule}, "", "  ")

	fmt.Println(string(serialized))
}
