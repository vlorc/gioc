package depend

import (
	"github.com/vlorc/gioc/types"
	"go/token"
	"go/scanner"
)

func makeFlagHandle(flag types.DependencyFlag) TagHandle {
	return func(_ types.DependencyFactory,des types.PropertyDescriptor,_ []string) (interface{},error) {
		des.SetFlags(des.Flags() | flag)
		return flag,nil
	}
}

func (df *CoreDependencyFactory)resolveTag(tag string, des types.PropertyDescriptor) {
	src := []byte(tag)
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		_, tk, str := s.Scan()
		switch tk {
		case token.IDENT:
			df.callIdent(str,des)
		case token.CHAR,token.STRING:
			des.SetName(str)
		case token.EOF:
			return
		}
	}

	return
}

func (df *CoreDependencyFactory)callIdent(ident string, des types.PropertyDescriptor) {
	handle,ok := df.tagHandle[ident]
	if !ok {
		return
	}
	for _,f := range handle {
		f(df,des,nil)
	}
}

