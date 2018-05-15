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

	var resTimes []*alfredV3.InputMsg
	resTimes = append(resTimes, alfredV3.NewInputMsg(resTimeFormatBar, resTimeFormatBar, resTimeFormatBar))

	resTimeFormatSlash := time.Unix(*rawTime, 0).Format("2006/01/02 03:04:05PM")
	resTimes = append(resTimes, alfredV3.NewInputMsg(resTimeFormatSlash, resTimeFormatSlash, resTimeFormatSlash))

	resTimeFormat24H := time.Unix(*rawTime, 0).Format("2006-01-02 15:04:05")
	resTimes = append(resTimes, alfredV3.NewInputMsg(resTimeFormat24H, resTimeFormat24H, resTimeFormat24H))

	resTimeFormatDetail := time.Unix(*rawTime, 0).Format("2006/01/02 15:04:05 '06 -0700")

	resTimes = append(resTimes, alfredV3.NewInputMsg(resTimeFormatDetail, resTimeFormatDetail, resTimeFormatDetail))
	af := alfredV3.NewAlfredMsg(resTimes)
	af.FormatAndPrint()
}
