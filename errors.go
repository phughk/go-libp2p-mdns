package mdns 

import (
	"errors"
)

var (
	ErrInvalidServiceName = errors.New("INVALID_SERVICE_NAME")
	ErrExcessBodySize = errors.New("EXCESS_BODY_SIZE")
)