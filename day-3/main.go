package main

import (
	"github.com/kjkondratuk/2021-advent-of-code/lib"
	"log"
	"math"
	"strconv"
)

// Generate:
//   - Gamma Rate - most common bit in position over whole dataset
//   - Epsilon Rate - least common bit in position over the whole dataset
//
//  Answer: multiply gamma rate by epsilon rate & convert to decimal
func main() {
	data := lib.ReadData("inputs/day-3.txt")

	size, counts := summarizeDataByColumn(data)

	log.Printf("Total records: %d", len(data))
	log.Printf("Counts: %+v", counts)

	gamma := make([]int, size)
	epsilon := make([]int, size)
	for i, n := range counts {
		threshold := len(data) / 2
		if n > threshold {
			gamma[i] = 1
			epsilon[i] = 0
		} else if n == threshold {
			// freak out because it didn't say what to do in the instructions
			panic("equal value distribution!")
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}

	gammaRate := binaryArrayToDecimal(gamma)
	log.Printf("Gamma rate: %+v - %d", gamma, gammaRate)
	epsilonRate := binaryArrayToDecimal(epsilon)
	log.Printf("Epsilon rate: %+v - %d", epsilon, epsilonRate)
	log.Printf("Consumption rate: %d", gammaRate*epsilonRate)

	or := findOxyRating(gamma, data)

	log.Printf("Oxy Rating: %s", or)
}

func binaryArrayToDecimal(arr []int) int {
	var acc float64 = 0
	for p, n := range arr {
		//log.Printf("multiplying: %f %f", float64(n), math.Pow(2, float64(len(arr)-p)))
		//log.Printf("2^x : %d", len(arr)-1-p)
		acc += float64(n) * math.Pow(2, float64(len(arr)-1-p))
	}
	return int(acc)
}

func summarizeDataByColumn(data []string, ) (int, []int) {
	bufferSize := 0
	counts := make([]int, 0)
	for r, reading := range data {
		buffer := []rune(reading)
		for i, c := range buffer {
			if r == 0 {
				counts = append(counts, 0)
				bufferSize++
			}

			if c == '1' {
				counts[i]++
			}
		}
	}
	return bufferSize, counts
}

func matchCharAt(s string, c int, i int) bool {
	v, _ :=  strconv.Atoi(string(rune(s[i])))
	return v == c
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func findOxyRating(gamma []int, data []string) string {
	for i := 0; i < len(gamma); i++ {
		log.Printf("Checking gamma index: %d", i)
		for _, _ = range data {
			log.Printf("ranging data: %s", data[i])
			if matchCharAt(data[i], gamma[i], i) {
				//log.Printf("Matched: %s - %s", string(data[i][i]), fmt.Sprintf("%d", gamma[i]))
				if len(data) > 1 {
					log.Printf("Removing: %d - %s", i, data[i])
					data = remove(data, i)
				}
			}/* else {
				log.Printf("Not matched: %s - %s", string(data[i][i]), fmt.Sprintf("%d", gamma[i]))
			}*/
		}
	}
	return ""
}