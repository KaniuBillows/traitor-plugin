package traitor_plugin

type Executable interface {
	Require() map[string]any
	GetName() string
}
