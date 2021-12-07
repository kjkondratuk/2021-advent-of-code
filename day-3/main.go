package main

import (
	"errors"
	"fmt"
	"github.com/kjkondratuk/2021-advent-of-code/lib"
	"log"
	"math"
)

// Generate:
//   - Gamma Rate - most common bit in position over whole dataset
//   - Epsilon Rate - least common bit in position over the whole dataset
//
//  Answer: multiply gamma rate by epsilon rate & convert to decimal

var (
	ErrNoWinner = errors.New("there was a tie")
)

func main() {
	data := lib.NewDataReader("inputs/day-3.txt").Read().([]string)

	log.Printf("Total records: %d", len(data))

	arr := create2DBooleanArray(data)

	lineLen := len(arr[0])

	// PART 1
	gamma := make([]bool, 0)
	for i := 0; i < lineLen; i++ {
		summary, err := summarizeDataByColumn(arr, i)
		if err != nil {
			panic(err)
		}
		gamma = append(gamma, summary)
	}
	epsilon := inverse(gamma)

	gammaRate := binaryArrayToDecimal(gamma)
	log.Printf("Gamma rate: %+v - %d", gamma, gammaRate)
	epsilonRate := binaryArrayToDecimal(epsilon)
	log.Printf("Epsilon rate: %+v - %d", epsilon, epsilonRate)
	log.Printf("Power Consumption rate: %d", gammaRate*epsilonRate)

	// Added to prevent regresssion during refactoring
	//if gammaRate*epsilonRate != 1540244 {
	//	panic("incorrect power consumption")
	//}

	// PART 2
	o2rating := FilterableBinaryList(arr).ProgressivePopularityFilterWithDefault(true)
	co2rating := FilterableBinaryList(arr).ProgressivePopularityFilterWithDefault(false)

	o2RatingDecimal := binaryArrayToDecimal(o2rating)
	log.Printf("Oxy Rating: %d", o2RatingDecimal)

	co2RatingDecimal := binaryArrayToDecimal(co2rating)
	log.Printf("Co2 Rating: %d", co2RatingDecimal)

	log.Printf("Life Support Rating: %d", o2RatingDecimal*co2RatingDecimal)
}

type FilterableBinaryList [][]bool

type BinaryListPredicate func(data []bool) bool

func (l FilterableBinaryList) Filter(p BinaryListPredicate) [][]bool {
	filtered := make([][]bool, 0)
	for _, item := range l {
		if p(item) {
			fmt.Printf("preserving: %+v\n", item)
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// ProgressivePopularityFilterWithDefault : filters progressively based on the popularity of a binary value
// starting at the left-most position and moving right until there is a single element left.
func (l FilterableBinaryList) ProgressivePopularityFilterWithDefault(dflt bool) []bool {
	var v []bool
	arrCopy := l
	i := 0
	for {
		pop, err := summarizeDataByColumn(arrCopy, i)
		if err == ErrNoWinner {
			fmt.Printf("using default of: %t\n", dflt)
			if dflt {
				pop = dflt
			} else {
				// invert the chosen default because we are inverting the comparison in the filter below
				pop = !dflt
			}
		}

		log.Printf("pop is: %t", pop)
		log.Printf("Filtering: %t at index %d", !pop, i)
		log.Printf("Len (B): %d", len(arrCopy))
		arrCopy = arrCopy.Filter(func(data []bool) bool {
			if dflt {
				//log.Printf("Filter most popular: %t - %t", data[i], pop)
				return data[i] == pop
			} else {
				//log.Printf("Filter least popular: %t - %t", data[i], pop)
				return data[i] == !pop
			}
		})
		log.Printf("Len (A): %d", len(arrCopy))

		if len(arrCopy) == 1 {
			v = arrCopy[0]
			break
		}
		i++
	}
	return v
}

func binaryArrayToDecimal(arr []bool) int {
	var acc float64 = 0
	for p, n := range arr {
		var v float64
		if n {
			v = 1
		}
		acc += v * math.Pow(2, float64(len(arr)-1-p))
	}
	return int(acc)
}

func inverse(i []bool) []bool {
	r := make([]bool, 0)
	for _, item := range i {
		r = append(r, !item)
	}
	return r
}

func summarizeDataByColumn(data [][]bool, col int) (bool, error) {
	var total float64 = 0
	for _, row := range data {
		if row[col] {
			total++
		}
	}

	majority := float64(len(data)) / 2.0
	fmt.Printf("total: %f - majority: %f - len: %d\n", total, majority, len(data))
	if total > majority {
		fmt.Println("total greater than majority")
		return true, nil
	} else if total == majority {
		fmt.Println("total equal to majority")
		return false, ErrNoWinner
	} else {
		fmt.Println("total less than majority")
		return false, nil
	}
}

func create2DBooleanArray(data []string) [][]bool {
	arr := make([][]bool, 0)
	for a, d := range data {
		arr = append(arr, make([]bool, 0))
		for _, c := range d {
			v := false
			if c == '1' {
				v = true
			}
			arr[a] = append(arr[a], v)
		}
	}
	return arr
}
