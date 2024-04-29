package builders

import "github.com/zacharytamas/karabiner-config/internal/karabiner"

type ConditionBuilder struct {
	condition *karabiner.Condition
}

func (b *ConditionBuilder) Name(name string) *ConditionBuilder {
	b.condition.Name = name
	return b
}

func (b *ConditionBuilder) Type(typeName string) *ConditionBuilder {
	b.condition.Type = typeName
	return b
}

func (b *ConditionBuilder) Value(value interface{}) *ConditionBuilder {
	b.condition.Value = value
	return b
}

func (b ConditionBuilder) Build() karabiner.Condition {
	return *b.condition
}

func NewCondition() *ConditionBuilder {
	return &ConditionBuilder{
		condition: &karabiner.Condition{},
	}
}
