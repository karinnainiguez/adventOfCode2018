package main

import (
	"fmt"

	"github.com/karinnainiguez/adventOfCode2018/puzzles"
)

func main() {

	// // DAY ONE
	// parsedDayOne := parseCalibrations()
	// dayOnePuzzleOne := puzzles.EndingFrequency(parsedDayOne)
	// fmt.Println(dayOnePuzzleOne)
	// dayOnePuzzleTwo := puzzles.RepeatedFrequency(parsedDayOne)
	// fmt.Println(dayOnePuzzleTwo)

	// // DAY TWO
	// parsedDayTwo := parsePackageIDs()
	// dayTwoPuzzleOne := puzzles.CheckSum(parsedDayTwo)
	// fmt.Println(dayTwoPuzzleOne)
	// dayTwoPuzzleTwo := puzzles.CommonLetters(parsedDayTwo)
	// fmt.Println(dayTwoPuzzleTwo)

	// // DAY Three
	// parsedDayThree := parseClaimAreas()
	// dayThreePuzzleOne := puzzles.OverlappingFabric(parsedDayThree)
	// fmt.Println(dayThreePuzzleOne)
	// dayThreePuzzleTwo := puzzles.NonOverlappingClaim(parsedDayThree)
	// fmt.Println(dayThreePuzzleTwo)

	// // DAY Four
	// parsedDayFour := parseDayFour()
	// dayFourPuzzleOneGuard := puzzles.MostAsleepGuard(parsedDayFour)
	// fmt.Println(dayFourPuzzleOneGuard)
	// dayFourPuzzleOneMinute := puzzles.BestMinute(parsedDayFour, dayFourPuzzleOneGuard)
	// fmt.Println(dayFourPuzzleOneMinute)
	// guard, min := puzzles.GuardAsleepMostAtMinute(parsedDayFour)
	// fmt.Printf("   Guard: %v    \n   Minute: %v\n", guard, min)

	// DAY Five
	parsedDayFive := parsedayFive()
	// fmt.Println(parsedDayFive)
	fmt.Println(len(parsedDayFive))
	// dayFivePuzzleOne := puzzles.DayFiveOne(parsedDayFive)
	// fmt.Println(len(dayFivePuzzleOne))
	// fmt.Println(dayFivePuzzleOne)
	dayFivePuzzleTwo := puzzles.DayFiveTwo(parsedDayFive)
	fmt.Println(dayFivePuzzleTwo)
	fmt.Println(len(dayFivePuzzleTwo))

}
