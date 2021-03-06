# harasou/alfred

This is the go library for creating Alfred 2 workflows

## Setup

```
go get "github.com/harasou/alfred"
```

## Example Usage

```go
package main

import (
	"github.com/harasou/alfred"
)

func main() {

	wf := alfred.Workflow()
	defer wf.Print()

	wf.AddItem(&alfred.Item{
		Uid:          "desktop",
		Arg:          "~/Desktop",
		Valid:        "YES",
		Autocomplete: "Desktop",
		Type:         "file",
		Title:        "Desktop",
		Subtitle:     "~/Desktop",
		IconType:     alfred.IconType{Text: "~/Desktop", Type: "fileicon"},
	})

	wf.AddItem(&alfred.Item{
		Uid:          "flickr",
		Valid:        "no",
		Autocomplete: "flickr",
		Title:        "Flickr",
		Icon:         "flickr.png",
	})

	wf.AddItem(&alfred.Item{
		Uid:          "image",
		Autocomplete: "My holiday photo",
		Type:         "file",
		Title:        "My holiday photo",
		Subtitle:     "~/Pictures/My holiday photo.jpg",
		IconType:     alfred.IconType{Text: "public.jpeg", Type: "filetype"},
	})

	wf.AddItem(&alfred.Item{
		Uid:          "home",
		Arg:          "~/",
		Valid:        "YES",
		Autocomplete: "Home",
		Type:         "file",
		Title:        "Home Folder",
		IconType:     alfred.IconType{Text: "~/", Type: "fileicon"},
		Subtitle:     "Home folder ~/",
		SubtitleMod: []alfred.SubtitleMod{
			{Text: "Subtext when shift is pressed", Mod: "shift"},
			{Text: "Subtext when fn is pressed", Mod: "fn"},
			{Text: "Subtext when ctrl is pressed", Mod: "ctrl"},
			{Text: "Subtext when alt is pressed", Mod: "alt"},
			{Text: "Subtext when cmd is pressed", Mod: "cmd"},
		},
		TextType: []alfred.TextType{
			{Text: "Text when copying", Type: "copy"},
			{Text: "Text for LargeType", Type: "largetype"},
		},
	})

}
```

Output

```xml
<?xml version="1.0"?>
<items>
  <item uid="desktop" arg="~/Desktop" valid="YES" autocomplete="Desktop" type="file">
    <title>Desktop</title>
    <subtitle>~/Desktop</subtitle>
    <icon type="fileicon">~/Desktop</icon>
  </item>
  <item uid="flickr" valid="no" autocomplete="flickr">
    <title>Flickr</title>
    <icon>flickr.png</icon>
  </item>
  <item uid="image" autocomplete="My holiday photo" type="file">
    <title>My holiday photo</title>
    <subtitle>~/Pictures/My holiday photo.jpg</subtitle>
    <icon type="filetype">public.jpeg</icon>
  </item>
  <item uid="home" arg="~/" valid="YES" autocomplete="Home" type="file">
    <title>Home Folder</title>
    <subtitle mod="shift">Subtext when shift is pressed</subtitle>
    <subtitle mod="fn">Subtext when fn is pressed</subtitle>
    <subtitle mod="ctrl">Subtext when ctrl is pressed</subtitle>
    <subtitle mod="alt">Subtext when alt is pressed</subtitle>
    <subtitle mod="cmd">Subtext when cmd is pressed</subtitle>
    <subtitle>Home folder ~/</subtitle>
    <icon type="fileicon">~/</icon>
    <text type="copy">Text when copying</text>
    <text type="largetype">Text for LargeType</text>
  </item>
</items>
```
