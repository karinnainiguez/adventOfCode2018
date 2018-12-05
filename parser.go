package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/karinnainiguez/adventOfCode2018/puzzles"
)

func parseCalibrations() []int {
	csvFile, _ := os.Open("inputs/dayOneInput.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var calibrations []int
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		num, _ := strconv.Atoi(line[0])
		calibrations = append(calibrations, num)
	}
	return calibrations[1:]
}

func parsePackageIDs() []string {
	csvFile, _ := os.Open("inputs/dayTwoInput.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var packageIDs []string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		ID := line[0]
		packageIDs = append(packageIDs, ID)
	}
	return packageIDs[1:]
}

func parseClaimAreas() [][]string {
	csvFile, _ := os.Open("inputs/dayThreeInput.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var claimAreas [][]string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		claimAreas = append(claimAreas, line)
	}
	return claimAreas[1:]
}

// type event struct {
// 	kind string
// 	time time.Time
// }

func parseDayFour() []puzzles.Event {
	csvFile, _ := os.Open("inputs/dayFourInput.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var events []string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		events = append(events, line[0])
	}
	events = events[1:]
	newEvents := []puzzles.Event{}
	for _, e := range events {

		str := strings.Split(e, "]")
		// fmt.Println(strings.Replace(str[0], "[", "", 1))
		tm, _ := time.Parse("2006-01-02 15:04", strings.Replace(str[0], "[", "", 1))

		// fmt.Println(tm.String())
		ne := puzzles.Event{
			Kind: str[1],
			Time: tm,
		}
		newEvents = append(newEvents, ne)

	}

	return newEvents
}

func parsedayFive() string {
	csvFile, _ := os.Open("inputs/dayFiveInput.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var claimAreas []string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		claimAreas = append(claimAreas, line[0])
	}
	return claimAreas[1]
}
