package alfred

import (
	"encoding/xml"
	"fmt"
	"log"
)

const xmlHeader = `<?xml version="1.0"?>`

type Items struct {
	XMLName struct{} `xml:"items"`

	Items []Item
}

type Item struct {
	XMLName struct{} `xml:"item"`

	Uid          string `xml:"uid,attr,omitempty"`
	Arg          string `xml:"arg,attr,omitempty"`
	Valid        string `xml:"valid,attr,omitempty"`
	Autocomplete string `xml:"autocomplete,attr,omitempty"`
	Type         string `xml:"type,attr,omitempty"`

	Title    string     `xml:"title"`
	Subtitle []Subtitle `xml:"subtitle,omitempty"`
	Icon     Icon       `xml:"icon,omitempty"`
	Text     []Text     `xml:"text,omitempty"`
}

type Subtitle struct {
	Mod  string `xml:"mod,attr,omitempty"`
	Text string `xml:",chardata"`
}

type Icon struct {
	Type string `xml:"type,attr,omitempty"`
	Text string `xml:",chardata"`
}

type Text struct {
	Type string `xml:"type,attr,omitempty"`
	Text string `xml:",chardata"`
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
