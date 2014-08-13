package testutils

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func exampleData(fn string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("test_data/%s.json", fn))
	return bytes, err
}

func decodeCheckData(b []byte, c interface{}) (err error) {
    return json.Unmarshal(b, c)
}

type Event struct {
	Action string
	Occurrences uint64
	
}

