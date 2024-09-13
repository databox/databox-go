package main

import (
	"context"
	"fmt"
	"os"

	databox "github.com/databox/databox-go/databox"
)

const t = "<token>" // Your Databox token

func main() {

	// Create a context with basic auth
	auth := context.WithValue(context.Background(), databox.ContextBasicAuth, databox.BasicAuth{UserName: t})

	// Create a configuration
	configuration := databox.NewConfiguration()
	configuration.DefaultHeader["Content-Type"] = "application/json"
	configuration.DefaultHeader["Accept"] = "application/vnd.databox.v2+json"

	// Create an API client
	apiClient := databox.NewAPIClient(configuration)

	// Create a new PushData object
	data := databox.NewPushData()
	data.SetKey("test")
	data.SetValue(1.0)

	// Push the data
	r, err := apiClient.DefaultAPI.DataPost(auth).PushData([]databox.PushData{*data}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DataPost``: %v\n", err)
	}

	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
}
