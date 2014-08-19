package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bobtfish/sensuplugingo"
	"io/ioutil"
	"os"
)

func GetEventFromStdin() (e sensuplugingo.Event, err error) {
	return GetEventFrom(os.Stdin)
}

func GetEventFrom(in *os.File) (e sensuplugingo.Event, err error) {
	b, err := ioutil.ReadAll(in)
	if err != nil {
		panic("Could not readall stdin")
	}
	err = json.Unmarshal(b, &e)
	if err != nil {
		panic(fmt.Sprintf("Could not decode json: %s", err))
	}
	return e, err
}
