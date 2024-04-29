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

func (b *HyperLayerBuilder) AddManipulators(manipulators ...*builders.ManipulatorBuilder) *HyperLayerBuilder {
	for _, m := range manipulators {
		m.Conditions(builders.NewCondition().Name(b.varName).Type("variable_if").Value(1))
		b.ruleBuilder.AddManipulators(m)
	}

	return b
}

func (b *HyperLayerBuilder) AddAppKey(key karabiner.KeyCode, appName string) *HyperLayerBuilder {
	b.AddManipulators(
		builders.NewManipulator().
			From(builders.NewFromEvent().KeyCode(key).MandatoryModifiers("any")).
			To(builders.NewToEvent().ShellCommand(fmt.Sprintf("open -a %q.app", appName))),
	)
	return b
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
				KeyCode(key).
				MandatoryModifiers("command", "option", "control", "shift"),
			).
			To(builders.NewToEvent().SetVariable(varName, 1)).
			ToAfterKeyUp(builders.NewToEvent().SetVariable(varName, 0)).
			ToIfAlone(builders.NewToEvent().KeyCode(key)),
	)
	return &builder
}
