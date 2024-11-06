package main

import (
	"bufio"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

var theGrid Grid
var ctx *cli.Context

// getGridSize reads the grid dimensions from the first line of the input file.
func getGridSize(scanner *bufio.Scanner) (grid Grid, err error) {
	scanner.Scan()
	stringArray := strings.Split(scanner.Text(), " ")
	grid.x, err = strconv.Atoi(stringArray[0])
	grid.y, err = strconv.Atoi(stringArray[1])
	return
}

// MarsRover runs the main simulation. It reads input commands for the grid size and each robot,
// then executes the instructions, updating and printing the robot states.
func MarsRover() (err error) {
	var file *os.File
	file, err = os.Open(ctx.String("input-file"))
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	theGrid, err = getGridSize(scanner)
	if err != nil {
		return
	}

	for scanner.Scan() {
		instruction := Instruction{}
		if err = instruction.Read(scanner.Text()); err != nil {
			return
		}
		if err = instruction.Run(); err != nil {
			return
		}
		instruction.Print()
		if err = scanner.Err(); err != nil {
			return
		}
	}
	return
}

// main initialises the CLI application and parses flags. It runs MarsRover based on input arguments.
func main() {
	app := cli.NewApp()
	app.Name = "mars-rover"
	app.Usage = "simulates robots navigating a grid on Mars"
	app.Action = func(c *cli.Context) (err error) {
		ctx = c
		err = MarsRover()
		if err != nil {
			return err
		}
		return err
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "input-file",
			Aliases: []string{"i"},
			Usage:   "path to the input file with grid and robot instructions",
			Value:   "test1.txt",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
