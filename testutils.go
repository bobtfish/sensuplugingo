package testutils

import (
	"fmt"
	"io/ioutil"
)

func exampleData(fn string) (string, error) {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("test_data/%s.json", fn))
	return string(bytes), err
}
