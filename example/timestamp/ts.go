package main

import (
	"flag"
	"time"

	"github.com/scbizu/alfredV3"
)

func main() {
	rawTime := flag.Int64("stamp", time.Now().Unix(), "parse timestamp")
	flag.Parse()

	resTimeFormatBar := time.Unix(*rawTime, 0).Format("2006-01-02 03:04:05PM")

	resTimes := []string{}
	resTimes = append(resTimes, resTimeFormatBar)

	resTimeFormatSlash := time.Unix(*rawTime, 0).Format("2006/01/02 03:04:05PM")
	resTimes = append(resTimes, resTimeFormatSlash)

	resTimeFormat24H := time.Unix(*rawTime, 0).Format("2006-01-02 15:04:05")

	resTimes = append(resTimes, resTimeFormat24H)

	resTimeFormatDetail := time.Unix(*rawTime, 0).Format("2006/01/02 15:04:05 '06 -0700")

	resTimes = append(resTimes, resTimeFormatDetail)
	af := alfredV3.New(resTimes)
	af.FormatAndPrint()
}
