package util

import "time"

const (
	DATETIME_FORMAT = "2006-01-02 15:04:05"
)

func UnitTimeToString(unixtime int64) string {
	return time.Unix(unixtime, 0).Format(DATETIME_FORMAT)
}
