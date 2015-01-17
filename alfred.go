package alfred

import (
	"encoding/xml"
	"fmt"
)

const xmlHeader = `<?xml version="1.0"?>`

type Items struct {
	Items   []Item
	XMLName struct{} `xml:"items"`
}

type Item struct {
	Uid          string `xml:"uid,attr,omitempty"`
	Arg          string `xml:"arg,attr,omitempty"`
	Valid        string `xml:"valid,attr,omitempty"`
	Autocomplete string `xml:"autocomplete,attr,omitempty"`
	Type         string `xml:"type,attr,omitempty"`
	Title        string `xml:"title"`
	Subtitle     string `xml:"subtitle,omitempty"`
	Icon         string `xml:"icon,omitempty"`

	XMLName struct{} `xml:"item"`
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
	var xmlOutput, _ = xml.MarshalIndent(workflow, "", "  ")
	fmt.Printf("%v\n%v", xmlHeader, string(xmlOutput))
}
