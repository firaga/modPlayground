// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"io"
	"os"
)

// InitializeEvent 声明injector的函数签名
func InitializeEvent(msg string) Event {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	wire.Bind()
	wire.FieldsOf()
	panic(wire.Build(NewGreeter, mySet))
	return Event{} //返回值没有实际意义，只需符合函数签名即可
}

var mySet = wire.NewSet(NewEvent, NewMessage)
