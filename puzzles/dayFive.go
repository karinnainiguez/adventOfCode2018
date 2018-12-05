package puzzles

import "strings"

func DayFiveOne(polymer string) string {

	ok, strArr := helper(polymer)
	for ok == false {
		ok, strArr = helper(strArr)
	}

	return strArr
}

func helper(polymer string) (bool, string) {
	strArr := []rune(polymer)

	pos := 0
	for pos < len(polymer) {
		if pos+1 < len(polymer) && (strArr[pos+1]-strArr[pos] == 'A'-'a' || strArr[pos]-strArr[pos+1] == 'A'-'a') {
			newStr := string(strArr[:pos]) + string(strArr[pos+2:])
			return false, newStr
		}

		if pos > 1 && (strArr[pos-1]-strArr[pos] == 'A'-'a' || strArr[pos]-strArr[pos-1] == 'A'-'a') {
			newStr := string(strArr[:pos-1]) + string(strArr[pos+1:])
			return false, newStr
		}
		pos++
	}

	return true, string(strArr)

}

func DayFiveTwo(polymer string) string {

	lets := make(map[string]string)

	lets["A"] = "a"
	lets["B"] = "b"
	lets["C"] = "c"
	lets["D"] = "d"
	lets["E"] = "e"
	lets["F"] = "f"
	lets["G"] = "g"
	lets["H"] = "h"
	lets["I"] = "i"
	lets["J"] = "j"
	lets["K"] = "k"
	lets["L"] = "l"
	lets["M"] = "m"
	lets["N"] = "n"
	lets["O"] = "o"
	lets["P"] = "p"
	lets["Q"] = "q"
	lets["R"] = "r"
	lets["S"] = "s"
	lets["T"] = "t"
	lets["U"] = "u"
	lets["V"] = "v"
	lets["W"] = "w"
	lets["X"] = "x"
	lets["Y"] = "y"
	lets["Z"] = "z"

	options := make(map[string]string)

	for k, v := range lets {
		ns := strings.Replace(polymer, k, "", -1)
		ns = strings.Replace(ns, v, "", -1)
		options[k] = DayFiveOne(ns)
	}

	minString := options["A"]
	for _, str := range options {
		if len(str) < len(minString) {
			minString = str
		}
	}

	return minString
}

func helper2(polymer string, let rune) string {
	strArr := []rune(polymer)

	newStr := []rune{}
	pos := 0
	for pos < len(polymer) {
		if strArr[pos] != let && strArr[pos] != let-32 {
			newStr = append(newStr, strArr[pos])
		}

		pos++
	}

	return string(newStr)

}
