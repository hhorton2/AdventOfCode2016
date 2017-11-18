package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"sync"
)

func main() {
	dat, err := ioutil.ReadFile("Day03/input.txt")
	check(err)
	regex := regexp.MustCompile(".*\n")
	matches := regex.FindAllString(string(dat[:]), -1)
	dayOneTriangles := getDayOneTriangles(matches)
	dayTwoTriangles := getDayTwoTriangles(dayOneTriangles)
	validCount := solve(dayOneTriangles)
	fmt.Println(validCount)
	validCount = solve(dayTwoTriangles)
	fmt.Println(validCount)
}
func getDayOneTriangles(matches []string) []Triangle {
	var dayOneTriangles []Triangle
	for i, match := range matches {
		matches[i] = match[:len(match)-1]
		regex := regexp.MustCompile("\\d+")
		sides := regex.FindAllString(match, -1)
		side1, err1 := strconv.Atoi(sides[0])
		side2, err2 := strconv.Atoi(sides[1])
		side3, err3 := strconv.Atoi(sides[2])
		if err1 == nil && err2 == nil && err3 == nil {
			dayOneTriangles = append(dayOneTriangles, Triangle{side1: side1, side2: side2, side3: side3})
		}
	}
	return dayOneTriangles
}
func getDayTwoTriangles(dayOneTriangles []Triangle) []Triangle {
	var dayTwoTriangles []Triangle
	for i := 0; i < len(dayOneTriangles); i += 3 {
		dayTwoTriangles = append(dayTwoTriangles, Triangle{side1: dayOneTriangles[i].side1, side2: dayOneTriangles[i+1].side1, side3: dayOneTriangles[i+2].side1})
		dayTwoTriangles = append(dayTwoTriangles, Triangle{side1: dayOneTriangles[i].side2, side2: dayOneTriangles[i+1].side2, side3: dayOneTriangles[i+2].side2})
		dayTwoTriangles = append(dayTwoTriangles, Triangle{side1: dayOneTriangles[i].side3, side2: dayOneTriangles[i+1].side3, side3: dayOneTriangles[i+2].side3})
	}
	return dayTwoTriangles
}
func solve(triangles []Triangle) int {
	wg := new(sync.WaitGroup)
	validTriangleChannel := make(chan bool, len(triangles))
	for _, triangle := range triangles {
		wg.Add(1)
		go checkIfTriangle(triangle, wg, validTriangleChannel)
	}
	wg.Wait()
	close(validTriangleChannel)
	validCount := 0
	for result := range validTriangleChannel {
		if result {
			validCount++
		}
	}
	return validCount
}
func checkIfTriangle(triangle Triangle, wg *sync.WaitGroup, isValidChn chan<- bool) {
	defer wg.Done()
	isValid := triangle.side1 < triangle.side2+triangle.side3
	isValid = isValid && triangle.side2 < triangle.side1+triangle.side3
	isValid = isValid && triangle.side3 < triangle.side2+triangle.side1
	isValidChn <- isValid
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Triangle struct {
	side1 int
	side2 int
	side3 int
}
