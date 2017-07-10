// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package depend

import (
	"testing"
	"reflect"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

type depTable map[int]struct {
	des *types.DependencyDescription
	sub depTable
}

func Test_NewDependencyFactory(t *testing.T) {
	test_factory(t,NewDependencyFactory())
}

func test_factory(t *testing.T,factory types.DependencyFactory) {
	if nil == factory{
		t.Errorf("can't allocate a DependencyFactory")
	}
}

func test_struct(t *testing.T,typ reflect.Type,table depTable) {
	factory := NewDependencyFactory()
	test_factory(t,factory)

	dep,err := factory.Instance(typ)
	if nil != err{
		t.Errorf("can't instance dependency error: %s",err.Error())
	}
	if typ != dep.Type() {
		t.Errorf("can't matching dependency type: %s",dep.Type().String())
	}
	if typ.NumField() != dep.Length(){
		t.Errorf("can't matching dependency length: %d",dep.Length())
	}
	compare_Dependency(t,dep,table)
}

func Test_AnonymousStruct(t *testing.T) {
	bean := struct {
		Id int64
		Name *string `inject:"'alias' optional"`
		Identity struct {
			Username string `inject:"lower"`
			Password string `inject:"upper"`
		}
	}{}
	typ := utils.DirectlyType(reflect.TypeOf(bean))
	table := depTable{
		0: {
			des :&types.DependencyDescription{
				Index:0,
				Type: typ.Field(0).Type,
				Name: "Id",
			},
		},
		1: {
			des : &types.DependencyDescription{
				Index: 1,
				Type:  typ.Field(1).Type,
				Name:  "alias",
				Flags: types.DEPENDENCY_FLAG_OPTIONAL,
			},
		},
		2: {
			des : &types.DependencyDescription{
				Index: 2,
				Type:  typ.Field(2).Type,
				Name:  "Identity",
				Flags: types.DEPENDENCY_FLAG_EXTENDS,
			},
			sub:depTable{
				0: {
					des :&types.DependencyDescription{
						Index:0,
						Type: typ.Field(2).Type.Field(0).Type,
						Name: "username",
					},
				},
				1: {
					des : &types.DependencyDescription{
						Index: 1,
						Type:  typ.Field(2).Type.Field(1).Type,
						Name:  "PASSWORD",
					},
				},
			},
		},
	}
	test_struct(t,typ,table)

}

func compare_Dependency(t *testing.T, dep types.Dependency, table depTable){
	for scan := dep.AsScan(); scan.Next(); {
		compare_Description(t,scan,table[scan.Index()])
	}
}

func compare_Description(t *testing.T, dst types.DescriptorGetter, table struct{des *types.DependencyDescription;sub depTable}) {
	if nil == table.des{
		t.Errorf("can't matching dependency field %s",dst.Name())
	}

	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(table.des).Elem()

	hasExtends := 0
	if 0 != table.des.Flags & types.DEPENDENCY_FLAG_EXTENDS{
		hasExtends = 1
	}

	for i,n := 0,srcVal.NumField() - hasExtends; i < n; i++ {
		srcField := srcVal.Field(i)
		key := srcVal.Type().Field(i).Name
		dstField := dstVal.MethodByName(key).Call(nil)[0]
		if srcField.Interface() != dstField.Interface() {
			t.Errorf("can't matching dependency field %s,%s: %v != %v",
				dst.Name(),
				key,
				dstField.Interface(),
				srcField.Interface())
		}
	}

	if 0 != hasExtends{
		compare_Dependency(t,dst.Depend(),table.sub)
	}
}
