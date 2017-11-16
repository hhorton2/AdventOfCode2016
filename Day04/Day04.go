package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	dat, err := ioutil.ReadFile("Day4/input.txt")
	check(err)
	datString := string(dat)
	rooms := getRooms(datString)
	sectorSum := solveDayOne(rooms)
	fmt.Println(sectorSum)
	rooms = getRoomsShifted(datString)
	solveDayTwo(rooms)
}
func solveDayTwo(rooms []Room) {
	for _, room := range rooms {
		if strings.Contains(room.encryptedName, "north") {
			fmt.Printf("%v | %v\n", room.encryptedName, room.sectorId)
		}
	}
}
func solveDayOne(rooms []Room) int {
	wg := new(sync.WaitGroup)
	sectorChan := make(chan int, len(rooms))
	for _, room := range rooms {
		wg.Add(1)
		go getSectorIfValid(room, wg, sectorChan)
	}
	wg.Wait()
	close(sectorChan)
	sectorSum := 0
	for sector := range sectorChan {
		sectorSum += sector
	}
	return sectorSum
}
func getSectorIfValid(room Room, wg *sync.WaitGroup, sectors chan<- int) {
	defer wg.Done()
	letters := getMapFromRoomName(room.encryptedName)
	actualChecksum := getChecksumFromMap(letters)
	if actualChecksum == room.checksum {
		sectors <- room.sectorId
	}
}
func getMapFromRoomName(roomName string) map[rune]int {
	letters := map[rune]int{}
	for _, currentRune := range roomName {
		letters[currentRune]++
	}
	return letters
}

func getChecksumFromMap(letters map[rune]int) string {
	var ss []kv
	for k, v := range letters {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})
	var str string
	for i := 0; i < 5; i++ {
		str += string(ss[i].Key)
	}
	return str
}
func getRooms(datString string) []Room {
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
func getRoomsShifted(datString string) []Room {
	regex := regexp.MustCompile(".*\n")
	matches := regex.FindAllString(datString, -1)
	var rooms []Room
	for _, match := range matches {
		checksum := getChecksum(match)
		sector := getSectorId(match)
		roomName := getRoomNameShifted(match, sector)
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
func getRoomNameShifted(match string, sector int) string {
	roomRegex := regexp.MustCompile("-\\d+.*\n")
	roomName := roomRegex.ReplaceAllString(match, "")
	var shiftedRoomName string
	for _, currentRune := range roomName {
		letterByte := byte(currentRune)
		for i := 0; i < sector; i++ {
			if letterByte == 'z' {
				letterByte = 'a'
			} else {
				letterByte = letterByte + 1
			}
		}
		shiftedRoomName += string(letterByte)
	}
	return shiftedRoomName
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
	checksum = checksum[1 : len(checksum)-1]
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
	Key   rune
	Value int
}
