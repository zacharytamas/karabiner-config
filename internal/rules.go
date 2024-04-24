package rules

type Manipulator struct {
	Type string `json:"type"`
	From From   `json:"from"`
}

type ManipulatorBuilder struct {
	manipulator Manipulator
}

func (b *ManipulatorBuilder) Build() Manipulator {
	return b.manipulator
}

type FromKeyParam string

type From struct {
	KeyCode FromKeyParam `json:"key_code"`
}

func Map(key FromKeyParam) ManipulatorBuilder {
	return ManipulatorBuilder{
		manipulator: Manipulator{
			Type: "basic",
			From: From{KeyCode: key},
		},
	}
}

type Rule struct {
	Description  string        `json:"description"`
	Manipulators []Manipulator `json:"manipulators"`
}

type RuleBuilder struct {
	Description  string
	manipulators []ManipulatorBuilder
}

func (b *RuleBuilder) Manipulators(manipulators ...ManipulatorBuilder) *RuleBuilder {
	b.manipulators = append(b.manipulators, manipulators...)
	return b
}

func (b RuleBuilder) Build() Rule {
	manipulators := make([]Manipulator, len(b.manipulators))
	for i, m := range b.manipulators {
		manipulators[i] = m.Build()
	}

	return Rule{
		Description:  b.Description,
		Manipulators: manipulators,
	}
}

func NewBuilder(description string) *RuleBuilder {
	return &RuleBuilder{
		Description:  description,
		manipulators: []ManipulatorBuilder{},
	}
}
