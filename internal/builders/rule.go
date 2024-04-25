package builders

import "github.com/zacharytamas/karabiner-config/internal/karabiner"

type RuleBuilder struct {
	rule *karabiner.Rule
}

func (b *RuleBuilder) Description(description string) *RuleBuilder {
	b.rule.Description = description
	return b
}

func (b *RuleBuilder) AddManipulators(manipulators ...*ManipulatorBuilder) *RuleBuilder {
	for _, m := range manipulators {
		b.rule.Manipulators = append(b.rule.Manipulators, m.Build())
	}
	return b
}

func (b RuleBuilder) Build() karabiner.Rule {
	return *b.rule
}

func NewRule() *RuleBuilder {
	return &RuleBuilder{
		rule: &karabiner.Rule{},
	}
}
