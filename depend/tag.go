// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"fmt"
	"reflect"
	"github.com/vlorc/gioc/utils"
	"strings"
)

func NewTagParser() *TagParser{
	obj := &TagParser{}

	obj.tagHandle = map[string][]TagHandle{
		"optional":[]TagHandle{
			flagsHandle(types.DEPENDENCY_FLAG_OPTIONAL),
		},
		"extends":[]TagHandle{
			flagsHandle(types.DEPENDENCY_FLAG_EXTENDS),
			extendsHandle,
		},
		"default":[]TagHandle{
			flagsHandle(types.DEPENDENCY_FLAG_DEFAULT),
			defaultHandle,
		},
		"id":[]TagHandle{nameHandle},
		"name":[]TagHandle{nameHandle},
		"lower":[]TagHandle{lowerCaseHandle},
		"upper":[]TagHandle{upperCaseHandle},
	}
	return obj
}

func lowerCaseHandle(factory types.DependencyFactory,des types.Descriptor,_ []string) error {
	des.SetName(strings.ToLower(des.Name()))
	return nil
}

func upperCaseHandle(factory types.DependencyFactory,des types.Descriptor,_ []string) error {
	des.SetName(strings.ToUpper(des.Name()))
	return nil
}


func nameHandle(factory types.DependencyFactory,des types.Descriptor,param []string) error {
	des.SetName(param[0])
	return nil
}

func defaultHandle(factory types.DependencyFactory,des types.Descriptor,_ []string) error {
	val := reflect.Zero(des.Type())
	des.SetDefault(val)
	return nil
}

func extendsHandle(factory types.DependencyFactory,des types.Descriptor,_ []string) error {
	dep, err := factory.Instance(des.Type())
	des.SetDepend(dep)
	return err
}

func flagsHandle(flag types.DependencyFlag) TagHandle {
	return func(_ types.DependencyFactory,des types.Descriptor,_ []string) error {
		des.SetFlags(des.Flags() | flag)
		return nil
	}
}

func (tp *TagParser)Resolve(factory types.DependencyFactory,tag string, des types.Descriptor) {
	s := utils.NewTokenScan()
	s.Init(tag)

	for {
		token, offset, length := s.Scan()
		switch token {
		case utils.TOKEN_IDENT:
			tp.Invoke(factory,tag[offset:length],des)
		case utils.TOKEN_CHART,utils.TOKEN_STRING:
			des.SetName(tag[offset + 1 :length - 1])
		case utils.TOKEN_EOF:
			return
		}
	}

	return
}

func (tp *TagParser)Invoke(factory types.DependencyFactory,key string, des types.Descriptor) {
	handle,ok := tp.tagHandle[key]
	if !ok {
		panic(fmt.Errorf("can't find token '%s'",key))
	}

	for _,f := range handle {
		if err := f(factory,des,nil);nil != err {
			panic(err)
		}
	}
}

