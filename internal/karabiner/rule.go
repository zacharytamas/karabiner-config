package karabiner

type Rule struct {
	Description  string             `json:"description"`
	Manipulators []BasicManipulator `json:"manipulators"`
}
