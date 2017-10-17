// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package mapper

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"sync"
)

type NamedMapping struct {
	lock  sync.Locker
	table map[string]types.BeanFactory
}

type SelectorMapping struct {
	typ      reflect.Type
	selector types.SelectorGetter
}
