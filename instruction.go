package main

import (
	"fmt"
	"strings"
)

// Instruction represents a command sequence for a single robot.
type Instruction struct {
	robot Robot
	moves string
}

// Read parses a line of instruction in the format "(x, y, orientation) commands" and sets up the initial state and movement commands for the robot.
func (i *Instruction) Read(instructionLine string) (err error) {
	trimString := strings.Trim(instructionLine, "(")
	replaceString := strings.Replace(trimString, ")", ",", -1)
	noSpaceString := strings.ReplaceAll(replaceString, " ", "")
	splitString := strings.Split(noSpaceString, ",")
	if err = i.robot.Read(splitString[0], splitString[1], splitString[2]); err != nil {
		return
	}
	i.moves = splitString[3]
	return
}

// Run executes each move command ('F', 'L', 'R') on the robot, updating its state as it processes each command.
func (i *Instruction) Run() (err error) {
	for _, move := range i.moves {
		switch move {
		case int32('F'):
			i.robot.MoveForward()
		case int32('L'):
			i.robot.RotateLeft()
		case int32('R'):
			i.robot.RotateRight()
		default:
			return fmt.Errorf("invalid move instruction: %s", string(move))
		}
	}
	return
}

// Print outputs the current state of the robot, including position, orientation, and if it is marked as "LOST".
func (i *Instruction) Print() {
	i.robot.print()
}
