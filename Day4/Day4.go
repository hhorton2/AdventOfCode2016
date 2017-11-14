package main

import (
	"io/ioutil"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("Day4/input.txt")
	check(err)
	datString := string(dat)
	rooms := getRooms(datString)
	for _,room := range rooms{
		fmt.Printf("%v | %v | %v\n", room.encryptedName, room.sectorId, room.checksum)
	}
}
func getRooms(datString string) ([]Room) {
	regex := regexp.MustCompile(".*\n")
	matches := regex.FindAllString(datString, -1)
	var rooms []Room
	regex = regexp.MustCompile("\n")
	for i, match := range matches {
		matches[i] = regex.ReplaceAllString(match, "")
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
