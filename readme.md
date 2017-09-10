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

```go
res:=[]string{"the","result","you","want" ,"to","show","in","the","bar"}
af := alfredV3.New(res)
af.FormatAndPrint()
```
