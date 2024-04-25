package rules

import (
	"github.com/zacharytamas/karabiner-config/internal/builders"
	"github.com/zacharytamas/karabiner-config/internal/karabiner"
)

func Map(key karabiner.KeyCode) *builders.ManipulatorBuilder {
	return builders.NewManipulator().From(builders.NewFromEvent().KeyCode(key))
}
