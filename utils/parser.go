package utils

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func PopSlashSource(selection *goquery.Selection) string {
	res, _ := selection.Attr("src")
	split := strings.Split(res, "/")
	return split[len(split)-1]
}
