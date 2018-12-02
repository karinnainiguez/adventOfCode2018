package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func parseCalibrations() []int {
	csvFile, _ := os.Open("dayOneInput.csv")
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
	csvFile, _ := os.Open("dayTwoInput.csv")
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
