// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package operation

import (
	"github.com/vlorc/gioc/module"
)

type EventHandle func(*EventContext)

func Event(handle ...EventHandle) module.ModuleInitHandle {
	return EventScore("",handle...)
}

func EventScore(name string,handle ...EventHandle) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		c := &EventContext{ctx,nil}
		c.Parent().AsProvider().Assign(&c.Listener,name)
		for _,v := range handle {
			v(c)
		}
	}
}

func On(event string,fn interface{}) EventHandle{
	return func(ctx *EventContext) {
		ctx.Listener.On(event,fn)
	}
}

func OnWith(event string,fn interface{}) EventHandle{
	return func(ctx *EventContext) {
		ctx.Listener.OnWith(lazyProvider(ctx.Container),event,fn)
	}
}

func Emit(event string,args ...interface{}) EventHandle{
	return func(ctx *EventContext) {
		ctx.Listener.Emit(event,args...)
	}
}

func EmitWith(event string,args ...interface{}) EventHandle{
	return func(ctx *EventContext) {
		ctx.Listener.EmitWith(lazyProvider(ctx.Container),event,args...)
	}
}
