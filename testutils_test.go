package testutils

import (
	"fmt"
	"testing"
)

func Test_exampleData(t *testing.T) {
	b, err := exampleData("check01")
	st := string(b)
	if err != nil {
		t.Error("error reading check01")
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
	d, err := exampleData("check01")
	if err != nil {
                t.Error("error loading check01")
        }
	var c map[string]interface{}
	err = decodeCheckData(d, &c)
	if err != nil {
                t.Error(fmt.Sprintf("error decodingCheckData for check01: %v", err))
        }
	if c["action"] == "create" {
		t.Log("Can get fields")
	} else {
		t.Error("Cannot get fields")
	}
}

func Test_decodeDataType(t *testing.T) {
	d, err := exampleData("check01")
        if err != nil {
                t.Error("error loading check01")
        }
        var c CheckResult
	err = decodeCheckData(d, &c)
	if err != nil {
                t.Error(fmt.Sprintf("error decodingCheckDataType for check01: %v", err))
        }
	if c.Action == "create" {
		t.Log("Can get fields")
        } else {
                t.Error(fmt.Sprintf("Cannot get fields: %s", c.Action))
        }
}

