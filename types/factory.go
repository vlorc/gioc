// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package types

import (
	"fmt"
	"github.com/vlorc/gioc/utils"
	"reflect"
	"sort"
)

type WrapperFactory struct {
	value   []int
	factory []func(BeanFactory) BeanFactory
}

type RawStringFactory string

func NewWrapperFactory() *WrapperFactory {
	return &WrapperFactory{}
}

type NameFactory string

func NewStringFactory(s string) StringFactory {
	return RawStringFactory(s)
}

func NewNameFactory(s string) StringFactory {
	return NameFactory(s)
}

func (l *WrapperFactory) Reset() {
	l.value = nil
	l.factory = nil
}

func (l *WrapperFactory) Append(value int, factory func(BeanFactory) BeanFactory) {
	l.value = append(l.value, value)
	l.factory = append(l.factory, factory)
}

func (l *WrapperFactory) Instance(b BeanFactory) (BeanFactory, error) {
	sort.Sort(l)

	for i := range l.factory {
		b = l.factory[i](b)
	}
	return b, nil
}

func (l *WrapperFactory) Len() int {
	return len(l.value)
}

func (l *WrapperFactory) Less(i, j int) bool {
	return l.value[i] > l.value[j]
}

func (l *WrapperFactory) Swap(i, j int) {
	l.value[i], l.value[j] = l.value[j], l.value[i]
	l.factory[i], l.factory[j] = l.factory[j], l.factory[i]
}

func (f RawStringFactory) Instance(provider Provider) (string, error) {
	return string(f), nil
}

func (f NameFactory) Instance(provider Provider) (string, error) {
	b := provider.Factory(StringType, string(f), -1)
	if nil == b {
		if b = provider.Factory(nil, string(f), -1); nil == b {
			utils.Panic(NewWithError(ErrFactoryNotFound, StringType, string(f)))
		}
	}

	val, err := b.Instance(provider)
	if nil != err {
		return "", err
	}

	if str, ok := val.(string); ok {
		return str, nil
	}
	if str, ok := val.(fmt.Stringer); ok {
		return str.String(), nil
	}
	if v := utils.ValueOf(val); reflect.Ptr == v.Kind() {
		return fmt.Sprint(utils.IndirectValue(v).Interface()), nil
	}
	return fmt.Sprint(val), nil
}
