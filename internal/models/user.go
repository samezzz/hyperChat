package models

type UserState struct {
	Stage     int               // Current stage in the conversation
	Responses map[string]string // Stores user responses like age, weight, etc.
}

// NewUserState initializes a new user state
func NewUserState() *UserState {
	return &UserState{
		Stage:     0,
		Responses: make(map[string]string),
	}
}
