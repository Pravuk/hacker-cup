//
// .  https://www.facebook.com/codingcompetitions/hacker-cup/2023/practice-round/problems/A2
//

package main

import (
	"bufio"
	"flag"
	"fmt"
	"hacker-cup/common"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	inputFileName  *string
	outputFileName *string
)

func init() {
	inputFileName = flag.String("i", "2023/a2/cheeseburger_corollary_2_input.txt", "")
	//inputFileName = flag.String("i", "2023/a2/cheeseburger_corollary_2_validation_input.txt", "")
	//inputFileName = flag.String("i", "2023/a2/cheeseburger_corollary_2_sample_input.txt", "")
	//outputFileName = flag.String("o", "", "")
	outputFileName = flag.String("o", "2023/a2/output.txt", "")
}

func solution(s, d, c int) int {
	singles := common.HackingFloor(c, s)
	moneyLeft := c - singles*s
	doubles := common.HackingFloor(moneyLeft, d)
	patties := singles + doubles*2

	doubles2 := common.HackingFloor(c, d)
	moneyLeft2 := c - doubles2*d
	singles2 := common.HackingFloor(moneyLeft2, s)
	patties2 := singles2 + doubles2*2

	if patties2+patties == 0 {
		return 0
	}

	if patties > patties2 {
		buns := singles*2 + doubles*2
		return int(math.Min(float64(patties), float64(buns-1)))
	} else {
		buns := singles2*2 + doubles2*2
		return int(math.Min(float64(patties2), float64(buns-1)))
	}
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
		c, _ := strconv.ParseInt(lineSplit[2], 10, 64)
		result := solution(int(s), int(d), int(c))
		printF(fmt.Sprintf("Case #%v: %d", i+1, result))
		//if solution(int(s), int(d), int(c)) {
		//	printF(fmt.Sprintf("Case #%v: YES", i+1))
		//} else {
		//	printF(fmt.Sprintf("Case #%v: NO", i+1))
		//}
		//printF("")
	}
}
