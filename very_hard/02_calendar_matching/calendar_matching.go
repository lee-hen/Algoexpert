package calendar_matching
// important question

import (
	"fmt"
	"strconv"
	"strings"
)

type StringMeeting struct {
	Start string
	End   string
}

type Meeting struct {
	Start int
	End   int
}

// CalendarMatching
// O(c1 + c2) time | O(c1 + c2) space - where c1 and c2 are the respective
// numbers of meetings in calendar1 and calendar2.
func CalendarMatching(
	calendar1 []StringMeeting, dailyBounds1 StringMeeting,
	calendar2 []StringMeeting, dailyBounds2 StringMeeting,
	meetingDuration int,
) []StringMeeting {
	updatedCalendar1 := updateCalendar(calendar1, dailyBounds1)
	updatedCalendar2 := updateCalendar(calendar2, dailyBounds2)
	mergedCalendar := mergeCalendars(updatedCalendar1, updatedCalendar2)
	flattenedCalendar := flattenCalendar(mergedCalendar)
	return getMatchingAvailabilities(flattenedCalendar, meetingDuration)
}

func updateCalendar(calendar []StringMeeting, dailyBounds StringMeeting) []Meeting {
	updatedCalendar := append([]StringMeeting{
		{Start: "0:00", End: dailyBounds.Start},
	}, calendar...)
	updatedCalendar = append(updatedCalendar, StringMeeting{
		Start: dailyBounds.End, End: "23:59",
	})

	meetings := make([]Meeting, 0)
	for _, i := range updatedCalendar {
		meetings = append(meetings, Meeting{
			Start: timeToMinutes(i.Start),
			End:   timeToMinutes(i.End),
		})
	}
	return meetings
}

func mergeCalendars(calendar1, calendar2 []Meeting) []Meeting {
	merged := make([]Meeting, 0)
	i, j := 0, 0
	for i < len(calendar1) && j < len(calendar2) {
		meeting1, meeting2 := calendar1[i], calendar2[j]
		if meeting1.Start < meeting2.Start {
			merged = append(merged, meeting1)
			i++
		} else {
			merged = append(merged, meeting2)
			j++
		}
	}

	for i < len(calendar1) {
		merged = append(merged, calendar1[i])
		i++
	}
	for j < len(calendar2) {
		merged = append(merged, calendar2[j])
		j++
	}
	return merged
}

func flattenCalendar(calendar []Meeting) []Meeting {
	flattened := []Meeting{calendar[0]}
	for i := 1; i < len(calendar); i++ {
		currentMeeting := calendar[i]
		previousMeeting := flattened[len(flattened)-1]
		if previousMeeting.End >= currentMeeting.Start {
			newPreviousMeeting := Meeting{
				Start: previousMeeting.Start,
				End:   max(previousMeeting.End, currentMeeting.End),
			}
			flattened[len(flattened)-1] = newPreviousMeeting
		} else {
			flattened = append(flattened, currentMeeting)
		}
	}
	return flattened
}

func getMatchingAvailabilities(calendar []Meeting, meetingDuration int) []StringMeeting {
	matchingAvailabilities := make([]StringMeeting, 0)
	for i := 1; i < len(calendar); i++ {
		start := calendar[i-1].End
		end := calendar[i].Start
		availabilityDuration := end - start
		if availabilityDuration >= meetingDuration {
			matchingAvailabilities = append(matchingAvailabilities, StringMeeting{
				Start: minutesToTime(start),
				End:   minutesToTime(end),
			})
		}
	}
	return matchingAvailabilities
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func timeToMinutes(time string) int {
	split := strings.SplitN(time, ":", 2)
	hours, _ := strconv.Atoi(split[0])
	minutes, _ := strconv.Atoi(split[1])
	return hours*60 + minutes
}

func minutesToTime(minutes int) string {
	hours, minutes := minutes/60, minutes%60
	return fmt.Sprintf("%d:%02d", hours, minutes)
}

// calendarMatching
// my solution
func calendarMatching(
	calendar1 []StringMeeting, dailyBounds1 StringMeeting,
	calendar2 []StringMeeting, dailyBounds2 StringMeeting,
	meetingDuration int,
) []StringMeeting {
	var commonDailyBound StringMeeting

	commonDailyBound.Start = dailyBounds1.Start
	if strings.Compare(formatTimeString(dailyBounds1.Start), formatTimeString(dailyBounds2.Start)) < 0 {
		commonDailyBound.Start = dailyBounds2.Start
	}

	commonDailyBound.End = dailyBounds1.End
	if strings.Compare(formatTimeString(dailyBounds1.End), formatTimeString(dailyBounds2.End)) > 0 {
		commonDailyBound.End = dailyBounds2.End
	}

	var merge1, merge2 []*StringMeeting
	merge1 = mergeOverlapping(calendar1)
	merge2 = mergeOverlapping(calendar2)
	calendar := mergeOverlapping(mergeCalendar(merge1, merge2))

	if len(calendar) == 0 {
		return []StringMeeting{
			commonDailyBound,
		}
	}

	potentialMatches := make([]StringMeeting, 0)
	if strings.Compare(formatTimeString(commonDailyBound.Start), formatTimeString(calendar[0].Start)) < 0 {
		matching := StringMeeting{
			Start: commonDailyBound.Start,
			End: calendar[0].Start,
		}
		potentialMatches = append(potentialMatches, matching)
	}

	for i := 1; i < len(calendar); i++ {
		j := i-1
		prev := calendar[j]
		curr := calendar[i]

		matching := StringMeeting{
			Start: prev.End,
			End: curr.Start,
		}
		potentialMatches = append(potentialMatches, matching)
	}

	if strings.Compare(formatTimeString(calendar[len(calendar)-1].End), formatTimeString(commonDailyBound.End)) < 0 {
		matching := StringMeeting{
			Start: calendar[len(calendar)-1].End,
			End: commonDailyBound.End,
		}
		potentialMatches = append(potentialMatches, matching)
	}

	matches := make([]StringMeeting, 0)
	for _, matching := range potentialMatches {
		start := matching.Start
		end := matching.End

		// 早めに開始したら、対象外になる。
		if strings.Compare(formatTimeString(start), formatTimeString(commonDailyBound.Start)) < 0 {
			continue
		}

		start = addDuration(start, meetingDuration)
		if strings.Compare(formatTimeString(end), formatTimeString(start)) < 0 {
			continue
		}

		matches = append(matches, matching)
	}

	return matches
}

func addDuration(time string, d int) string {
	times := strings.Split(time, ":")

	hour, _ := strconv.Atoi(times[0])
	minute, _ := strconv.Atoi(times[1])

	if minute + d >= 60 {
		hour++
		minute = minute+d-60
	} else {
		minute = minute+d
	}

	str := strconv.Itoa(hour) + ":" + strconv.Itoa(minute)
	if minute == 0 {
		str+="0"
	}

	return str
}

func mergeCalendar(calendar1, calendar2 []*StringMeeting) []StringMeeting {
	var i, j, k int
	newCalendar := make([]StringMeeting, len(calendar1) + len(calendar2))
	for i < len(calendar1) && j < len(calendar2) {
		if strings.Compare(formatTimeString(calendar1[i].Start), formatTimeString(calendar2[j].Start)) < 0 {
			newCalendar[k] = *calendar1[i]
			i++
		} else {
			newCalendar[k] = *calendar2[j]
			j++
		}
		k++
	}

	for i < len(calendar1) {
		newCalendar[k] = *calendar1[i]
		i++
		k++
	}

	for j < len(calendar2) {
		newCalendar[k] = *calendar1[j]
		j++
		k++
	}
	return newCalendar
}

func mergeOverlapping(calendar []StringMeeting) []*StringMeeting {
	newCalendar := make([]*StringMeeting, 0)

	for i := 0; i < len(calendar); i++ {
		intervalTime := calendar[i]
		if len(newCalendar) > 0 {
			lastIntervalTime := newCalendar[len(newCalendar)-1]
			if strings.Compare(formatTimeString(lastIntervalTime.End), formatTimeString(intervalTime.Start)) >= 0 {
				if strings.Compare(formatTimeString(lastIntervalTime.End), formatTimeString(intervalTime.End)) < 0 {
					lastIntervalTime.End = intervalTime.End
				}
			}  else {
				newCalendar = append(newCalendar, &intervalTime)
			}
		} else {
			newCalendar = append(newCalendar, &intervalTime)
		}
	}

	return newCalendar
}

func formatTimeString(time string) string {
	if len(time) == 5 {
		return time
	}

	times := strings.Split(time, ":")
	return fmt.Sprintf("%02s:", times[0]) + ":" + times[1]
}
