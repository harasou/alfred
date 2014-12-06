package alfred

import (
	"encoding/xml"
	"fmt"
)

const xmlHeader = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"

type Items struct {
	Items   []Item
	XMLName struct{} `xml:"items"`
}

type Item struct {
	Uid          string `xml:"uid,attr"`
	Valid        string `xml:"valid,attr"`
	Arg          string `xml:"arg,attr"`
	Autocomplete string `xml:"autocomplete,attr"`
	Title        string `xml:"title"`
	Subtitle     string `xml:"subtitle"`
	Icon         string `xml:"icon"`

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
	fmt.Println(xmlHeader, string(xmlOutput))
}
