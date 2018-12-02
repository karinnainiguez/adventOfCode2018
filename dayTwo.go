package main

import "strings"

func checkSum(packageIDs []string) int {

	check2 := 0
	check3 := 0

	for _, p := range packageIDs {
		occ := make(map[rune]int)
		for _, char := range p {
			if _, ok := occ[char]; ok {
				occ[char]++
			} else {
				occ[char] = 1
			}
		}
		local2 := 0
		local3 := 0
		for _, o := range occ {
			if o == 2 {
				local2 = 1
			} else if o == 3 {
				local3 = 1
			}
		}
		check2 += local2
		check3 += local3

	}

	return check2 * check3
}

func commonLetters(packageIDs []string) string {
	str1, str2 := similarStrings(packageIDs)
	common := commonStrings(str1, str2)
	return common
}

func similarStrings(strs []string) (str1 string, str2 string) {

	for p, s := range strs {
		str1 = s

		for np := p + 1; np < len(strs); np++ {
			str2 = strs[np]
			difference := 0
			if len(str1) == len(str2) {

				for i := 0; i < len(str1); i++ {
					if str1[i] != str2[i] {
						difference++
					}
				}

				if difference == 1 {
					return str1, str2
				}

			}

		}
	}
	return
}
func commonStrings(str1 string, str2 string) string {

	var newString strings.Builder

	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			newString.WriteByte(str1[i])
		}
	}

	return newString.String()
}
