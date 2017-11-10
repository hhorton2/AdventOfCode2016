package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
)

func main() {
	dat, err := ioutil.ReadFile("Day1/input.txt")
	check(err)
	dataArray := strings.Split(strings.Replace(string(dat), "\n", "", -1), ", ")
	solvePartOne(dataArray)
	solvePartTwo(dataArray)
}
func solvePartOne(dataArray []string) {
	North := new(Direction)
	East := new(Direction)
	South := new(Direction)
	West := new(Direction)
	North.left = West
	North.right = East
	East.left = North
	East.right = South
	South.left = East
	South.right = West
	West.left = South
	West.right = North
	currentDirection := North
	for _, instruction := range dataArray {
		direction := getDirection(instruction)
		steps := getSteps(instruction)
		if direction == "R" {
			currentDirection = currentDirection.right
		} else {
			currentDirection = currentDirection.left
		}
		currentDirection.steps += steps
	}
	fmt.Printf("north: %v, east: %v, south: %v, west: %v\n", North.steps, East.steps, South.steps, West.steps)
	fmt.Printf("blocks away: %v\n", math.Abs(float64(North.steps-South.steps))+math.Abs(float64(East.steps-West.steps)))
}
func solvePartTwo(dataArray []string) {
	North := new(Direction)
	East := new(Direction)
	South := new(Direction)
	West := new(Direction)
	North.left = West
	North.right = East
	East.left = North
	East.right = South
	South.left = East
	South.right = West
	West.left = South
	West.right = North
	currentDirection := North
	var grid []Vertex
	currentVertex := Vertex{x: 0, y: 0}
	grid = append(grid, currentVertex)
	foundHq := false
	for _, instruction := range dataArray {
		direction := getDirection(instruction)
		steps := getSteps(instruction)
		if direction == "R" {
			currentDirection = currentDirection.right
		} else {
			currentDirection = currentDirection.left
		}
		for i := 0; i < steps; i++ {
			currentDirection.steps = currentDirection.steps + 1
			newVertex := currentVertex
			switch currentDirection {
			case North:
				foundHq = checkIfAlreadyVisited(grid, Vertex{x: currentVertex.x, y: currentVertex.y + 1}, foundHq)
				newVertex.y += 1
			case East:
				foundHq = checkIfAlreadyVisited(grid, Vertex{x: currentVertex.x + 1, y: currentVertex.y}, foundHq)
				newVertex.x += 1
			case South:
				foundHq = checkIfAlreadyVisited(grid, Vertex{x: currentVertex.x, y: currentVertex.y - 1}, foundHq)
				newVertex.y -= 1
			default:
				foundHq = checkIfAlreadyVisited(grid, Vertex{x: currentVertex.x - 1, y: currentVertex.y}, foundHq)
				newVertex.x -= 1
			}
			if foundHq {
				break
			} else {
				grid = append(grid, newVertex)
				currentVertex = newVertex
			}
		}
		if foundHq {
			break
		}
	}

	fmt.Printf("north: %v, east: %v, south: %v, west: %v\n", North.steps, East.steps, South.steps, West.steps)
	fmt.Printf("blocks away: %v\n", math.Abs(float64(North.steps-South.steps))+math.Abs(float64(East.steps-West.steps)))
}
func checkIfAlreadyVisited(grid []Vertex, checkVertex Vertex, foundHq bool) bool {
	for _, vertex := range grid {
		if vertex.x == checkVertex.x && vertex.y == checkVertex.y {
			foundHq = true
			break
		}
	}
	return foundHq
}
func getSteps(instruction string) int {
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(err)
	} else {
		return steps
	}
}
func getDirection(instruction string) string {
	return instruction[0:1]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Direction struct {
	right *Direction
	left  *Direction
	steps int
}

type Vertex struct {
	x int
	y int
}
