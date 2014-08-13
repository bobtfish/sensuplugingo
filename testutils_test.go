package testutils

import (
	"fmt"
	"testing"
)

func Test_exampleData(t *testing.T) {
	st, err := exampleData("check01")
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

