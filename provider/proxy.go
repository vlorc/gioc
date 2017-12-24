// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package provider

import (
	"reflect"
)

func (pp *ProxyProvider) Resolve(impType interface{}, args ...string) (interface{}, error) {
	return pp.provider.Resolve(impType, args...)
}

func (pp *ProxyProvider) ResolveType(typ reflect.Type, name string, deep int) (interface{}, error) {
	return pp.provider.ResolveType(typ, name, deep)
}
func (pp *ProxyProvider) ResolveNamed(impType interface{}, name string, deep int) (interface{}, error) {
	return pp.provider.ResolveNamed(impType, name, deep)
}

func (pp *ProxyProvider) Assign(dst interface{}, args ...string) (err error) {
	return pp.provider.Assign(dst, args...)
}

func (pp *ProxyProvider) AssignType(dst reflect.Value, typ reflect.Type, name string, deep int) error {
	return pp.provider.AssignType(dst, typ, name, deep)
}

func (pp *ProxyProvider) AssignNamed(dst interface{}, impType interface{}, name string, deep int) error {
	return pp.provider.AssignNamed(dst, impType, name, deep)
}
