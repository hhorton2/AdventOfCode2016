package main

import "io/ioutil"

func main() {
	dat, err := ioutil.ReadFile("Day5/input.txt")
	check(err)
	datString := string(dat)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
