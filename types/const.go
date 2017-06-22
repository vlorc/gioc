// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package types

type DependencyFlag int

type ErrorCode int

const (
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
}

const (
	DEPENDENCY_FLAG_EXTENDS DependencyFlag = 1 << iota
	DEPENDENCY_FLAG_OPTIONAL
	DEPENDENCY_FLAG_DEFAULT
)

const DEFAULT_NAME string = ""
