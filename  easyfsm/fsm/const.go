package fsm

import "context"
import "errors"

type (
	BusinessName string

	State int64

	ExecHandler func(context.Context, interface{}) (interface{}, error)
)

var ParamConvertInvalidError = errors.New("param convert invalid")
