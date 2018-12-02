package main

import (
	"fmt"
)

func main() {

	// // DAY ONE
	// parsedDayOne := parseCalibrations()
	// dayOnePuzzleOne := endingFrequency(parsedDayOne)
	// fmt.Println(dayOnePuzzleOne)
	// dayOnePuzzleTwo := repeatedFrequency(parsedDayOne)
	// fmt.Println(dayOnePuzzleTwo)

	// DAY TWO
	parsedDayTwo := parsePackageIDs()
	dayTwoPuzzleOne := checkSum(parsedDayTwo)
	fmt.Println(dayTwoPuzzleOne)
	dayTwoPuzzleTwo := commonLetters(parsedDayTwo)
	fmt.Println(dayTwoPuzzleTwo)

}
