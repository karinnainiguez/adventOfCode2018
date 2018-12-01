package main

import (
	"fmt"
)

func main() {

	parsedDayOne := parseCalibrations()
	dayOnePuzzleOne := endingFrequency(parsedDayOne)
	fmt.Println(dayOnePuzzleOne)

	dayOnePuzzleTwo := repeatedFrequency(parsedDayOne)
	fmt.Println(dayOnePuzzleTwo)
}
