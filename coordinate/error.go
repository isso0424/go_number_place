package coordinate

import "fmt"

type ErrorCode int

type CoordinateError interface {
	error
	GetErrorCode() ErrorCode
}

type ValueNotFound struct {
	value int
}

func (e ValueNotFound) Error() string {
	return fmt.Sprintf("%d is not found", e.value)
}

func (e ValueNotFound) GetErrorCode() ErrorCode {
	return 0
}
