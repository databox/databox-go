package databox

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"runtime"
	"encoding/json"
)

const (
	apiUrl = "https://push2new.databox.com"
	clientVersion = "0.1.1"
)

type Client struct {
	PushToken string
	PushHost  string
}

type KPI struct {
	Key   string
	Value float32
	Date  string
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

func (client *Client) LastPush() ([]byte, error) {
	return postRequest(client, "/lastpushes/1", []byte(``))
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
func (kpi *KPI) ToJsonData() (map[string]interface{}) {
	var payload = make(map[string]interface{})
	payload["$" + kpi.Key] = kpi.Value

	if kpi.Date != "" {
		payload["date"] = kpi.Date
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
