package databox

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"runtime"
	"encoding/json"
	"fmt"
	"time"
)

const (
	apiUrl = "https://push2new.databox.com"
	clientVersion = "0.1.2"
)

type Client struct {
	PushToken string
	PushHost  string
}

type KPI struct {
	Key        string
	Value      float32
	Date       string
	Attributes map[string]interface{}
}

type KPIWrap struct {
	Data []map[string]interface{} `json:"data"`
}

func NewClient(pushToken string) *Client {
	return &Client{
		PushToken:    pushToken,
		PushHost:    apiUrl,
	}
}

type ResponseStatus    struct {
	Status string `json:"status"`
}

var postRequest = func(client *Client, path string, payload []byte) ([]byte, error) {
	userAgent := "Databox/" + clientVersion + " (" + runtime.Version() + ")"
	request, err := http.NewRequest("POST", (apiUrl + path), bytes.NewBuffer(payload))
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(client.PushToken, "")

	if err != nil {
		return nil, err
	}

	response, err_2 := (&http.Client{}).Do(request)
	if err_2 != nil {
		return nil, err_2
	}

	data, err_3 := ioutil.ReadAll(response.Body)
	return data, err_3
}

var getRequest = func(client *Client, path string) ([]byte, error) {
	userAgent := "Databox/" + clientVersion + " (" + runtime.Version() + ")"
	request, err := http.NewRequest("GET", (apiUrl + path), nil)
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(client.PushToken, "")

	if err != nil {
		return nil, err
	}

	response, err_2 := (&http.Client{}).Do(request)
	if err_2 != nil {
		return nil, err_2
	}

	data, err_3 := ioutil.ReadAll(response.Body)
	return data, err_3
}

type LastPush struct {
	Push           string `json: push`
	Time           time.Time `json: datetime`
	NumberOfErrors int `json: no_err`
	Errors         string `json: err`
	Keys           string `json: keys`
}

/*
TODO: Use this for better and fine grained response parsing
func (lastPush *LastPush) UnmarshalJSON(raw []byte) error {
	var preParsed map[string]interface{}
	if err := json.Unmarshal(raw, &preParsed); err != nil {
		return err
	}

	parsedTime, err_1 := time.Parse(time.RFC3339, preParsed["datetime"].(string))
	if err_1 != nil {
		return err_1
	}

	lastPush.Time = parsedTime
	return nil
}
*/

func (client *Client) LastPushes(n int) ([]LastPush, error) {
	response, err := getRequest(client, fmt.Sprintf("/lastpushes/%d", n))
	if err != nil {
		return nil, err
	}

	lastPushes := make([]LastPush, 0)
	err_1 := json.Unmarshal(response, &lastPushes);
	if err_1 != nil {
		return nil, err_1
	}

	return lastPushes, nil
}

func (client *Client) LastPush() (LastPush, error) {
	lastPushes, err := client.LastPushes(1)
	if err != nil {
		return LastPush{}, err
	}

	return lastPushes[0], nil
}

func (client *Client) Push(kpi *KPI) (*ResponseStatus, error) {
	payload, err := serializeKPIs([]KPI{*kpi})
	if err != nil {
		fmt.Println("serialize")
		return &ResponseStatus{}, err
	}

	response, err_2 := postRequest(client, "/", payload)
	if err_2 != nil {
		fmt.Println("post response:")
		return &ResponseStatus{}, err_2
	}

	var responseStatus = &ResponseStatus{}
	if err_3 := json.Unmarshal(response, &responseStatus); err_3 != nil {
		return &ResponseStatus{}, err_3
	}

	return responseStatus, nil
}

/* Serialisation */
func (kpi *KPI) ToJsonData() (map[string]interface{}) {
	var payload = make(map[string]interface{})
	payload["$" + kpi.Key] = kpi.Value

	if kpi.Date != "" {
		payload["date"] = kpi.Date
	}

	if len(kpi.Attributes) != 0 {
		for key, value := range kpi.Attributes {
			payload[key] = value
		}
	}

	return payload
}

func serializeKPIs(kpis []KPI) ([]byte, error) {
	wrap := KPIWrap{
		Data: make([]map[string]interface{}, 0),
	}

	for _, kpi := range (kpis) {
		wrap.Data = append(wrap.Data, kpi.ToJsonData())
	}

	return json.Marshal(wrap)
}
