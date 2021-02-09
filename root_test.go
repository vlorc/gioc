// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package gioc

import (
	"fmt"
	"github.com/vlorc/gioc/types"
	"testing"
)

func Test_Invoker(t *testing.T) {
	root := NewRootContainer()

	getKey := func(id int64, name *string) (r string) {
		if nil != name {
			r = fmt.Sprintf("id(%d) - name(%s)", id, *name)
		} else {
			r = fmt.Sprintf("id(%d)", id)
		}
		return
	}

	name := "angel"
	root.AsRegister().RegisterInstance(&name)
	var dependFactory types.DependencyFactory
	var invokerFactory types.InvokerFactory
	root.AsProvider().Assign(&dependFactory)
	root.AsProvider().Assign(&invokerFactory)

	dep, err := dependFactory.Instance(getKey)
	if nil != err {
		t.Errorf("can't allocate a depend error : %s", err.Error())
	}
	invoker, err := invokerFactory.Instance(getKey, dep)
	if nil != err {
		t.Errorf("can't allocate a invoker error : %s", err.Error())
	}

	results := invoker.ApplyWith(root.AsProvider(), 10)
	t.Log("getKey", results[0].Interface())
}
