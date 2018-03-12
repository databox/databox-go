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
	"time"
)

func main() {
    client := databox.NewClient("<push token>")

	if _, err := client.Push(&databox.KPI{
		Key:	"temp.ny",
		Value: 	52.0,
		Date: 	"2015-01-01 09:00:00",
	}); err != nil {
		fmt.Println("Inserted.")
	}

  if data, _ := client.LastPushes(1) ; data != nil {
		for _, push := range data {
			printPush(push)
		}
  }

	// Additional attributes
  var attributes = make(map[string]interface{})
  attributes["test.number"] = 10
  attributes["test.string"] = "Oto Brglez"

	if _, err := client.Push(&databox.KPI{
		Key: "testing.this",
		Value: 10.0,
		Date: time.Now().Format(time.RFC3339),
		Attributes: attributes,
	}); err != nil {
		fmt.Println("Inserted.")
	}

  // Retriving last push
	if data, _ := client.LastPushes(1) ; data != nil {
		for _, push := range data {
			printPush(push)
		}
  }
}

func printPush(push databox.LastPush) {
	fmt.Println("Request")
	fmt.Println(" date: " + push.Request.Date)
	fmt.Print(" body: ")
	fmt.Println(push.Request.Body)
	fmt.Print(" errors: ")
	fmt.Println(push.Request.Errors)
	fmt.Println("")

	fmt.Println("Response")
	fmt.Println(" date: " + push.Response.Date)
	fmt.Print(" body: ")
	fmt.Println(push.Response.Body)
	fmt.Println("")

	fmt.Println("Metrics")
	for _, value := range push.Metrics {
		fmt.Println(" " + value)
	}
	fmt.Println("")
}

```

## Development


```bash
# develop
docker run -it --rm -v "$PWD":/usr/local/go/src/databox golang:1.5 bash

# test
docker run -it --rm -v "$PWD":/usr/local/go/src/databox golang:1.5 make
```

## Author
- [Oto Brglez](https://github.com/otobrglez)
- [Vlada Petrovic](https://github.com/vladapetrovic)

[BuildStatus-Url]: https://travis-ci.org/databox/databox-go
[BuildStatus-Image]: https://travis-ci.org/databox/databox-go.svg                     
[ReportCard-Url]: http://goreportcard.com/report/databox/databox-go
[ReportCard-Image]: http://goreportcard.com/badge/databox/databox-go
