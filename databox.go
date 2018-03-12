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
	apiUrl        = "https://push.databox.com"
	clientVersion = "2.0.0"
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
		PushToken: pushToken,
		PushHost:  apiUrl,
	}
}

type ResponseStatus struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

var postRequest = func(client *Client, path string, payload []byte) ([]byte, error) {
	userAgent := "databox-go/" + clientVersion
	accept := "application/vnd.databox.v" + strings.Split(clientVersion, ".")[0] + "+json"
	request, err := http.NewRequest("POST", (apiUrl + path), bytes.NewBuffer(payload))
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", accept)
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
	if err_3 != nil {
		return data, err_3
	}

	if response.StatusCode != 200 {
		var responseStatus = &ResponseStatus{}
		json.Unmarshal(data, &responseStatus)
		err_4 := errors.New(responseStatus.Type + ": " + responseStatus.Message)
		return nil, err_4
	}

	return data, nil
}

var getRequest = func(client *Client, path string) ([]byte, error) {
	userAgent := "databox-go/" + clientVersion
	accept := "application/vnd.databox.v" + strings.Split(clientVersion, ".")[0] + "+json"
	request, err := http.NewRequest("GET", (apiUrl + path), nil)
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", accept)
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

type PushRequest struct {
	Date   string   `json:"date"`
	Body   KPIWrap  `json:"body"`
	Errors []string `json:"errors"`
}

type PushResponse struct {
	Date string         `json:"date"`
	Body ResponseStatus `json:"body"`
}

type LastPush struct {
	Request  PushRequest  `json:"request"`
	Response PushResponse `json:"response"`
	Metrics  []string     `json:"metrics"`
}

func (client *Client) LastPushes(n int) ([]LastPush, error) {
	response, err := getRequest(client, fmt.Sprintf("/lastpushes?limit=%d", n))
	if err != nil {
		return nil, err
	}

	lastPushes := make([]LastPush, 0)
	err_1 := json.Unmarshal(response, &lastPushes)
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
		return &ResponseStatus{}, err
	}

	response, err_2 := postRequest(client, "/", payload)
	if err_2 != nil {
		return &ResponseStatus{}, err_2
	}

	var responseStatus = &ResponseStatus{}
	if err_3 := json.Unmarshal(response, &responseStatus); err_3 != nil {
		return &ResponseStatus{}, err_3
	}

	return responseStatus, nil
}

/* Serialisation */
func (kpi *KPI) ToJsonData() map[string]interface{} {
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

func serializeKPIs(kpis []KPI) ([]byte, error) {
	wrap := KPIWrap{
		Data: make([]map[string]interface{}, 0),
	}

	for _, kpi := range kpis {
		wrap.Data = append(wrap.Data, kpi.ToJsonData())
	}

	return json.Marshal(wrap)
}
