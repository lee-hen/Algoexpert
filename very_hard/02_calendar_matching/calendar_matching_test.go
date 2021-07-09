package calendar_matching

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	calendar1 := []StringMeeting{{"9:00", "10:30"}, {"12:00", "13:00"}, {"16:00", "18:00"}}
	dailyBounds1 := StringMeeting{"9:00", "20:00"}
	calendar2 := []StringMeeting{{"10:00", "11:30"}, {"12:30", "14:30"}, {"14:30", "15:00"}, {"16:00", "17:00"}}
	dailyBounds2 := StringMeeting{"10:00", "18:30"}
	meetingDuration := 30
	expected := []StringMeeting{{"11:30", "12:00"}, {"15:00", "16:00"}, {"18:00", "18:30"}}
	result := CalendarMatching(calendar1, dailyBounds1, calendar2, dailyBounds2, meetingDuration)
	require.Equal(t, result, expected)
}

func TestCase2(t *testing.T) {
	calendar1 := []StringMeeting{{"10:00", "10:30"}, {"10:45", "11:15"}, {"11:30", "13:00"}, {"14:15", "16:00"}, {"16:00", "18:00"}}
	dailyBounds1 := StringMeeting{"9:30", "20:00"}
	calendar2 := []StringMeeting{{"10:00", "11:00"}, {"12:30", "13:30"}, {"14:30", "15:00"}, {"16:00", "17:00"}}
	dailyBounds2 := StringMeeting{"9:00", "18:30"}
	meetingDuration := 15
	expected := []StringMeeting{{"9:30", "10:00"}, {"11:15", "11:30"}, {"13:30", "14:15"}, {"18:00", "18:30"}}
	result := CalendarMatching(calendar1, dailyBounds1, calendar2, dailyBounds2, meetingDuration)
	require.Equal(t, expected, result)
}

func TestCase3(t *testing.T) {
	calendar1 := []StringMeeting{}
	dailyBounds1 := StringMeeting{"9:30", "20:00"}
	calendar2 := []StringMeeting{}
	dailyBounds2 := StringMeeting{"9:00", "18:30"}
	meetingDuration := 15
	expected := []StringMeeting{}
	result := CalendarMatching(calendar1, dailyBounds1, calendar2, dailyBounds2, meetingDuration)
	require.Equal(t, expected, result)
}
