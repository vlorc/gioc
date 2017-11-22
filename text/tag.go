// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package text

import (
	"github.com/vlorc/gioc/types"
)

func NewTagParser() types.TextParser {
	return NewTextParser(map[string][]types.IdentHandle{
		"optional": []types.IdentHandle{
			flagsHandle(types.DEPENDENCY_FLAG_OPTIONAL),
		},
		"extends": []types.IdentHandle{
			flagsHandle(types.DEPENDENCY_FLAG_EXTENDS),
			extendsHandle,
		},
		"default": []types.IdentHandle{
			flagsHandle(types.DEPENDENCY_FLAG_DEFAULT),
			defaultHandle,
		},
		"lazy": []types.IdentHandle{
			flagsHandle(types.DEPENDENCY_FLAG_LAZY),
			lazyHandle,
		},
		"id":    []types.IdentHandle{nameHandle},
		"name":  []types.IdentHandle{nameHandle},
		"lower": []types.IdentHandle{lowerCaseHandle},
		"upper": []types.IdentHandle{upperCaseHandle},
	},NewParamFactory())
}