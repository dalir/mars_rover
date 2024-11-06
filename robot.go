package main

import (
	"fmt"
	"strconv"
)

// Orientation vectors
var North = Pair{0, 1}
var West = Pair{-1, 0}
var South = Pair{0, -1}
var East = Pair{1, 0}

type Status string

const (
	LOST     Status = "LOST"
	IN_RANGE Status = "IN_RANGE"
)

// Pair represents a coordinate or vector on the grid.
type Pair struct {
	x int
	y int
}

// Grid defines the grid's dimensions using Pair struct.
type Grid Pair

// Position represents the robot's position on the grid.
type Position Pair

// Robot holds the current state of the robot, including position, orientation, and status.
type Robot struct {
	curPosition Position
	orientation Pair
	status      Status
}

// updateStatus checks if a position is within grid boundaries, setting the robot's status accordingly.
func (r *Robot) updateStatus(x int, y int) Status {
	if (x < 0 || x > theGrid.x) ||
		(y < 0 || y > theGrid.y) {
		r.status = LOST
	} else {
		r.status = IN_RANGE
	}
	return r.status
}

// RotateLeft rotates the robot 90 degrees counterclockwise.
func (r *Robot) RotateLeft() {
	if r.status == LOST {
		return
	}
	switch r.orientation {
	case North:
		r.orientation = West
	case West:
		r.orientation = South
	case South:
		r.orientation = East
	case East:
		r.orientation = North
	}
}

// RotateRight rotates the robot 90 degrees clockwise.
func (r *Robot) RotateRight() {
	if r.status == LOST {
		return
	}
	switch r.orientation {
	case North:
		r.orientation = East
	case West:
		r.orientation = North
	case South:
		r.orientation = West
	case East:
		r.orientation = South
	}
}

// MoveForward moves the robot one unit forward in its current orientation.
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

// Read initializes the robot's position and orientation based on string inputs.
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
	default:
		return fmt.Errorf("invalid orientation instruction: %s", orientation)
	}
	return
}

// print outputs the robot's position, orientation, and status (LOST if applicable).
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
