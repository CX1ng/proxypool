package common

import (
	"errors"
)

var (
	ModelLimitInvalid = errors.New("limit Invalid")
	StorageNotSupport = errors.New("Storage Type Not Support")
)
