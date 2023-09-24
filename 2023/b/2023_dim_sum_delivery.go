//
// .  https://www.facebook.com/codingcompetitions/hacker-cup/2023/practice-round/problems/B
//

package main

import (
	"bufio"
	"flag"
	"fmt"
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
	inputFileName = flag.String("i", "2023/b/dim_sum_delivery_input.txt", "")
	//outputFileName = flag.String("o", "", "")
	outputFileName = flag.String("o", "2023/b/output.txt", "")
}

func solution(r, c int) bool {
	return r > c
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

		if solution(int(r), int(c)) {
			printF(fmt.Sprintf("Case #%v: YES", i+1))
		} else {
			printF(fmt.Sprintf("Case #%v: NO", i+1))
		}
		//printF("")
	}
}
