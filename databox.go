package databox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	apiURL        = "https://push.databox.com"
	clientVersion = "2.0.0"
)

// Client struct holds push token and host to Databox service
type Client struct {
	PushToken string
	PushHost  string
}

// KPI struct holds information about item in push request
type KPI struct {
	Key        string
	Value      float32
	Date       string
	Attributes map[string]interface{}
}

// KPIWrap struct is just a wrapper around KPI with root key "data"
type KPIWrap struct {
	Data []map[string]interface{} `json:"data"`
}

// ResponseStatus struct is for different response variations
type ResponseStatus struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

// PushRequest struct holds information about Request returned from LastPush request
type PushRequest struct {
	Date   string   `json:"date"`
	Body   KPIWrap  `json:"body"`
	Errors []string `json:"errors"`
}

// PushResponse struct holds information about Response returned from LastPush request
type PushResponse struct {
	Date string         `json:"date"`
	Body ResponseStatus `json:"body"`
}

// LastPush struct holds item information from LastPush request
type LastPush struct {
	Request  PushRequest  `json:"request"`
	Response PushResponse `json:"response"`
	Metrics  []string     `json:"metrics"`
}

// NewClient returns object for making calls against a Databox service.
func NewClient(pushToken string) *Client {
	return &Client{
		PushToken: pushToken,
		PushHost:  apiURL,
	}
}

var postRequest = func(client *Client, path string, payload []byte) ([]byte, error) {
	userAgent := "databox-go/" + clientVersion
	accept := "application/vnd.databox.v" + strings.Split(clientVersion, ".")[0] + "+json"
	request, err := http.NewRequest("POST", (apiURL + path), bytes.NewBuffer(payload))
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", accept)
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(client.PushToken, "")

	if err != nil {
		return nil, err
	}

	response, err2 := (&http.Client{}).Do(request)
	if err2 != nil {
		return nil, err2
	}

	data, err3 := ioutil.ReadAll(response.Body)
	if err3 != nil {
		return data, err3
	}

	if response.StatusCode != 200 {
		var responseStatus = &ResponseStatus{}
		json.Unmarshal(data, &responseStatus)
		err4 := errors.New(responseStatus.Type + ": " + responseStatus.Message)
		return nil, err4
	}

	return data, nil
}

var getRequest = func(client *Client, path string) ([]byte, error) {
	userAgent := "databox-go/" + clientVersion
	accept := "application/vnd.databox.v" + strings.Split(clientVersion, ".")[0] + "+json"
	request, err := http.NewRequest("GET", (apiURL + path), nil)
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", accept)
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(client.PushToken, "")

	if err != nil {
		return nil, err
	}

	response, err2 := (&http.Client{}).Do(request)
	if err2 != nil {
		return nil, err2
	}

	data, err3 := ioutil.ReadAll(response.Body)
	return data, err3
}

// LastPushes returns n last pushes from Databox service
func (client *Client) LastPushes(n int) ([]LastPush, error) {
	response, err := getRequest(client, fmt.Sprintf("/lastpushes?limit=%d", n))
	if err != nil {
		return nil, err
	}

	lastPushes := make([]LastPush, 0)
	err1 := json.Unmarshal(response, &lastPushes)
	if err1 != nil {
		return nil, err1
	}

	return lastPushes, nil
}

// LastPush returns latest push from Databox service
func (client *Client) LastPush() (LastPush, error) {
	lastPushes, err := client.LastPushes(1)
	if err != nil {
		return LastPush{}, err
	}

	return lastPushes[0], nil
}

// Push makes push request against Databox service
func (client *Client) Push(kpi *KPI) (*ResponseStatus, error) {
	payload, err := serializeKPIs([]KPI{*kpi})
	if err != nil {
		return &ResponseStatus{}, err
	}

	response, err2 := postRequest(client, "/", payload)
	if err2 != nil {
		return &ResponseStatus{}, err2
	}

	var responseStatus = &ResponseStatus{}
	if err3 := json.Unmarshal(response, &responseStatus); err3 != nil {
		return &ResponseStatus{}, err3
	}

	return responseStatus, nil
}

// ToJSONData serializes KPI to json
func (kpi *KPI) ToJSONData() map[string]interface{} {
	var payload = make(map[string]interface{})
	payload["$"+kpi.Key] = kpi.Value

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

// serializeKPIs traverse all kpis and return json representation
func serializeKPIs(kpis []KPI) ([]byte, error) {
	wrap := KPIWrap{
		Data: make([]map[string]interface{}, 0),
	}

	for _, kpi := range kpis {
		wrap.Data = append(wrap.Data, kpi.ToJSONData())
	}

	return json.Marshal(wrap)
}
