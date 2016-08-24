package cloudsearchhelper

import "fmt"

type (
	// Action is indexing action
	Action int

	// Datum is data of cloudsearch index object
	Datum struct {
		Type   Action      `json:"type"`
		ID     string      `json:"id"`
		Fields interface{} `json:"fields,omitempty"`
	}
)

const (
	Add Action = iota + 1
	Delete
)

// MarshalJSON for json
func (a Action) MarshalJSON() ([]byte, error) {
	switch a {
	case Add:
		return []byte("\"add\""), nil
	case Delete:
		return []byte("\"delete\""), nil
	}

	return nil, fmt.Errorf("Action not found")
}

// UnmarshalJSON for json
func (a *Action) UnmarshalJSON(data []byte) error {
	s := string(data)

	switch s {
	case "\"add\"":
		*a = Add
	case "\"delete\"":
		*a = Delete
	default:
		return fmt.Errorf("action (%s) not found", s)
	}

	return nil
}
