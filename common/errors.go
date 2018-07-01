package common

import (
	"errors"
)

var (
	ErrModelLimitInvalid    = errors.New("limit Invalid")
	ErrStorageNotSupport    = errors.New("Storage Type Not Support")
	ErrMysqlHandlerNotInit  = errors.New("Mysql Handler Not Init")
	ErrConfigHandlerNotInit = errors.New("Config Handler Not Init")
	ErrParserNotSupport     = errors.New("The Parser Not Support")
)
