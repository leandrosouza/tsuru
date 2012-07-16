package bind

import (
	"github.com/timeredbull/tsuru/api/unit"
)

// AppContainer provides methdos for a container of apps.
//
// The container stores only the names of the apps.
type AppContainer interface {
	// Adds an app to the container.
	AddApp(string) error

	// Finds an app in the container, returning an index a value >= 0 if it is
	// present, and -1 if not present.
	FindApp(string) int

	// Removes an app form the container.
	RemoveApp(name string) error
}

type EnvVar interface {
	Name() string
	Value() string
	Public() bool
	InstanceName() string
}

type App interface {
	Units() []unit.Unit
	SetEnvs([]EnvVar) error
	UnsetEnvs([]string) error
}

type Binder interface {
	Bind(AppContainer, App) error
	Unbind(AppContainer, App) error
}
