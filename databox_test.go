package databox

import (
	"testing"
	"reflect"
	"os"
)

var originalPostRequest = postRequest
var originalGetRequest = getRequest

func getToken() (pushToken string) {
	pushToken = "adxg1kq5a4g04k0wk0s4wkssow8osw84"
	envPushToken := "" + os.Getenv("DATABOX_PUSH_TOKEN")
	if envPushToken != "" {
		pushToken = os.Getenv("DATABOX_PUSH_TOKEN")
	}
	return
}

func TestSimpleInit(t *testing.T) {
	token := getToken()
	client := NewClient(getToken())

	if reflect.ValueOf(client).Kind().String() != "ptr" {
		t.Error("Not pointer")
	}

	if client.PushToken != token {
		t.Error("Token is not set.")
	}
}

func TestLastPush(t *testing.T) {
	getRequest = func(client *Client, path string) ([]byte, error) {
		return []byte(`[
  {
    "push": "{\"data\":[{\"$sales\":203},{\"$sales\":103,\"date\":\"2015-01-01 17:00:00\"}]}",
    "err": "[]",
    "no_err": 0,
    "datetime": "2016-01-25T22:08:20.704Z",
    "keys": "[\"2850|sales\"]"
  }
]`), nil
	}


	lastPush, err := NewClient(getToken()).LastPush()
	if err != nil {
		t.Error("Error was raised", err)
	}

	if lastPush.NumberOfErrors != 0 {
		t.Error("Number of errors in last push must equal 0!")
	}

	if lastPush.Push == "" {
		t.Error("Push must not be nil")
	}
}

func TestKPI_ToJsonData(t *testing.T) {
	a := (&KPI{Key:"a", Value:float32(33)}).ToJsonData()
	if a["$a"] != float32(33) {
		t.Error("Conversion error")
	}

	date := "2015-01-01 09:00:00"
	b := (&KPI{Key:"a", Date:date}).ToJsonData()
	if b["date"] != date {
		t.Error("Conversion error")
	}
}

func TestSuccessfulPush(t *testing.T) {
	postRequest = func(client *Client, path string, payload []byte) ([]byte, error) {
		return []byte(`{"status":"ok"}`), nil
	}

	if status, _ := NewClient(getToken()).Push(&KPI{
		Key:"temp.ny",
		Value: 60.0,
	}); status.Status != "ok" {
		t.Error("Not inserted")
	}
}

func TestFailedPush(t *testing.T) {
	postRequest = func(client *Client, path string, payload []byte) ([]byte, error) {
		return []byte(`{"status":"error"}`), nil
	}

	if status, _ := NewClient(getToken()).Push(&KPI{
		Key:"temp.ny",
		Value: 52.0,
		Date: "2015-01-01 09:00:00",
	}); status.Status == "ok" {
		t.Error("This should not be \"ok\"")
	}
}
