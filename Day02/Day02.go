package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("Day02/input.txt")
	check(err)
	dataArray := strings.Split(string(dat), "\n")
	dataArray = dataArray[:len(dataArray)-1]
	one := new(KeypadNumber)
	two := new(KeypadNumber)
	three := new(KeypadNumber)
	four := new(KeypadNumber)
	five := new(KeypadNumber)
	six := new(KeypadNumber)
	seven := new(KeypadNumber)
	eight := new(KeypadNumber)
	nine := new(KeypadNumber)
	a := new(KeypadNumber)
	b := new(KeypadNumber)
	c := new(KeypadNumber)
	d := new(KeypadNumber)
	initializeDayOneRelationships(one, two, four, three, five, six, seven, eight, nine)
	solve(five, dataArray)
	initializeDayTwoRelationships(one, two, four, three, five, six, seven, eight, nine, a, b, c, d)
	solve(five, dataArray)
}
func solve(startingKey *KeypadNumber, instructions []string) {
	currentKey := startingKey
	for _, instructionSet := range instructions {
		for _, instruction := range strings.Split(instructionSet, "") {
			switch instruction {
			case "L":
				currentKey = currentKey.left
			case "R":
				currentKey = currentKey.right
			case "U":
				currentKey = currentKey.up
			case "D":
				currentKey = currentKey.down
			}
		}
		fmt.Print(currentKey.number)
	}
	fmt.Println()
}

func initializeDayOneRelationships(one *KeypadNumber, two *KeypadNumber, four *KeypadNumber, three *KeypadNumber, five *KeypadNumber, six *KeypadNumber, seven *KeypadNumber, eight *KeypadNumber, nine *KeypadNumber) {
	one.left = one
	one.right = two
	one.up = one
	one.down = four
	two.left = one
	two.right = three
	two.up = two
	two.down = five
	three.left = two
	three.right = three
	three.up = three
	three.down = six
	four.left = four
	four.right = five
	four.up = one
	four.down = seven
	five.left = four
	five.right = six
	five.up = two
	five.down = eight
	six.left = five
	six.right = six
	six.up = three
	six.down = nine
	seven.left = seven
	seven.right = eight
	seven.up = four
	seven.down = seven
	eight.left = seven
	eight.right = nine
	eight.up = five
	eight.down = eight
	nine.left = eight
	nine.right = nine
	nine.up = six
	nine.down = nine
	one.number = "1"
	two.number = "2"
	three.number = "3"
	four.number = "4"
	five.number = "5"
	six.number = "6"
	seven.number = "7"
	eight.number = "8"
	nine.number = "9"

}

func initializeDayTwoRelationships(one *KeypadNumber, two *KeypadNumber, four *KeypadNumber, three *KeypadNumber, five *KeypadNumber, six *KeypadNumber, seven *KeypadNumber, eight *KeypadNumber, nine *KeypadNumber, a *KeypadNumber, b *KeypadNumber, c *KeypadNumber, d *KeypadNumber) {
	one.left = one
	one.right = one
	one.up = one
	one.down = three
	two.left = two
	two.right = three
	two.up = two
	two.down = six
	three.left = two
	three.right = four
	three.up = one
	three.down = seven
	four.left = three
	four.right = four
	four.up = four
	four.down = eight
	five.left = five
	five.right = six
	five.up = five
	five.down = five
	six.left = five
	six.right = seven
	six.up = two
	six.down = a
	seven.left = six
	seven.right = eight
	seven.up = three
	seven.down = b
	eight.left = seven
	eight.right = nine
	eight.up = four
	eight.down = c
	nine.left = eight
	nine.right = nine
	nine.up = nine
	nine.down = nine
	a.left = a
	a.right = b
	a.up = six
	a.down = a
	b.left = a
	b.right = c
	b.up = seven
	b.down = d
	c.left = b
	c.right = c
	c.up = eight
	c.down = c
	d.left = d
	d.right = d
	d.up = b
	d.down = d
	a.number = "a"
	b.number = "b"
	c.number = "c"
	d.number = "d"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type KeypadNumber struct {
	number string
	left   *KeypadNumber
	right  *KeypadNumber
	up     *KeypadNumber
	down   *KeypadNumber
}
