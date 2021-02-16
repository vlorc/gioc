package operation

import (
	"github.com/vlorc/gioc/module"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
	"reflect"
	"strings"
)

func Condition(cond module.ModuleCondHandle, handle ...module.ModuleInitHandle) module.ModuleInitHandle {
	return func(ctx *module.ModuleInitContext) {
		if cond(ctx) {
			for _, v := range handle {
				v(ctx)
			}
		}
	}
}

func havingValue(eq func(types.BeanFactory, types.Provider) bool, typ reflect.Type, names ...string) module.ModuleCondHandle {
	return func(ctx *module.ModuleInitContext) bool {
		if len(names) == 0 {
			if f := ctx.Container().AsProvider().Factory(typ, "", -1); nil != f {
				return eq(f, ctx.Container().AsProvider())
			}
			return false
		}
		for _, v := range names {
			if strings.HasPrefix(v, "${") && strings.HasSuffix(v, "") {
				if str, err := types.NewNameFactory(v).Instance(ctx.Container().AsProvider()); nil != err {
					// add check error
					continue
				} else {
					v = str
				}
			}
			if f := ctx.Container().AsProvider().Factory(typ, v, -1); nil != f && eq(f, ctx.Container().AsProvider()) {
				return true
			}
		}

		return false
	}
}

func HavingBean(impType interface{}, names ...string) module.ModuleCondHandle {
	return havingValue(func(factory types.BeanFactory, provider types.Provider) bool {
		return nil != factory
	}, utils.TypeOf(impType), names...)
}

func HavingValue(eq func(types.BeanFactory, types.Provider) bool, impType interface{}, names ...string) module.ModuleCondHandle {
	return havingValue(eq, utils.TypeOf(impType), names...)
}

func Not(cond ...module.ModuleCondHandle) module.ModuleCondHandle {
	return func(c *module.ModuleInitContext) bool {
		for _, v := range cond {
			if v(c) {
				return false
			}
		}
		return true
	}
}

func And(cond ...module.ModuleCondHandle) module.ModuleCondHandle {
	return func(c *module.ModuleInitContext) bool {
		for _, v := range cond {
			if !v(c) {
				return false
			}
		}
		return true
	}
}

func Or(cond ...module.ModuleCondHandle) module.ModuleCondHandle {
	return func(c *module.ModuleInitContext) bool {
		for _, v := range cond {
			if v(c) {
				return true
			}
		}
		return false
	}
}

func Equal(val interface{}) func(types.BeanFactory, types.Provider) bool {
	return func(factory types.BeanFactory, provider types.Provider) bool {
		instance, err := factory.Instance(provider)
		if nil != err {
			return false
		}
		return reflect.DeepEqual(val, instance)
	}
}
