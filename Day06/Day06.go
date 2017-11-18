package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	dat, err := ioutil.ReadFile("Day06/input.txt")
	check(err)
	datString := string(dat)
	rowStream := make(chan string, 10)
	positionCountStream := make(chan PositionCount, 10)
	positionCountStreamForkA := make(chan PositionCount, 10)
	positionCountStreamForkB := make(chan PositionCount, 10)
	dayOneResultStream := make(chan string)
	dayTwoResultStream := make(chan string)
	go produceRows(datString, rowStream)
	go getPositionCounts(rowStream, positionCountStream)
	go duplicatePositionCounts(positionCountStream, positionCountStreamForkA, positionCountStreamForkB)
	go aggregateMaximumPositionCounts(positionCountStreamForkA, dayOneResultStream)
	go aggregateMinimumPositionCounts(positionCountStreamForkB, dayTwoResultStream)
	fmt.Printf("Day One Result: %v\nExecution took: %v\n", <-dayOneResultStream, time.Since(start))
	fmt.Printf("Day Two Result: %v\nExecution took: %v\n", <-dayTwoResultStream, time.Since(start))

}
func aggregateMinimumPositionCounts(countStream chan PositionCount, resultStream chan string) {
	defer close(resultStream)
	countMap := make(map[int]map[string]int)
	for count := range countStream {
		if _, ok := countMap[count.position]; ok {
			countMap[count.position][count.value]++
		} else {
			countMap[count.position] = make(map[string]int)
			countMap[count.position][count.value]++
		}
	}
	result := [8]byte{}
	for position, letterMap := range countMap {
		positionMin := math.MaxInt32
		positionLetter := ""
		for letter, count := range letterMap {
			if count < positionMin {
				positionMin = count
				positionLetter = letter
			}
		}
		result[position] = []byte(positionLetter)[0]
	}
	resultStream <- string(result[:])

}
func duplicatePositionCounts(positionCounts <-chan PositionCount, forkA chan<- PositionCount, forkB chan<- PositionCount) {
	defer close(forkA)
	defer close(forkB)
	for position := range positionCounts {
		forkA <- position
		forkB <- position
	}
}
func aggregateMaximumPositionCounts(countStream <-chan PositionCount, resultStream chan<- string) {
	defer close(resultStream)
	countMap := make(map[int]map[string]int)
	for count := range countStream {
		if _, ok := countMap[count.position]; ok {
			countMap[count.position][count.value]++
		} else {
			countMap[count.position] = make(map[string]int)
			countMap[count.position][count.value]++
		}
	}
	result := [8]byte{}
	for position, letterMap := range countMap {
		positionMax := 0
		positionLetter := ""
		for letter, count := range letterMap {
			if count > positionMax {
				positionMax = count
				positionLetter = letter
			}
		}
		result[position] = []byte(positionLetter)[0]
	}
	resultStream <- string(result[:])
}
func getPositionCounts(rowStream <-chan string, positionCounts chan<- PositionCount) {
	defer close(positionCounts)
	for row := range rowStream {
		for i := 0; i < len(row); i++ {
			positionCounts <- PositionCount{position: i, value: row[i : i+1]}
		}
	}
}
func produceRows(input string, rowStream chan<- string) {
	defer close(rowStream)
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		rowStream <- row
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type PositionCount struct {
	position int
	value    string
}
