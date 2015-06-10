package databox

import (
	"testing"
	"reflect"
)

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

func TestSimplePush(t *testing.T) {
	client := NewClient("adxg1kq5a4g04k0wk0s4wkssow8osw84")
	client.Push("temp.boston", 52.3)
}

