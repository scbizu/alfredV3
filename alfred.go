package alfredV3

import (
	"encoding/json"
	"fmt"
)

// InputMsg defines the input message
type InputMsg struct {
	title    string
	subTitle string
	arg      string
}

// AlfredMsg is the msg object
type AlfredMsg struct {
	rawMsg []*InputMsg
	output string
}

// NewInputMsg init an alfred instance
func NewInputMsg(title string, subTitle string, arg string) *InputMsg {
	i := new(InputMsg)
	i.title = title
	i.subTitle = subTitle
	i.arg = arg
	return i
}

// NewAlfredMsg init the alfred msg obj
func NewAlfredMsg(iMsg []*InputMsg) *AlfredMsg {
	return &AlfredMsg{
		rawMsg: iMsg,
	}
}

// Format convert the raw std output into Alfred-formatted output
func (al *AlfredMsg) Format() {
	items := new(AlfredJSONFormat)
	for _, i := range al.rawMsg {
		item := new(AlfredJSONItem)
		item.Title = i.title
		item.SubTitle = i.subTitle
		item.Arg = i.arg
		// item.UID = ""
		text := new(Text)
		text.Copy = item.Title
		text.LargeType = item.Title
		item.Valid = true
		item.Text = text
		items.Items = append(items.Items, item)
	}

	res, _ := json.Marshal(items)
	al.output = string(res)
	return
}

//Print print the standard result
func (al *AlfredMsg) Print() {
	fmt.Print(al.output)
}

//FormatAndPrint combine Format and Print method
func (al *AlfredMsg) FormatAndPrint() {
	al.Format()
	al.Print()
}

//Error returns the std workflow error.
func Error(err error) {
	f := new(AlfredJSONFormat)
	item := new(AlfredJSONItem)
	item.Title = err.Error()
	item.UID = "error"
	item.Valid = false
	f.Items = append(f.Items, item)
	res, _ := json.Marshal(f)
	am := new(AlfredMsg)
	am.output = string(res)
	am.FormatAndPrint()
}
