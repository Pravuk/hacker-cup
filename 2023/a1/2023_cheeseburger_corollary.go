//
// .  https://www.facebook.com/codingcompetitions/hacker-cup/2023/practice-round/problems/A1
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
	//inputFileName = flag.String("i", "2023/a1/cheeseburger_corollary_1_sample_input.txt", "")
	inputFileName = flag.String("i", "2023/a1/cheeseburger_corollary_1_input.txt", "")
	//inputFileName = flag.String("i", "2023/a1/cheeseburger_corollary_1_validation_input.txt", "")
	//outputFileName = flag.String("o", "", "")
	outputFileName = flag.String("o", "2023/a1/output.txt", "")
}
func solution(s, d, k int) bool {
	//buns := s * 2
	//buns += d * 2
	//patties := s
	//patties += d * 2
	//return k <= buns-1 && k <= patties
	return k <= s*2+d*2-1 && k <= s+d*2
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
		s, _ := strconv.ParseInt(lineSplit[0], 10, 64)
		d, _ := strconv.ParseInt(lineSplit[1], 10, 64)
		k, _ := strconv.ParseInt(lineSplit[2], 10, 64)
		if solution(int(s), int(d), int(k)) {
			printF(fmt.Sprintf("Case #%v: YES", i+1))
		} else {
			printF(fmt.Sprintf("Case #%v: NO", i+1))
		}
		//printF("")
	}
}
