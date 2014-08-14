package filter

import (
	"testing"
        "github.com/bobtfish/sensuplugingo"
	"github.com/bobtfish/sensuplugingo/testutils"
)

func getExampleData(t *testing.T) sensuplugingo.Event {
	d, err := testutils.ExampleData("event01")
        if err != nil {
                t.Error("error loading event01")
        }
        var e sensuplugingo.Event
        err = testutils.DecodeCheckData(d, &e)
        if err != nil {
                t.Error("error decoding event01")
        }
	return e
}

func Test_FilterDisabled(t *testing.T) {
	e := getExampleData(t)
	f, err := FilterDisabled(e)
	if err != nil {
                t.Error("error filtering with FilterDisabled event01")
        }
	if f == false {
		t.Log("Did not filter")
	} else {
		t.Error("Did filter")
	}
}

func Test_FilterRepeated(t *testing.T) {
	e := getExampleData(t)
	f, err := FilterRepeated(e)
	if err != nil {
                t.Error("error filtering with FilterRepeated event01")
        }
	if f == false {
		t.Log("Did not filter")
	} else {
		t.Error("Did filter")
	}
}

func Test_FilterSilenced(t *testing.T) {
	e := getExampleData(t)
	f, err := FilterSilenced(e)
	if err != nil {
                t.Error("error filtering with FilterSilenced event01")
        }
	if f == false {
		t.Log("Did not filter")
	} else {
		t.Error("Did filter")
	}
}

func Test_FilterDependencies(t *testing.T) {
	e := getExampleData(t)
	f, err := FilterDisabled(e)
	if err != nil {
                t.Error("error filtering with FilterDependencies event01")
        }
	if f == false {
		t.Log("Did not filter")
	} else {
		t.Error("Did filter")
	}
}

