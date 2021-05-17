// +build wireinject

package main

import "github.com/google/wire"

func NewApp() (*App, func(), error) {
	panic(wire.Build(
		provideFile,
		provideNetConn,
		wire.Struct(new(App), "*"),
	))
}
