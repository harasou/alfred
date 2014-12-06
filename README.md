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
    Title:        "Desktop",
    Subtitle:     "~/Desktop",
    Icon:         "~/Desktop",
  })

  item := &alfred.Item{}
  item.Uid = "flickr"
  item.Valid = "no"
  item.Autocomplete = "flickr"
  item.Title = "Flickr"
  item.Icon = "flickr.png"

  wf.AddItem(item)

}
```

Output

```xml
<?xml version="1.0" encoding="UTF-8"?>
<items>
  <item uid="desktop" valid="YES" arg="~/Desktop" autocomplete="Desktop">
    <title>Desktop</title>
    <subtitle>~/Desktop</subtitle>
    <icon>~/Desktop</icon>
  </item>
  <item uid="flickr" valid="no" arg="" autocomplete="flickr">
    <title>Flickr</title>
    <subtitle></subtitle>
    <icon>flickr.png</icon>
  </item>
</items>
```
