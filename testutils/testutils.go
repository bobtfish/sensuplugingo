package testutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ExampleData(fn string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("test_data/%s.json", fn))
	return bytes, err
}

func DecodeCheckData(b []byte, c interface{}) (err error) {
	return json.Unmarshal(b, c)
}
