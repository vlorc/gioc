// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"github.com/vlorc/gioc/types"
	"go/token"
	"go/scanner"
	"errors"
	"fmt"
)

func NewTagParser() *TagParser{
	obj := &TagParser{
		tagHandle:make(map[string][]TagHandle),
	}

	obj.tagHandle["optional"] = []TagHandle{
		flagsHandle(types.DEPENDENCY_FLAG_OPTIONAL),
	}
	obj.tagHandle["extends"] = []TagHandle{
		flagsHandle(types.DEPENDENCY_FLAG_EXTENDS),
		extendsHandle,
	}
	obj.tagHandle["default"] = []TagHandle{
		flagsHandle(types.DEPENDENCY_FLAG_DEFAULT),
	}

	return obj
}

func extendsHandle(factory types.DependencyFactory,des types.PropertyDescriptor,_ []string) (interface{}, error) {
	dep, err := factory.Instance(des.Type())
	des.SetDepend(dep)
	return dep,err
}

func flagsHandle(flag types.DependencyFlag) TagHandle {
	return func(_ types.DependencyFactory,des types.PropertyDescriptor,_ []string) (interface{},error) {
		des.SetFlags(des.Flags() | flag)
		return flag,nil
	}
}

func (tp *TagParser)Resolve(factory types.DependencyFactory,tag string, des types.PropertyDescriptor) {
	src := []byte(tag)
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		_, tk, str := s.Scan()
		switch tk {
		case token.IDENT:
			tp.Invoke(factory,str,des)
		case token.CHAR,token.STRING:
			des.SetName(str[1 : len(str) - 1])
		case token.EOF:
			return
		}
	}

	return
}

func (tp *TagParser)Invoke(factory types.DependencyFactory,key string, des types.PropertyDescriptor) {
	handle,ok := tp.tagHandle[key]
	if !ok {
		panic(errors.New(fmt.Sprintf("can't find token '%s'",key)))
	}

	for _,f := range handle {
		_,err := f(factory,des,nil)
		if nil != err {
			panic(err)
		}
	}
}

