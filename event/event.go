// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package event

import (
	"github.com/vlorc/gioc/types"
	"strings"
	"github.com/vlorc/gioc/invoker"
)

func(el *CoreEventListener) On(event string,fn interface{}) error {
	return el.on(el.provider,event, fn)
}

func(el *CoreEventListener) Emit(event string,params ...interface{}) error{
	return el.EmitWith(el.provider,event,params...)
}

func(el *CoreEventListener) OnWith(provider func() types.Provider,event string,fn interface{}) error {
	return el.on(provider,event, func() types.Invoker{
		return invoker.NewInvokerWith(provider,toInvoker(fn))
	})
}

func(el *CoreEventListener) EmitWith(provider func() types.Provider,event string,params ...interface{}) error {
	for _, v := range el.get(event) {
		v.ApplyWith(provider(),params...)
	}
	return nil
}

func(el *CoreEventListener) on(provider func() types.Provider,event string,fn interface{}) error {
	pos := strings.Index(event,"::")
	if pos > 0 {
		var l types.EventListener
		provider().Assign(&l, event[:pos])
		return l.On(event[pos+2:], fn)
	}
	if 0 == pos {
		event = event[2:]
	}
	el.set(event,toInvoker(fn))
	return nil
}

func(el *CoreEventListener) set(event string,inv types.Invoker) {
	el.lock.Lock()
	defer el.lock.Unlock()

	if "*" == event {
		el.defaults = append(el.defaults,inv)
	} else {
		el.listeners[event] = append(el.listeners[event],inv)
	}
}

func(el *CoreEventListener) get(event string) []types.Invoker{
	el.lock.RLock()
	inv,ok := el.listeners[event]
	el.lock.RUnlock()
	if !ok {
		inv = el.defaults
	}
	return inv
}

func toInvoker(fn interface{}) types.Invoker{
	var inv types.Invoker
	switch r := fn.(type) {
	case types.Invoker:
		inv = r
	case func() types.Invoker:
		inv = r()
	default:
		inv = invoker.NewInvoker(fn,nil)
	}
	return inv
}