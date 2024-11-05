package main

import (
	"fmt"
	"strconv"
)

var North = Pair{0, 1}
var South = Pair{0, -1}
var East = Pair{1, 0}
var West = Pair{-1, 0}

type Status string

const LOST Status = "LOST"
const IN_RANGE Status = "IN_RANGE"

type Pair struct {
	x int
	y int
}

type Grid Pair

type Position Pair

type Robot struct {
	curPosition Position
	orientation Pair
	status      Status
}

func (r *Robot) updateStatus(x int, y int) Status {
	if (x < 0 || x > theGrid.x) ||
		(y < 0 || y > theGrid.y) {
		r.status = LOST
	} else {
		r.status = IN_RANGE
	}
	return r.status
}

func (r *Robot) RotateLeft() {
	if r.status == LOST {
		return
	}
	r.orientation.x = (r.orientation.x - 1) % 2
	r.orientation.y = (r.orientation.y + 1) % 2
}

func (r *Robot) RotateRight() {
	if r.status == LOST {
		return
	}
	r.orientation.x = (r.orientation.x + 1) % 2
	r.orientation.y = (r.orientation.y - 1) % 2
}

func (r *Robot) MoveForward() {
	if r.status == LOST {
		return
	}
	xCandidate := r.curPosition.x + r.orientation.x
	yCandidate := r.curPosition.y + r.orientation.y
	if r.updateStatus(xCandidate, yCandidate) == IN_RANGE {
		r.curPosition.x = xCandidate
		r.curPosition.y = yCandidate
	}
}

func (r *Robot) Read(x string, y string, orientation string) (err error) {
	r.curPosition.x, err = strconv.Atoi(x)
	if err != nil {
		return
	}
	r.curPosition.y, err = strconv.Atoi(y)
	if err != nil {
		return
	}
	switch orientation {
	case "N":
		r.orientation = North
	case "S":

		r.orientation = South
	case "W":
		r.orientation = West
	case "E":
		r.orientation = East
	}
	return
}

func (r *Robot) print() {
	var orientation string
	switch r.orientation {
	case North:
		orientation = "N"
	case South:
		orientation = "S"
	case West:
		orientation = "W"
	case East:
		orientation = "E"
	}
	if r.status == LOST {
		fmt.Printf("(%d, %d, %s) LOST\n", r.curPosition.x, r.curPosition.y, orientation)
	} else {
		fmt.Printf("(%d, %d, %s)\n", r.curPosition.x, r.curPosition.y, orientation)
	}
}
