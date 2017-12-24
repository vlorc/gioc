// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package text

import (
	"github.com/vlorc/gioc/types"
)

func NewTagParser() types.TextParser {
	return NewTextParser(map[string][]types.IdentHandle{
		"optional": {
			flagsHandle(types.DEPENDENCY_FLAG_OPTIONAL),
		},
		"extends": {
			flagsHandle(types.DEPENDENCY_FLAG_EXTENDS),
			extendsHandle,
		},
		"default": {
			flagsHandle(types.DEPENDENCY_FLAG_DEFAULT),
			defaultHandle,
		},
		"lazy": {
			flagsHandle(types.DEPENDENCY_FLAG_LAZY),
			lazyHandle,
		},
		"id":    {nameHandle},
		"name":  {nameHandle},
		"lower": {lowerCaseHandle},
		"upper": {upperCaseHandle},
	}, NewParamFactory())
}
