package main

import (
	"fmt"
	"time"
)

func NextWorkday(start time.Time) time.Time {
	nextDay := start.AddDate(0, 0, 1)
	switch nextDay.Weekday() {
	case time.Saturday:
		return nextDay.AddDate(0, 0, 2)
	case time.Sunday:
		return nextDay.AddDate(0, 0, 1)
	default:
		return nextDay
	}
}

func TimeAgo(pastTime time.Time) string {
	now := time.Now()
	duration := now.Sub(pastTime)

	if duration < 0 {
		return "in the future"
	}

	seconds := int(duration.Seconds())
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	months := days / 30
	years := days / 365

	switch {
	case years > 0:
		if years == 1 {
			return "1 year ago"
		}
		return fmt.Sprintf("%d years ago", years)
	case months > 0:
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	case days > 0:
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	case hours > 0:
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	case minutes > 0:
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	default:
		if seconds == 1 {
			return "1 second ago"
		}
		return fmt.Sprintf("%d seconds ago", seconds)
	}
}

func ParseStringToTime(dateString, format string) (time.Time, error) {
	return time.Parse(format, dateString)
}

func TimeDifference(start, end time.Time) time.Duration {
	return end.Sub(start)
}

func FormatTimeToString(timestamp time.Time, format string) string {
	return timestamp.Format(format)
}

func main() {
	dateString := "2023-10-23 02:41:49"
	format := "2006-01-02 15:04:05"
	result, err := ParseStringToTime(dateString, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: 2023-10-23 02:41:49 +0000 UTC

	pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
	result1 := TimeAgo(pastTime)
	fmt.Println(result1) // Output: 1 month ago
}
