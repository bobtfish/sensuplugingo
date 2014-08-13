package testutils

import (
	"fmt"
	"testing"
)

func Test_exampleData(t *testing.T) {
	b, err := exampleData("event01")
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

func Test_exampleDataFileNotFound(t *testing.T) {
        _, err := exampleData("doesnotexist")
        if err != nil {
                t.Log("error reading non existant file")
        } else {
		t.Error("Didn't report an error")
	}
}

func Test_decodeDataSimple(t *testing.T) {
	d, err := exampleData("event01")
	if err != nil {
                t.Error("error loading event01")
        }
	var e map[string]interface{}
	err = decodeCheckData(d, &e)
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
	d, err := exampleData("event01")
        if err != nil {
                t.Error("error loading event01")
        }
        var e Event 
	err = decodeCheckData(d, &e)
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
	// e.Client.Subscriptions all, frontend, proxy
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
	// e.Check.Subscribers frontend
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
	// e.Check.History "0", "2" 
	if e.Check.Flapping == false {
                t.Log("Can get fields")
        } else {
                t.Error(fmt.Sprintf("Cannot get fields: %s", e.Check.Flapping))
        }
}

