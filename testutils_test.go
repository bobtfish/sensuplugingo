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
