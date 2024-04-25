package karabiner

// TODO Narrow Modifier
type Modifier string

type FromModifiers struct {
	Mandatory []Modifier `json:"mandatory,omitempty"`
	Optional  []Modifier `json:"optional,omitempty"`
}

type ToModifiers struct {
	Mandatory []Modifier `json:"mandatory,omitempty"`
	Optional  []Modifier `json:"optional,omitempty"`
}

type FromEvent struct {
	KeyCode   KeyCode        `json:"key_code"`
	Modifiers *FromModifiers `json:"modifiers,omitempty"`
}

type ToEvent struct {
	KeyCode   KeyCode    `json:"key_code"`
	Modifiers []Modifier `json:"modifiers,omitempty"`
}
