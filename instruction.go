package main

import (
	"fmt"
	"strings"
)

type Instruction struct {
	robot Robot
	moves string
}

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

func (i *Instruction) Print() {
	i.robot.print()
}
