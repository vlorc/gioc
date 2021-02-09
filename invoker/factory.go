package invoker

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
)

func NewInvokerFactory() types.InvokerFactory {
	return &CoreInvokerFactory{}
}

func (fi *CoreInvokerFactory) Instance(method interface{}, dependency types.Dependency) (invoker types.Invoker, err error) {
	defer utils.Recover(&err)

	invoker = NewInvoker(method, dependency)
	return
}

func NewNoArgsInvoker(method interface{}) types.Invoker {
	if src := utils.ValueOf(method); reflect.Func == src.Kind() {
		return NoParamInvoker(src)
	}
	panic(types.NewWithError(types.ErrTypeNotFunction, method))
}

func NewSimpleInvoker(method interface{}) types.Invoker {
	if src := utils.ValueOf(method); reflect.Func == src.Kind() {
		return SimpleInvoker(src)
	}
	panic(types.NewWithError(types.ErrTypeNotFunction, method))
}

func NewInvokerWith(provider func() types.Provider, invoker types.Invoker) types.Invoker {
	return &WithInvoker{
		provider: provider,
		invoker:  invoker,
	}
}
