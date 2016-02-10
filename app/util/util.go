package util

import "time"

const (
	DATETIME_FORMAT = "2006-01-02 15:04:05"
	DAY_FORMAT      = "2006-01-02"
	DAY_FORMAT2     = "20060102"
)

func UnixTimeToDateString(unixtime int64) string {
	return time.Unix(unixtime, 0).Format(DATETIME_FORMAT)
}

func UnixTimeToDayString(unixtime int64) string {
	return time.Unix(unixtime, 0).Format(DAY_FORMAT)
}

func DayStringToUnixTime(day string) int64 {
	unixtime, _ := time.Parse(DAY_FORMAT2, day)
	return unixtime.Unix()
}
