package main

import (
	"fmt"

	"github.com/karinnainiguez/adventOfCode2018/puzzles"
)

func main() {

	// DAY ONE
	parsedDayOne := parseCalibrations()
	dayOnePuzzleOne := puzzles.EndingFrequency(parsedDayOne)
	fmt.Println(dayOnePuzzleOne)
	dayOnePuzzleTwo := puzzles.RepeatedFrequency(parsedDayOne)
	fmt.Println(dayOnePuzzleTwo)

	// DAY TWO
	parsedDayTwo := parsePackageIDs()
	dayTwoPuzzleOne := puzzles.CheckSum(parsedDayTwo)
	fmt.Println(dayTwoPuzzleOne)
	dayTwoPuzzleTwo := puzzles.CommonLetters(parsedDayTwo)
	fmt.Println(dayTwoPuzzleTwo)

	// DAY Three
	parsedDayThree := parseClaimAreas()
	dayThreePuzzleOne := puzzles.OverlappingFabric(parsedDayThree)
	fmt.Println(dayThreePuzzleOne)
	dayThreePuzzleTwo := puzzles.NonOverlappingClaim(parsedDayThree)
	fmt.Println(dayThreePuzzleTwo)

}
