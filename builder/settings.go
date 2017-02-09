package builder

// An abstracted settings provider that can unmarshal to a target
type SettingsProvider interface {
	AssignSettings(target interface{}) error
}
