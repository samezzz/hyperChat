package models

type UserState struct {
	Onboarding            bool
	Stage                 int // Current stage in the conversation
	Age                   string
	Weight                string
	BloodPressure         string
	Medications           string
	PreExistingConditions string
	LanguageStage         int
	LanguageSelected      string

	FeatureSelected bool
	CurrentFeature  string
	BPLogStage      int
	ReminderStage   int
}

// NewUserState initializes a new user state
func NewUserState() *UserState {
	return &UserState{
		Stage:                 0,
		Onboarding:            false,
		Age:                   "",
		Weight:                "",
		BloodPressure:         "",
		Medications:           "",
		PreExistingConditions: "",
	}
}
