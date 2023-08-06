//go:build wireinject
// +build wireinject

package application

import (
	"github.com/google/wire"
)

// Init initializes the App by using the dependency injection framework, Wire.
// It builds the registry of dependencies required by the App and panics if the build fails.
func Init() (*App, func()) {
	panic(wire.Build(registry))
}
