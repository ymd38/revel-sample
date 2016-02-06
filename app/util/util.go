package util

import "time"

const (
	DATETIME_FORMAT = "2006-01-02 15:04:05"
	DAY_FORMAT      = "2006-01-02"
)

func UnitTimeToDateString(unixtime int64) string {
	return time.Unix(unixtime, 0).Format(DATETIME_FORMAT)
}

func UnitTimeToDayString(unixtime int64) string {
	return time.Unix(unixtime, 0).Format(DAY_FORMAT)
}
