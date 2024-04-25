package karabiner

type BasicParameters map[string]any

type BasicManipulator struct {
	Type       string          `json:"type"`
	From       FromEvent       `json:"from"`
	To         []ToEvent       `json:"to,omitempty"`
	ToIfAlone  []ToEvent       `json:"to_if_alone,omitempty"`
	Parameters BasicParameters `json:"parameters,omitempty"`
}

func NewBasicManipulator() *BasicManipulator {
	return &BasicManipulator{
		Type:       "basic",
		To:         []ToEvent{},
		ToIfAlone:  []ToEvent{},
		Parameters: make(BasicParameters),
	}
}
