package utils

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"time"
)

func PopSlashSource(selection *goquery.Selection) string {
	res, _ := selection.Attr("src")
	split := strings.Split(res, "/")
	return split[len(split)-1]
}

func UnixTimeStringToTime(unixTime string) time.Time {
	if len(unixTime) == 0 {
		return time.Unix(0, 0)
	}
	dateStartInt64, _ := strconv.ParseInt(unixTime[:len(unixTime)-3], 10, 64)
	return time.Unix(dateStartInt64, 0)
}
