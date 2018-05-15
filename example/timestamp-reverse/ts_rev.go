package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/scbizu/alfredV3"
)

func main() {
	timeShortCut := flag.String("restamp", "now", "parse date to timestamp")
	flag.Parse()

	var t int64
	switch *timeShortCut {
	case "now", "0":
		t = time.Now().Unix()
	case "yesterday":
		t = time.Now().AddDate(0, 0, -1).Unix()

	case "tomorrow":
		t = time.Now().AddDate(0, 0, 1).Unix()

	default:
		if len(*timeShortCut) > 1 {
			//only support SECOND
			if (*timeShortCut)[0] == '+' {
				diff, err := strconv.ParseInt(string((*timeShortCut)[1:]), 10, 64)
				if err != nil {
					alfredV3.Error(err)
				}
				t = time.Now().Add(time.Duration(time.Date(0, 0, 0, 0, 0, int(diff), 0, time.UTC).Second())).Unix()
			} else if (*timeShortCut)[0] == '-' {
				diff, err := strconv.ParseInt(string((*timeShortCut)[1:]), 10, 64)
				if err != nil {
					alfredV3.Error(err)
				}
				t = time.Now().Add(time.Duration(time.Date(0, 0, 0, 0, 0, int(-diff), 0, time.UTC).Second())).Unix()
			} else {
				//specific date format
				parseTime, err := time.Parse("2006/01/02/15:04:05", *timeShortCut)
				if err != nil {
					alfredV3.Error(err)
				}
				t = parseTime.Unix()
			}
		}
	}
	nowStr := strconv.FormatInt(t, 10)
	nowStrFmt := fmt.Sprintf("%s => %s", *timeShortCut, nowStr)
	ts := []*alfredV3.InputMsg{alfredV3.NewInputMsg(nowStrFmt, nowStrFmt, nowStr)}
	af := alfredV3.NewAlfredMsg(ts)
	af.FormatAndPrint()
}
