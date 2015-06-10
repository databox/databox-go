package databox

import (
	"fmt"
)

const (
	apiUrl = "https://push2new.databox.com"
	clientVersion = "0.0.1"
)

func NewClient(pushToken string) *Client {
	return &Client{
		PushToken: 	pushToken,
		PushHost:	"https://push2new.databox.com",
	}
}

func (client *Client) Push(key string, number float32) {
	fmt.Println(key)
	fmt.Println(number)
}


type Client struct {
	PushToken string
	PushHost  string
}

func main() {
	c := NewClient("x")
	c.Push("oto", 23.0)
}