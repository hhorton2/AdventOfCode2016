package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("Day5/input.txt")
	check(err)
	datString := string(dat)
	datString = strings.Replace(datString, "\n", "", -1)
	password := solveDayOne(datString)
	fmt.Println(password)
	password = solveDayTwo(datString)
	fmt.Println(password)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solveDayOne(input string) string {
	password := ""
	for i := 0; len(password) < 8; i++ {
		h := md5.New()
		byteInput := []byte(input)
		byteInput = append(byteInput, []byte(strconv.Itoa(i))...)
		h.Write(byteInput)
		hash := hex.EncodeToString(h.Sum(nil))
		firstFive := hash[0:5]
		if firstFive == "00000" {
			password += hash[5:6]
			//fmt.Printf("%v", string(hash))
		}
	}
	return password
}

func solveDayTwo(input string) string {
	var password [8]byte
	var oneFound bool
	var twoFound bool
	var threeFound bool
	var fourFound bool
	var fiveFound bool
	var sixFound bool
	var sevenFound bool
	var eightFound bool
	var allFound bool
	for i := 0; !allFound; i++ {
		h := md5.New()
		byteInput := []byte(input)
		byteInput = append(byteInput, []byte(strconv.Itoa(i))...)
		h.Write(byteInput)
		hash := hex.EncodeToString(h.Sum(nil))
		firstFive := hash[0:5]
		if firstFive == "00000" {
			position := hash[5:6]
			value := hash[6:7]
			if position == "0" && !oneFound {
				password[0] = []byte(value)[0]
				oneFound = true
			} else if position == "1" && !twoFound {
				password[1] = []byte(value)[0]
				twoFound = true
			} else if position == "2" && !threeFound {
				password[2] = []byte(value)[0]
				threeFound = true
			} else if position == "3" && !fourFound {
				password[3] = []byte(value)[0]
				fourFound = true
			} else if position == "4" && !fiveFound {
				password[4] = []byte(value)[0]
				fiveFound = true
			} else if position == "5" && !sixFound {
				password[5] = []byte(value)[0]
				sixFound = true
			} else if position == "6" && !sevenFound {
				password[6] = []byte(value)[0]
				sevenFound = true
			} else if position == "7" && !eightFound {
				password[7] = []byte(value)[0]
				eightFound = true
			}
			allFound = oneFound && twoFound && threeFound && fourFound && fiveFound && sixFound && sevenFound && eightFound
		}
	}
	return string(password[:8])
}
