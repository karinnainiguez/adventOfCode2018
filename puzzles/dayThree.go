package puzzles

import (
	"strconv"
	"strings"
)

func OverlappingFabric(claimColl [][]string) int {
	squaresWithinTwoPlus := 0
	fullFabric := make([][]string, 1100)

	for i := 0; i < 1100; i++ {
		fullFabric[i] = make([]string, 1100)
	}

	for _, claim := range claimColl {

		split := strings.Split(claim[0], " @ ")
		startCol, _ := strconv.Atoi(split[1])
		split = strings.Split(claim[1], ": ")
		startRow, _ := strconv.Atoi(split[0])
		split = strings.Split(split[1], "x")
		width, _ := strconv.Atoi(split[0])
		height, _ := strconv.Atoi(split[1])

		for r := 0; r < height; r++ {

			for c := 0; c < width; c++ {
				current := fullFabric[startRow+r][startCol+c]

				if current == "" {
					fullFabric[startRow+r][startCol+c] = "taken"
				} else if current == "taken" {
					fullFabric[startRow+r][startCol+c] = "X"
				}

			}

		}

	}

	for _, row := range fullFabric {
		for _, col := range row {
			if col == "X" {
				squaresWithinTwoPlus++
			}
		}
	}

	return squaresWithinTwoPlus
}

func NonOverlappingClaim(claimColl [][]string) string {
	fullFabric := make([][]string, 1100)

	for i := 0; i < 1100; i++ {
		fullFabric[i] = make([]string, 1100)
	}

	for _, claim := range claimColl {

		split := strings.Split(claim[0], " @ ")
		startCol, _ := strconv.Atoi(split[1])
		split = strings.Split(claim[1], ": ")
		startRow, _ := strconv.Atoi(split[0])
		split = strings.Split(split[1], "x")
		width, _ := strconv.Atoi(split[0])
		height, _ := strconv.Atoi(split[1])

		for r := 0; r < height; r++ {

			for c := 0; c < width; c++ {
				current := fullFabric[startRow+r][startCol+c]

				if current == "" {
					fullFabric[startRow+r][startCol+c] = "taken"
				} else if current == "taken" {
					fullFabric[startRow+r][startCol+c] = "X"
				}

			}

		}

	}

	for _, claim := range claimColl {

		split := strings.Split(claim[0], " @ ")
		id := split[0]
		startCol, _ := strconv.Atoi(split[1])
		split = strings.Split(claim[1], ": ")
		startRow, _ := strconv.Atoi(split[0])
		split = strings.Split(split[1], "x")
		width, _ := strconv.Atoi(split[0])
		height, _ := strconv.Atoi(split[1])

		overlaps := false

		for r := 0; r < height; r++ {

			for c := 0; c < width; c++ {
				current := fullFabric[startRow+r][startCol+c]

				if current == "X" {
					overlaps = true
				}

			}

		}
		if overlaps == false {
			return id
		}

	}

	return ""

}
