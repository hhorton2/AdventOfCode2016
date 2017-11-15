package main

import (
	"io/ioutil"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sort"
)

func main() {
	dat, err := ioutil.ReadFile("Day4/input.txt")
	check(err)
	datString := string(dat)
	rooms := getRooms(datString)
	wg := new(sync.WaitGroup)
	sectorChan := make(chan int, len(rooms))
	for _, room := range rooms {
		//fmt.Printf("%v | %v | %v\n", room.encryptedName, room.sectorId, room.checksum)
		wg.Add(1)
		go getSectorIfValid(room, wg, sectorChan)
	}
	wg.Wait()
	close(sectorChan)
	sectorSum := 0
	for sector := range sectorChan {
		sectorSum += sector
	}
	fmt.Println(sectorSum)
}
func getSectorIfValid(room Room, wg *sync.WaitGroup, sectors chan<- int) {
	defer wg.Done()
	letters := getMapFromRoomName(room.encryptedName)
	actualChecksum := getChecksumFromMap(letters)
	if actualChecksum == room.checksum {
		sectors <- room.sectorId
	}
}
func getMapFromRoomName(roomName string) map[string]int {
	letters := map[string]int{}
	for _, currentRune := range roomName {
		currentLetter := strings.Replace(strconv.QuoteRune(currentRune), "'", "", -1)
		letters[currentLetter] = letters[currentLetter] + 1
	}
	return letters
}

func getChecksumFromMap(letters map[string]int) string {
	var ss []kv
	for k, v := range letters {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key > ss[i].Key
		}
		return ss[i].Value > ss[j].Value
	})
	var str string
	for i := 0; i < 5; i++ {
		str += ss[i].Key
	}
	return str
}
func getRooms(datString string) ([]Room) {
	regex := regexp.MustCompile(".*\n")
	matches := regex.FindAllString(datString, -1)
	var rooms []Room
	for _, match := range matches {
		checksum := getChecksum(match)
		sector := getSectorId(match)
		roomName := getRoomName(match)
		rooms = append(rooms, Room{checksum: checksum, encryptedName: roomName, sectorId: sector})
	}
	return rooms
}
func getRoomName(match string) string {
	roomRegex := regexp.MustCompile("-\\d+.*\n")
	roomName := roomRegex.ReplaceAllString(match, "")
	roomName = strings.Replace(roomName, "-", "", -1)
	return roomName
}
func getSectorId(match string) int {
	sectorRegex := regexp.MustCompile("\\d+")
	sectorString := sectorRegex.FindAllString(match, 1)[0]
	sector, err := strconv.Atoi(sectorString)
	check(err)
	return sector
}
func getChecksum(match string) string {
	checksumRegex := regexp.MustCompile("\\[.+]")
	checksum := checksumRegex.FindAllString(match, 1)[0]
	checksum = checksum[1:len(checksum)-1]
	return checksum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Room struct {
	checksum      string
	sectorId      int
	encryptedName string
}

type kv struct {
	Key   string
	Value int
}
