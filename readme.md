alfredV3
---

alfredV3 is a alfred3 go util lib ,it will use json as the data exchange rule.

## feature

* esaily convert your cmd apps into alfred apps
* use boltdb as cache

## WIP feature

* support for quick open URL
* other advance Workflow features
* auto deploy app to the workflow dir

## easy usage

> A timestamp converter script filter

```go
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
```
