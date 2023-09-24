//
// .  https://www.facebook.com/codingcompetitions/hacker-cup/2023/practice-round/problems/C
//

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/samber/lo"
	"hacker-cup/common"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	inputFileName  *string
	outputFileName *string
)

func init() {
	inputFileName = flag.String("i", "2023/c/two_apples_a_day_input.txt", "")
	//outputFileName = flag.String("o", "", "")
	outputFileName = flag.String("o", "2023/c/output.txt", "")
}

func numsEx2Nums(nums []int, target int) int {
	var sums = make(map[int][]int, len(nums))
	for i, num := range nums {
		if _, ok := sums[target-num]; !ok {
			sums[target-num] = []int{i}
		} else {
			sums[target-num] = append(sums[target-num], i)
		}
	}

	var unpairs []int
	for i, num := range nums {
		if _, ok := sums[num]; ok {
			if len(sums[num]) > 0 {
				j := sums[num][0]
				if i != j {
					pair := nums[j]
					if _, ok1 := sums[pair]; ok1 && len(sums[pair]) > 0 {
						sums[num] = common.Remove(sums[num], 0)
						sums[pair] = common.Remove(sums[pair], 0)
					}
				}
			}
		} else {
			unpairs = append(unpairs, num)
		}
	}
	if len(unpairs) == 1 {
		return unpairs[0]
	} else if len(unpairs) > 1 {
		return -1
	} else {
		for key, _ := range sums {
			if len(sums[key]) == 1 {
				return nums[sums[key][0]]
			} else if len(sums[key]) > 1 {
				return -1
			}
		}
	}
	return -1
}

func possibleEqualApples(apples []int) []int {
	sort.Ints(apples)

	if len(apples) < 3 {
		return apples
	}
	var result []int
	result = append(result, apples[0]+apples[len(apples)-1])
	result = append(result, apples[1]+apples[len(apples)-1])
	result = append(result, apples[0]+apples[len(apples)-2])
	if len(apples) > 3 {
		result = append(result, apples[1]+apples[len(apples)-2])
	}
	return result
}
func allEqual(apples []int) bool {
	last := apples[0]
	for i, cur := range apples {
		if i > 0 {
			if last != cur {
				return false
			}
		}
	}
	return true
}
func solution(days int, apples []int) int {
	if len(apples) == 1 {
		return apples[0]
	}
	if allEqual(apples) {
		return apples[0]
	}
	total := lo.Sum(apples)
	var solvedNums []int
	possiblesSums := possibleEqualApples(apples)
	possiblesSums = lo.UniqBy(possiblesSums, func(item int) int {
		return item
	})
	for _, sum := range possiblesSums {
		candidate := numsEx2Nums(apples, sum)
		if candidate >= 1 {
			solvedNums = append(solvedNums, sum-candidate)
		}
	}
	if len(solvedNums) > 0 {
		res := lo.Min(solvedNums)
		if res <= 0 {
			res = lo.Max(solvedNums)
			if res > 0 && (total+res)%days == 0 {
				return res
			}
		} else if (total+res)%days == 0 {
			return res
		}
	}
	return -1
}

func main() {
	start := time.Now()
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

	fileScanner := bufio.NewScanner(file) // MaxScanTokenSize overridden to 600000
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	casesAmount, err := strconv.ParseInt(fileScanner.Text(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	for i := int64(0); i < casesAmount; i++ {
		fileScanner.Scan()
		days, _ := strconv.ParseInt(fileScanner.Text(), 10, 64)
		fileScanner.Scan()
		line := fileScanner.Text()
		lineSplit := strings.Split(line, " ")
		apples := lo.Map(lineSplit, func(item string, index int) int {
			r, _ := strconv.ParseInt(item, 10, 64)
			return int(r)
		})
		result := solution(int(days), apples)
		printF(fmt.Sprintf("Case #%v: %d", i+1, result))
	}
	elapsed := time.Since(start)
	fmt.Printf("Done. Time execution %d microseconds", elapsed/1000)
}
