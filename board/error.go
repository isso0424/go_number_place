package board

import (
	"fmt"
	"isso0424/go_number_place/coordinate"
)

type BoardError interface {
	error
	errorCode() int
}

type AlreadyInputted struct {
	coor coordinate.Coordinate
}

func (e AlreadyInputted) Error() string {
	return fmt.Sprintf("Coordinate (%d %d) is already inputted.", e.coor.GetX(), e.coor.GetY())
}

func (e AlreadyInputted) errorCode() int {
	return 0
}
