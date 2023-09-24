package common

import (
	"log"
	"os"
)

type printFunc func(line string)

func GetPrint(print2file string) (printFunc, *os.File, error) {
	if len(print2file) == 0 {
		return func(line string) {
			println(line)
		}, nil, nil
	} else {
		outFile, err := os.Create(print2file)
		return func(line string) {
			_, err := outFile.WriteString(line)
			if err != nil {
				log.Fatal(err)
			}
			outFile.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}, outFile, err
	}

}

type Direction struct {
	RowOffset, ColOffset int
}

func NewNonDiagonalDirections() []Direction {
	return []Direction{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
}

func HackingFloor(a, b int) int {
	r := a % b
	return (a - r) / b
}

func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
