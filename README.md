# Databox bindings for Go

[![Build Status](https://travis-ci.org/databox/databox-go.svg)](https://travis-ci.org/databox/databox-go)

Go wrapper for [Databox](http://databox.com) - Mobile Executive Dashboard.

## Installation

```bash
go install github.com/databox/databox-go
go get github.com/databox/databox-go # or this. :)
```

## Usage
```go
package main

import (
	databox "github.com/databox/databox-go"
	"fmt"
)

func main() {
    client := databox.NewClient("<push token>")

	if status, _ := client.Push(&databox.KPI{
		Key:	"temp.ny",
		Value: 	52.0,
		Date: 	"2015-01-01 09:00:00",
	}); status.Status == "ok" {
		fmt.Println("Inserted.")
	}

    if data, _ := client.LastPush() ; data != nil {
        fmt.Println(string(data))
    }
}

```

## Author
- [Oto Brglez](https://github.com/otobrglez)
