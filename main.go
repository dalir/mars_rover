package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var theGrid Grid

func getGridSize(scanner *bufio.Scanner) (grid Grid, err error) {
	//reading first line Grid size
	scanner.Scan()
	stringArray := strings.Split(scanner.Text(), " ")
	grid.x, err = strconv.Atoi(stringArray[0])
	grid.y, err = strconv.Atoi(stringArray[1])
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	theGrid, err = getGridSize(scanner)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		instruction := Instruction{}
		if err = instruction.Read(scanner.Text()); err != nil {
			log.Fatal(err)
		}
		if err = instruction.Run(); err != nil {
			log.Fatal(err)
		}
		instruction.Print()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

}
