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
		Icon:         "~/Desktop",
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
		Icon:         "public.jpeg",
	})

	wf.AddItem(&alfred.Item{
		Uid:          "home",
		Arg:          "~/",
		Valid:        "YES",
		Autocomplete: "Home",
		Type:         "file",
		Title:        "Home Folder",
		Subtitle:     "Home folder ~/",
		Icon:         "~/",
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
    <icon>~/Desktop</icon>
  </item>
  <item uid="flickr" valid="no" autocomplete="flickr">
    <title>Flickr</title>
    <icon>flickr.png</icon>
  </item>
  <item uid="image" autocomplete="My holiday photo" type="file">
    <title>My holiday photo</title>
    <subtitle>~/Pictures/My holiday photo.jpg</subtitle>
    <icon>public.jpeg</icon>
  </item>
  <item uid="home" arg="~/" valid="YES" autocomplete="Home" type="file">
    <title>Home Folder</title>
    <subtitle>Home folder ~/</subtitle>
    <icon>~/</icon>
  </item>
</items>
```
