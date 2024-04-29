package karabiner

type BasicParameters map[string]any

type Condition struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type BasicManipulator struct {
	Type         string          `json:"type"`
	From         FromEvent       `json:"from"`
	To           []ToEvent       `json:"to,omitempty"`
	ToIfAlone    []ToEvent       `json:"to_if_alone,omitempty"`
	ToAfterKeyUp []ToEvent       `json:"to_after_key_up,omitempty"`
	Parameters   BasicParameters `json:"parameters,omitempty"`
	Conditions   []Condition     `json:"conditions,omitempty"`
}

func NewBasicManipulator() *BasicManipulator {
	return &BasicManipulator{
		Type:       "basic",
		To:         []ToEvent{},
		ToIfAlone:  []ToEvent{},
		Parameters: make(BasicParameters),
	}
}
