package puzzles

import (
	"sort"
	"strings"
	"time"
)

type Event struct {
	Kind string
	Time time.Time
}

func BestMinute(events []Event, guard string) int {

	sort.Slice(events, func(i, j int) bool {
		return events[i].Time.Before(events[j].Time)
	})

	var currentGuard string
	var asleepTime time.Time
	sleepingMinutes := make(map[int]int)

	for _, e := range events {
		if strings.Contains(e.Kind, "begins shift") {
			currentGuard = e.Kind
		} else if strings.Contains(e.Kind, "wakes up") {

			if currentGuard == guard {
				currMin := asleepTime.Minute()
				endMin := e.Time.Minute()
				for currMin < endMin {
					sleepingMinutes[currMin]++
					currMin++

				}
			}

		} else if strings.Contains(e.Kind, "falls asleep") {
			asleepTime = e.Time
		}

	}

	var mostSleep int
	var timesAsleep int

	for min, asleep := range sleepingMinutes {
		if asleep > timesAsleep {
			mostSleep = min
			timesAsleep = asleep
		}
	}

	return mostSleep
}

func MostAsleepGuard(events []Event) string {

	sort.Slice(events, func(i, j int) bool {
		return events[i].Time.Before(events[j].Time)
	})

	var currentGuard string
	var asleepTime time.Time
	sleepingMinutes := make(map[string]time.Duration)

	for _, e := range events {
		if strings.Contains(e.Kind, "begins shift") {
			currentGuard = e.Kind
		} else if strings.Contains(e.Kind, "wakes up") {
			sleepMin := e.Time.Sub(asleepTime)
			sleepingMinutes[currentGuard] += sleepMin

		} else if strings.Contains(e.Kind, "falls asleep") {
			asleepTime = e.Time
		}

	}

	var mostSleep time.Duration
	var sleepyGuard string
	for gd, dur := range sleepingMinutes {
		if dur > mostSleep {
			mostSleep = dur
			sleepyGuard = gd
		}
	}

	return sleepyGuard
}

func GuardAsleepMostAtMinute(events []Event) (string, int) {

	sort.Slice(events, func(i, j int) bool {
		return events[i].Time.Before(events[j].Time)
	})

	var currentGuard string
	var asleepTime time.Time
	sleepingMinutes := make(map[int][]string)
	for i := 0; i < 61; i++ {
		sleepingMinutes[i] = []string{}
	}

	for _, e := range events {
		if strings.Contains(e.Kind, "begins shift") {
			currentGuard = e.Kind
		} else if strings.Contains(e.Kind, "wakes up") {

			currMin := asleepTime.Minute()
			endMin := e.Time.Minute()
			for currMin < endMin {
				sleepingMinutes[currMin] = append(sleepingMinutes[currMin], currentGuard)
				currMin++
			}

		} else if strings.Contains(e.Kind, "falls asleep") {
			asleepTime = e.Time
		}

	}

	var timesAsleep int
	var sleepyGuard string
	var minute int
	for min, guards := range sleepingMinutes {

		occurrances := make(map[string]int)
		for _, gd := range guards {
			occurrances[gd]++
		}

		for gd, occ := range occurrances {
			if timesAsleep < occ {
				timesAsleep = occ
				sleepyGuard = gd
				minute = min
			}
		}

	}

	return sleepyGuard, minute
}
