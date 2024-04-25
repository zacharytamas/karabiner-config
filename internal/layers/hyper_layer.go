package layers

import (
	"fmt"

	"github.com/zacharytamas/karabiner-config/internal/builders"
	"github.com/zacharytamas/karabiner-config/internal/karabiner"
)

type HyperLayerBuilder struct {
	key         karabiner.KeyCode
	varName     string
	ruleBuilder *builders.RuleBuilder
}

func (b *HyperLayerBuilder) Build() karabiner.Rule {
	return b.ruleBuilder.Build()
}

func NewHyperLayer(key karabiner.KeyCode, varName string) *HyperLayerBuilder {
	builder := HyperLayerBuilder{
		key:         key,
		varName:     varName,
		ruleBuilder: builders.NewRule().Description(fmt.Sprintf("Layer - %s", varName)),
	}

	// Activation rule
	builder.ruleBuilder.AddManipulators(
		builders.NewManipulator().
			From(builders.NewFromEvent().
				KeyCode("o").
				MandatoryModifiers("command", "control", "shift", "option"),
			),
	)
	return &builder
}
