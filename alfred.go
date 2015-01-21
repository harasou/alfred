package alfred

import (
	"encoding/xml"
	"fmt"
	"log"
)

const xmlHeader = `<?xml version="1.0"?>`

type Items struct {
	XMLName struct{} `xml:"items"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	Uid          string
	Arg          string
	Valid        string
	Autocomplete string
	Type         string
	Title        string
	Subtitle     string
	SubtitleMod  []SubtitleMod
	Icon         string
	IconType     IconType
	Text         string
	TextType     []TextType
}

type SubtitleMod struct {
	Mod  string `xml:"mod,attr,omitempty"`
	Text string `xml:",chardata"`
}

type IconType struct {
	Type string `xml:"type,attr,omitempty"`
	Text string `xml:",chardata"`
}

type TextType struct {
	Type string `xml:"type,attr,omitempty"`
	Text string `xml:",chardata"`
}

func (i Item) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// 利便性のために string で受け取っていたメンバのつめ直し
	if i.Subtitle != "" {
		i.SubtitleMod = append([]SubtitleMod{{Text: i.Subtitle}}, i.SubtitleMod...)
	}
	if i.Icon != "" {
		i.IconType = IconType{Text: i.Icon}
	}
	if i.Text != "" {
		i.TextType = append([]TextType{{Text: i.Text}}, i.TextType...)
	}

	// 構造体を渡しても omitempty が効かないので、スライスにつめ直し
	var icon []IconType
	if i.IconType.Text != "" || i.IconType.Type != "" {
		icon = []IconType{i.IconType}
	}

	ii := struct {
		Uid          string        `xml:"uid,attr,omitempty"`
		Arg          string        `xml:"arg,attr,omitempty"`
		Valid        string        `xml:"valid,attr,omitempty"`
		Autocomplete string        `xml:"autocomplete,attr,omitempty"`
		Type         string        `xml:"type,attr,omitempty"`
		Title        string        `xml:"title"`
		SubtitleMod  []SubtitleMod `xml:"subtitle,omitempty"`
		IconType     []IconType    `xml:"icon,omitempty"`
		TextType     []TextType    `xml:"text,omitempty"`
	}{
		Uid:          i.Uid,
		Arg:          i.Arg,
		Valid:        i.Valid,
		Autocomplete: i.Autocomplete,
		Type:         i.Type,
		Title:        i.Title,
		SubtitleMod:  i.SubtitleMod,
		IconType:     icon,
		TextType:     i.TextType,
	}
	return e.EncodeElement(ii, start)
}

func Workflow() *Items {
	return new(Items).Init()
}

func (workflow *Items) Init() *Items {
	workflow.Items = []Item{}
	return workflow
}

func (workflow *Items) AddItem(item *Item) {
	workflow.Items = append(workflow.Items, *item)
}

func (workflow *Items) Print() {
	var xmlOutput, err = xml.MarshalIndent(workflow, "", "  ")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%v\n%v\n", xmlHeader, string(xmlOutput))
}
