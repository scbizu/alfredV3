package alfredV3

//AlfredJSONFormat defines the standard `Script Filter JSON Format`
//more details are mentioned below:
//https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
type AlfredJSONFormat struct {
	Items []*AlfredJSONItem `json:"items"`
}

//AlfredJSONItem defines a single item
type AlfredJSONItem struct {
	UID       string `json:"uid"`
	Title     string `json:"title"`
	SubTitle  string `json:"subtitle"`
	Arg       string `json:"arg"`
	Icon      *Icon  `json:"icon"`
	Valid     bool   `json:"valid"`
	AC        string `json:"autocomplete"`
	Type      string `json:"type"`
	Mods      *Mods  `json:"mods"`
	Text      *Text  `json:"text"`
	QuickLook string `json:"quicklookurl"`
}

//Icon defines icon object
type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

//Mods defines mod object
type Mods struct {
	Alt  ModOptions `json:"alt"`
	Cmd  ModOptions `json:"cmd"`
	Ctrl ModOptions `json:"ctrl"`
}

//ModOptions defines the extra option in mod
type ModOptions struct {
	Valid    bool   `json:"valid"`
	Arg      string `json:"arg"`
	Subtitle string `json:"subtitle"`
	Icon     Icon   `json:"icon"`
}

//Text defines the text option
type Text struct {
	Copy      string `json:"copy"`
	LargeType string `json:"largetype"`
}
