// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package types

import (
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

func NewStringFactory(s string) StringFactory {
	return RawStringFactory(s)
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
