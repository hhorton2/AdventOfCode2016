package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	dat, err := ioutil.ReadFile("Day07/input.txt")
	check(err)
	datString := string(dat)
	rowStream := make(chan string, 10)
	go getRows(datString, rowStream)
	ipStream := make(chan ipv7, 10)
	go parseRows(rowStream, ipStream)
	ipStreamA := make(chan ipv7, 10)
	ipStreamB := make(chan ipv7, 10)
	go duplicateStream(ipStream, ipStreamA, ipStreamB)
	tlsResultStream := make(chan int)
	sslResultStream := make(chan int)
	go getValidNumberOfTLSAddresses(ipStreamA, tlsResultStream)
	go getValidNumberOfSSLAddresses(ipStreamB, sslResultStream)
	fmt.Printf("Valid TLS Addresses: %v\nValid SSL Addresses: %v\nFinished in: %v", <-tlsResultStream, <-sslResultStream, time.Since(start))

}
func getValidNumberOfSSLAddresses(ipStream <-chan ipv7, resultStream chan<- int) {
	defer close(resultStream)
	validAddresses := 0
	for ip := range ipStream {
		ipAba := findAba(ip.ipSlices)
		hypernetAba := findAba(ip.hypernets)
	verify:
		for _, aba := range ipAba {
			for _, bab := range hypernetAba {
				if len(aba) > 0 && len(bab) > 0 && aba[0] == bab[1] && aba[1] == bab[0] {
					validAddresses++
					break verify
				}
			}
		}
	}
	resultStream <- validAddresses
}
func findAba(inputs []string) []string {
	aba := make([]string, 1, 8)
	for _, input := range inputs {
		for i := 2; i < len(input); i++ {
			if input[i] == input[i-2] && input[i] != input[i-1] {
				aba = append(aba, input[i-2:i+1])
			}
		}
	}
	return aba
}
func getValidNumberOfTLSAddresses(ipStream <-chan ipv7, resultStream chan<- int) {
	defer close(resultStream)
	validAddresses := 0
	for ip := range ipStream {
		ipAbbaFound := findAbba(ip.ipSlices)
		hypernetAbbaFound := findAbba(ip.hypernets)
		if ipAbbaFound && !hypernetAbbaFound {
			validAddresses++
		}
	}
	resultStream <- validAddresses
}
func findAbba(inputs []string) bool {
	for _, value := range inputs {
		for i := 3; i < len(value); i++ {
			if value[i-3] == value[i] && value[i-1] == value[i-2] && value[i] != value[i-1] {
				return true
			}
		}
	}
	return false
}
func parseRows(rowStream <-chan string, ipStream chan<- ipv7) {
	defer close(ipStream)

	for row := range rowStream {
		inHypernet := false
		currentIp := ipv7{hypernets: make([]string, 1, 4), ipSlices: make([]string, 1, 4)}
		currentHypernet := 0
		currentIpSlice := 0
		for _, char := range row {
			letter := string(char)
			if letter == "[" {
				inHypernet = true
				currentIpSlice++
				if currentIpSlice >= len(currentIp.ipSlices) {
					currentIp.ipSlices = append(currentIp.ipSlices, "")
				}
			} else if letter == "]" {
				inHypernet = false
				currentHypernet++
				if currentHypernet >= len(currentIp.hypernets) {
					currentIp.hypernets = append(currentIp.hypernets, "")
				}
			} else if inHypernet {
				currentIp.hypernets[currentHypernet] += letter
			} else {
				currentIp.ipSlices[currentIpSlice] += letter
			}
		}
		ipStream <- currentIp
	}
}
func getRows(input string, rowStream chan<- string) {
	defer close(rowStream)
	for _, row := range strings.Split(input, "\n") {
		rowStream <- row
	}
}

func duplicateStream(ipAddresses <-chan ipv7, forkA chan<- ipv7, forkB chan<- ipv7) {
	defer close(forkA)
	defer close(forkB)
	for position := range ipAddresses {
		forkA <- position
		forkB <- position
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ipv7 struct {
	hypernets, ipSlices []string
}
