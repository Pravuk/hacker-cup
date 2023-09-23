package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gammazero/deque"
	"hacker-cup/common"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	inputFileName  *string
	outputFileName *string
)

func init() {
	inputFileName = flag.String("i", "second_friend_input.txt", "")
	outputFileName = flag.String("o", "output.txt", "")
}

func solution(r, c int, forest []string) (bool, []string) {
	var queueToBeHappy deque.Deque[[]int]
	var pretendToPutTree = make([][]int, r)
	for i := 0; i < r; i++ {
		pretendToPutTree[i] = make([]int, c)
	}
	directs := common.NewNonDiagonalDirections()
	for i, _ := range forest {
		for j, _ := range forest[i] {
			if forest[i][j] == '#' {
				continue
			}

			for _, dir := range directs {
				newRow := i + dir.RowOffset
				newCol := j + dir.ColOffset
				if newRow < 0 || newRow >= len(forest) ||
					newCol < 0 || newCol >= len(forest[0]) ||
					forest[newRow][newCol] == '#' {
					continue
				}
				pretendToPutTree[i][j]++
			}
			if pretendToPutTree[i][j] < 2 && forest[i][j] == '^' {
				return false, nil
			}
			if pretendToPutTree[i][j] == 1 {
				queueToBeHappy.PushBack([]int{i, j})
			}
		}
	}

	for queueToBeHappy.Len() > 0 {
		size := queueToBeHappy.Len()
		for k := 0; k < size; k++ {
			tree := queueToBeHappy.PopFront()
			r = tree[0]
			c = tree[1]
			if forest[r][c] == '^' {
				return false, nil
			}
			pretendToPutTree[r][c] = 0
			for _, dir := range directs {
				newRow := r + dir.RowOffset
				newCol := c + dir.ColOffset
				if newRow < 0 || newRow >= len(forest) ||
					newCol < 0 || newCol >= len(forest[0]) ||
					pretendToPutTree[newRow][newCol] == 0 {
					continue
				}
				if forest[newRow][newCol] == '^' && pretendToPutTree[newRow][newCol] == 1 {
					return false, nil
				}
				pretendToPutTree[newRow][newCol]--
				if pretendToPutTree[newRow][newCol] == 1 {
					queueToBeHappy.PushBack([]int{newRow, newCol})
				}
			}
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if pretendToPutTree[i][j] == 0 {
				continue
			}
			row := []rune(forest[i])
			row[j] = '^'
			forest[i] = string(row)
		}
	}

	return true, forest
}

func main() {
	flag.Parse()

	file, err := os.Open(*inputFileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	printF, ofile, err := common.GetPrint(*outputFileName)
	if err != nil {
		log.Fatal(err)
	}

	if ofile != nil {
		defer ofile.Close()
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	casesAmount, err := strconv.ParseInt(fileScanner.Text(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	for i := int64(0); i < casesAmount; i++ {
		fileScanner.Scan()
		lineSplit := strings.Split(fileScanner.Text(), " ")
		r, _ := strconv.ParseInt(lineSplit[0], 10, 64)
		c, _ := strconv.ParseInt(lineSplit[1], 10, 64)
		var forest []string
		for j := int64(0); j < r; j++ {
			fileScanner.Scan()
			forest = append(forest, fileScanner.Text())
			r, forest := solution(int(r), int(c), forest)
			if r {
				printF(fmt.Sprintf("Case #%v: Possible", j+1))
				for _, row := range forest {
					printF(row)
				}
			} else {
				printF(fmt.Sprintf("Case #%v: Impossible", j+1))
			}
		}
	}
}
