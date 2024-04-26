package builders

import "github.com/zacharytamas/karabiner-config/internal/karabiner"

type ManipulatorBuilder struct {
	manipulator *karabiner.BasicManipulator
}

func (b *ManipulatorBuilder) From(from *FromEventBuilder) *ManipulatorBuilder {
	b.manipulator.From = from.Build()
	return b
}

func (b *ManipulatorBuilder) To(to ...*ToEventBuilder) *ManipulatorBuilder {
	for _, t := range to {
		b.manipulator.To = append(b.manipulator.To, t.Build())
	}
	return b
}

func (b *ManipulatorBuilder) ToIfAlone(to ...*ToEventBuilder) *ManipulatorBuilder {
	for _, t := range to {
		b.manipulator.ToIfAlone = append(b.manipulator.ToIfAlone, t.Build())
	}
	return b
}

func (b *ManipulatorBuilder) ToAfterKeyUp(to ...*ToEventBuilder) *ManipulatorBuilder {
	for _, t := range to {
		b.manipulator.ToAfterKeyUp = append(b.manipulator.ToAfterKeyUp, t.Build())
	}
	return b
}

func (b *ManipulatorBuilder) Parameters(parameters karabiner.BasicParameters) *ManipulatorBuilder {
	b.manipulator.Parameters = parameters
	return b
}

func (b *ManipulatorBuilder) Build() karabiner.BasicManipulator {
	return *b.manipulator
}

func NewManipulator() *ManipulatorBuilder {
	return &ManipulatorBuilder{
		manipulator: &karabiner.BasicManipulator{
			Type:       "basic",
			To:         []karabiner.ToEvent{},
			ToIfAlone:  []karabiner.ToEvent{},
			Parameters: make(karabiner.BasicParameters),
		},
	}
}
