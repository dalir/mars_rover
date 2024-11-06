package main

import (
	"bufio"
	"github.com/urfave/cli/v2"
	_ "github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

var theGrid Grid
var ctx *cli.Context

func getGridSize(scanner *bufio.Scanner) (grid Grid, err error) {
	//reading first line Grid size
	scanner.Scan()
	stringArray := strings.Split(scanner.Text(), " ")
	grid.x, err = strconv.Atoi(stringArray[0])
	grid.y, err = strconv.Atoi(stringArray[1])
	return
}

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

func main() {
	app := cli.NewApp()
	app.Name = "mars-rover"
	app.Usage = "running robots on a grid Mars"
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
			Usage:   "path to the input file to read the instruction for grid and the robots",
			Value:   "test1.txt",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
