package builders

import "github.com/zacharytamas/karabiner-config/internal/karabiner"

type FromEventBuilder struct {
	fromEvent karabiner.FromEvent
}

func (b *FromEventBuilder) KeyCode(key karabiner.KeyCode) *FromEventBuilder {
	b.fromEvent.KeyCode = key
	return b
}

func (b *FromEventBuilder) Modifiers(modifiers ...karabiner.Modifier) *FromEventBuilder {
	b.fromEvent.Modifiers.Optional = append(b.fromEvent.Modifiers.Optional, modifiers...)
	return b
}

func (b *FromEventBuilder) MandatoryModifiers(modifiers ...karabiner.Modifier) *FromEventBuilder {
	b.fromEvent.Modifiers.Mandatory = append(b.fromEvent.Modifiers.Mandatory, modifiers...)
	return b
}

func (b FromEventBuilder) Build() karabiner.FromEvent {
	if len(b.fromEvent.Modifiers.Mandatory) == 0 && len(b.fromEvent.Modifiers.Optional) == 0 {
		b.fromEvent.Modifiers = nil
	}
	return b.fromEvent
}

func NewFromEvent() *FromEventBuilder {
	return &FromEventBuilder{
		fromEvent: karabiner.FromEvent{Modifiers: &karabiner.FromModifiers{}},
	}
}

type ToEventBuilder struct {
	toEvent karabiner.ToEvent
}

func (b *ToEventBuilder) KeyCode(key karabiner.KeyCode) *ToEventBuilder {
	b.toEvent.KeyCode = key
	return b
}

func (b *ToEventBuilder) Modifiers(modifiers ...karabiner.Modifier) *ToEventBuilder {
	b.toEvent.Modifiers = append(b.toEvent.Modifiers, modifiers...)
	return b
}

func (b *ToEventBuilder) SetVariable(name string, value interface{}) *ToEventBuilder {
	switch value.(type) {
	case string, bool, int, float32, float64:
		b.toEvent.SetVariable = &karabiner.ToVariable{Name: name, Value: value}
	default:
		panic("Unsupported type")
	}
	return b
}

func (b *ToEventBuilder) Build() karabiner.ToEvent {
	return b.toEvent
}

func NewToEvent() *ToEventBuilder {
	return &ToEventBuilder{
		toEvent: karabiner.ToEvent{},
	}
}
