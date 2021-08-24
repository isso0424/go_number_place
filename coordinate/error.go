package coordinate

import "fmt"

type CoordinateError interface {
	error
	GetErrorCode() int
}
