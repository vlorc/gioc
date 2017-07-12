package invoker

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

func NewInvokerFactory() types.InvokerFactory {
	return &CoreInvokerFactory{}
}

func (fi *CoreInvokerFactory) Instance(method interface{}, build types.Builder) (invoker types.Invoker, err error) {
	defer utils.Recover(&err)

	invoker = NewInvoker(method, build)
	return
}
