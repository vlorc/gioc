// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package register

import (
	"github.com/vlorc/gioc/selector"
	"github.com/vlorc/gioc/types"
)

func (r *readOnlyRegister) Interface(instance interface{}, args ...string) {

}

func (r *readOnlyRegister) Set(instance interface{}, args ...string) {

}

func (r *readOnlyRegister) Add(instance interface{}, args ...string) {

}

func (r *readOnlyRegister) Put(instance interface{}, args ...string) bool {
	return true
}

func (r *readOnlyRegister) Factory(factory types.BeanFactory, impType interface{}, args ...string) {

}

func (r *readOnlyRegister) Selector() types.Selector {
	return selector.NewReadOnlyFactory(r.getter)
}
