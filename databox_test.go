package databox

import (
	"context"
	"errors"
	"os"
	"reflect"
	"testing"
	"time"
)

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
	getRequest = func(ctx context.Context, client *Client, path string) ([]byte, error) {
		return []byte(`[
    {
        "request": {
            "date": "2018-03-12T10:16:57.108Z",
            "body": {
                "data": [
                    {
                        "$temp.ny": 52,
                        "date": "2015-01-01 09:00:00"
                    }
                ]
            },
            "errors": []
        },
        "response": {
            "date": "2018-03-12T10:16:57.109Z",
            "body": {
                "id": "15208128000647621f06d2625a6231"
            }
        },
        "metrics": [
            "90565|temp.ny"
        ]
    }
]`), nil
	}

	_, err := NewClient(getToken()).LastPush()
	if err != nil {
		t.Error("Error was raised", err)
	}
}

func TestKPI_ToJSONData(t *testing.T) {
	a := (&KPI{Key: "a", Value: float32(33)}).ToJSONData()
	if a["$a"] != float32(33) {
		t.Error("Conversion error")
	}

	date := "2015-01-01 09:00:00"
	b := (&KPI{Key: "a", Date: date}).ToJSONData()
	if b["date"] != date {
		t.Error("Conversion error")
	}
}

func TestSuccessfulPush(t *testing.T) {
	postRequest = func(ctx context.Context, client *Client, path string, payload []byte) ([]byte, error) {
		return []byte(`{"id":"someRandomId"}`), nil
	}

	if _, err := NewClient(getToken()).Push(&KPI{
		Key:   "temp.ny",
		Value: 60.0,
	}); err != nil {
		t.Error("Not inserted")
	}
}

func TestFailedPush(t *testing.T) {
	pushError := errors.New("invalid_json: some error message")
	postRequest = func(ctx context.Context, client *Client, path string, payload []byte) ([]byte, error) {
		return []byte(`{"type":"invalid_json","message":"some error message"}`), pushError
	}

	if _, err := NewClient(getToken()).Push(&KPI{
		Key:   "temp.ny",
		Value: 52.0,
		Date:  "2015-01-01 09:00:00",
	}); err == nil {
		t.Error("This should not be \"ok\"")
	}
}

func TestWithAdditionalAttributes(t *testing.T) {
	postRequest = func(ctx context.Context, client *Client, path string, payload []byte) ([]byte, error) {
		return []byte(`{"id":"someRandomId"}`), nil
	}

	client := NewClient(getToken())

	var attributes = make(map[string]interface{})
	attributes["test.number"] = 10
	attributes["test.string"] = "Oto Brglez"

	if _, err := client.Push(&KPI{
		Key:        "test.TestWithAdditionalAttributes",
		Value:      10.0,
		Date:       time.Now().Format(time.RFC3339),
		Attributes: attributes,
	}); err != nil {
		t.Error("Must be nil")
	}

	if _, err := client.LastPush(); err != nil {
		t.Error("Must be nil")
	}
}
