package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"time"
)

const (
	DATETIME_FORMAT = "2006-01-02 15:04:05"
	DAY_FORMAT      = "2006-01-02"
	DAY_FORMAT2     = "20060102"
)

func UnixTimeToDateString(unixtime int64) string {
	return time.Unix(unixtime, 0).Format(DATETIME_FORMAT)
}

func UnixTimeToDayString(unixtime int64) string {
	return time.Unix(unixtime, 0).Format(DAY_FORMAT2)
}

func DayStringToUnixTime(day string) int64 {
	unixtime, _ := time.Parse(DAY_FORMAT2, day)
	return unixtime.Unix()
}

func ToMD5(src string) string {
	h := md5.New()
	io.WriteString(h, src)
	return hex.EncodeToString(h.Sum(nil))
}
