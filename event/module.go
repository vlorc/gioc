package event

import (
	. "github.com/vlorc/gioc/module"
	. "github.com/vlorc/gioc/module/operation"
	"github.com/vlorc/gioc/types"
)

func EventModuleFor(id ...string) types.ModuleFactory{
	return NewModuleForFactory(
		Declare(Method(NewEventListenerWith),Singleton(),Id("",id...)),
	)
}