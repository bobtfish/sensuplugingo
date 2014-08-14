package testutils

import (
	"fmt"
	"github.com/bobtfish/sensuplugingo"
	"testing"
)

func Test_ExampleData(t *testing.T) {
	b, err := ExampleData("event01")
	st := string(b)
	if err != nil {
		t.Error("error reading event01")
	}
	expLen := 644
	if len(st) == expLen {
		t.Log("got example data")
	} else {
		t.Error(fmt.Sprintf("bad example data: expected %d, got %d - %s", expLen, len(st), st))
	}
}

func Test_ExampleDataFileNotFound(t *testing.T) {
	_, err := ExampleData("doesnotexist")
	if err != nil {
		t.Log("error reading non existant file")
	} else {
		t.Error("Didn't report an error")
	}
}

func Test_decodeDataSimple(t *testing.T) {
	d, err := ExampleData("event01")
	if err != nil {
		t.Error("error loading event01")
	}
	var e map[string]interface{}
	err = DecodeCheckData(d, &e)
	if err != nil {
		t.Error(fmt.Sprintf("error decodingCheckData for event01: %v", err))
	}
	if e["action"] == "create" {
		t.Log("Can get fields")
	} else {
		t.Error("Cannot get fields")
	}
}

func Test_decodeDataType(t *testing.T) {
	d, err := ExampleData("event01")
	if err != nil {
		t.Error("error loading event01")
	}
	var e sensuplugingo.Event
	err = DecodeCheckData(d, &e)
	if err != nil {
		t.Error(fmt.Sprintf("error decodingCheckDataType for event01: %v", err))
	}
	if e.Action == "create" {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %s", e.Action))
	}
	if e.Occurrences == 1 {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %d", e.Occurrences))
	}
	if e.Client.Name == "host01" {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %s", e.Client.Name))
	}
	if e.Client.Address == "10.2.1.11" {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %s", e.Client.Address))
	}
	if len(e.Client.Subscriptions) == 3 && e.Client.Subscriptions[0] == "all" && e.Client.Subscriptions[1] == "frontend" && e.Client.Subscriptions[2] == "proxy" {
		t.Log("Can get fields")
	} else {
		t.Error("Cannot get fields for Subscriptions")
	}
	if e.Client.Timestamp == 1326390159 {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %d", e.Client.Timestamp))
	}
	if e.Check.Name == "frontend_http_check" {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %s", e.Check.Name))
	}
	if e.Check.Issued == 1326390169 {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %d", e.Check.Issued))
	}
	if e.Check.Output == "HTTP CRITICAL: HTTP/1.1 503 Service Temporarily Unavailable" {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %s", e.Check.Output))
	}
	if e.Check.Status == 2 {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %d", e.Check.Status))
	}
	if e.Check.Command == "check_http -I 127.0.0.1 -u http://web.example.com/healthcheck.html -R 'pageok'" {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %s", e.Check.Command))
	}
	if len(e.Check.Subscribers) == 1 && e.Check.Subscribers[0] == "frontend" {
		t.Log("Can get fields")
	} else {
		t.Error("Cannot get fields for Subscriptions")
	}
	if e.Check.Interval == 60 {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %s", e.Check.Interval))
	}
	if e.Check.Handler == "campfire" {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %s", e.Check.Handler))
	}
	if len(e.Check.History) == 2 && e.Check.History[0] == "0" && e.Check.History[1] == "2" {
		t.Log("Can get fields")
	} else {
		t.Error("Cannot get fields for Subscriptions")
	}
	if e.Check.Flapping == false {
		t.Log("Can get fields")
	} else {
		t.Error(fmt.Sprintf("Cannot get fields: %s", e.Check.Flapping))
	}
}
