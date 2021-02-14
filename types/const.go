// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package types

type Token int

type Kind uint

type DependencyFlag int

type ErrorCode int

const (
	_                             = iota
	ErrInstanceNotFound ErrorCode = iota
	ErrFactoryNotFound
	ErrTypeNotFound
	ErrDependencyNotFound
	ErrDependencyNotNeed
	ErrTypeNotSet
	ErrTypeNotMatch
	ErrTypeNotSupport
	ErrInjectNotAllot
	ErrTypeImplements
	ErrTypeNotPointer
	ErrTypeNotInterface
	ErrTypeNotFunction
	ErrTypeNotConvert
	ErrIndexNotSupport
	ErrTokenNotSupport
	ErrExtendNotSupport
)

var errFormatTable = map[ErrorCode]string{
	ErrInstanceNotFound:  `can't find instance of type '%s' - '%s'`,
	ErrFactoryNotFound:   `can't find factory of type '%s' - '%s'`,
	ErrTypeNotFound:      `can't find type '%s' - '%s'`,
	ErrTypeNotSet:        `can't set type '%s'`,
	ErrTypeNotMatch:      `can't matching type '%s'`,
	ErrTypeNotSupport:    `can't support type '%s'`,
	ErrDependencyNotNeed: `Don't need dependency '%s'`,
	ErrInjectNotAllot:    `can't need allot inject '%s'`,
	ErrTypeImplements:    `can't implements type '%s'`,
	ErrTypeNotPointer:    `don't is pointer type '%s'`,
	ErrTypeNotInterface:  `don't is interface type '%s'`,
	ErrTypeNotFunction:   `don't is interface type '%s'`,
	ErrTypeNotConvert:    `don't convert type '%s'`,
	ErrIndexNotSupport:   `can't support type '%s'`,
	ErrTokenNotSupport:   `can't support token '%s'`,
	ErrExtendNotSupport:  `can't support extend type '%s'`,
}

const (
	DEPENDENCY_FLAG_EXTENDS DependencyFlag = 1 << iota
	DEPENDENCY_FLAG_LAZY
	DEPENDENCY_FLAG_DEFAULT
	DEPENDENCY_FLAG_OPTIONAL
	DEPENDENCY_FLAG_UNSAFE
	DEPENDENCY_FLAG_SKIP
	DEPENDENCY_FLAG_REQUEST
	DEPENDENCY_FLAG_ANY
	DEPENDENCY_FLAG_MAKE
)

const (
	TOKEN_EOF     Token = -1
	TOKEN_ERROR         = 2
	TOKEN_CHART         = 4
	TOKEN_ECHART        = 7
	TOKEN_STRING        = 9
	TOKEN_ESTRING       = 12
	TOKEN_IDENT         = 13
	TOKEN_LPAREN        = 14
	TOKEN_RPAREN        = 15
	TOKEN_COMMA         = 16
	TOKEN_NUMBER        = 19
	TOKEN_FLOAT         = 20
	TOKEN_NULL          = 23
)

const (
	Invalid Kind = iota
	Int
	Float
	String
	Boolean
	Null
)
