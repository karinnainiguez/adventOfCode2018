package puzzles

func EndingFrequency(calibrations []int) int {
	st := 0
	for _, c := range calibrations {
		st += c
	}
	return st
}

func RepeatedFrequency(calibrations []int) int {
	seen := make(map[int]bool)
	st := 0
	seen[st] = true
	for {
		for _, c := range calibrations {
			st += c
			if _, ok := seen[st]; ok {
				return st
			}
			seen[st] = true
		}
	}
}
