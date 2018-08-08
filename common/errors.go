package common

import (
	"errors"
)

var (
	ErrModelLimitInvalid       = errors.New("limit Invalid")
	ErrStorageNotSupport       = errors.New("Storage Type Not Support")
	ErrMysqlHandlerNotInit     = errors.New("Mysql Handler Not Init")
	ErrConfigHandlerNotInit    = errors.New("Config Handler Not Init")
	ErrParserNotSupport        = errors.New("The Parser Not Support")
	ErrRedisHandlerNotInit     = errors.New("Redis Handler Not Init")
	ErrBeginPageNumLessThanOne = errors.New("BeginPageNum Shouldn't Less Than One")
	ErrEndPageNumLessThanZero  = errors.New("EndPageNum Shouldn't Less Than Zero")
)
