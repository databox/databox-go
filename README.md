# Databox bindings for Go

Go wrapper for [Databox](http://databox.com) - Mobile Executive Dashboard.

[![Build Status][BuildStatus-Image]][BuildStatus-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]


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
    
    // Additional attributes
    var attributes = make(map[string]interface{})
    attributes["test.number"] = 10
    attributes["test.string"] = "Oto Brglez"
    
	if status, _ := client.Push(&KPI{
		Key: "testing.this",
		Value: 10.0,
		Date: time.Now().Format(time.RFC3339),
		Attributes: attributes,
	}); status.Status == "ok" {
		t.Error("This status must be ok")
	}
    
    // Retriving last push
    lastPush, err := client.LastPush()
    if err != nil {
        t.Error("Error was raised", err)
    }
 
    fmt.Println("Number of errors in last push", lastPush.NumberOfErrors) 
}

```

## Author
- [Oto Brglez](https://github.com/otobrglez)


[BuildStatus-Url]: https://travis-ci.org/databox/databox-go
[BuildStatus-Image]: https://travis-ci.org/databox/databox-go.svg                     
[ReportCard-Url]: http://goreportcard.com/databox/databox-go/s
[ReportCard-Image]: http://goreportcard.com/databox/databox-go/s