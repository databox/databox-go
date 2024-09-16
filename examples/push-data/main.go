package main

import (
	"context"
	"fmt"
	"os"
	"time"

	databox "github.com/databox/databox-go/databox"
)

const t = "<token>" // Your Databox token

func main() {

	// Create a context with basic auth
	auth := context.WithValue(context.Background(), databox.ContextBasicAuth, databox.BasicAuth{UserName: t})

	// Create a configuration
	cfg := databox.NewConfiguration()
	cfg.DefaultHeader["Content-Type"] = "application/json"
	cfg.DefaultHeader["Accept"] = "application/vnd.databox.v2+json"

	// Create an API client
	api := databox.NewAPIClient(cfg)

	// Create a new PushDataAttribute object - this is optional and represent the dimensions of the data
	a := databox.NewPushDataAttribute()
	a.SetKey("currency")
	a.SetValue("USD")

	var d []databox.PushDataAttribute
	d = append(d, *a)

	// Create a new PushData object and set the data
	data := databox.NewPushData()
	data.SetKey("sales")
	data.SetValue(100.0)
	data.SetDate(time.Now().UTC().Format(time.RFC3339))
	data.SetUnit("USD")
	data.SetAttributes(d)

	// Push the data
	r, err := api.DefaultAPI.DataPost(auth).PushData([]databox.PushData{*data}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DataPost``: %v\n", err)
	}

	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
}
