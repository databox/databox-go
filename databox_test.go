package databox

import (
	"testing"
	"reflect"
)

var pushToken string = "adxg1kq5a4g04k0wk0s4wkssow8osw84"
var originalPostRequest = postRequest

func TestSimpleInit(t *testing.T) {
	token := "token"
	client := NewClient(token)

	if reflect.ValueOf(client).Kind().String() != "ptr" {
		t.Error("Not pointer")
	}

	if client.PushToken != token {
		t.Error("Toke is not set.")
	}
}

func TestLastPush(t *testing.T) {
	postRequest = func(client *Client, path string, payload []byte) ([]byte, error) {
		return []byte(`{"err":"[]"}`), nil
	}

	_, err := NewClient(pushToken).LastPush()
	if err != nil {
		t.Error("Error was raised", err)
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

	if status, _ := NewClient(pushToken).Push(&KPI{
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

	if status, _ := NewClient(pushToken).Push(&KPI{
		Key:"temp.ny",
		Value: 52.0,
		Date: "2015-01-01 09:00:00",
	}); status.Status == "ok" {
		t.Error("This should not be \"ok\"")
	}
}
