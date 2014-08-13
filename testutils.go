package testutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func exampleData(fn string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("test_data/%s.json", fn))
	return bytes, err
}

func decodeCheckData(b []byte, c interface{}) (err error) {
	return json.Unmarshal(b, c)
}

type Client struct {
	Name          string
	Address       string
	Subscriptions []string
	Timestamp     uint64
}

type CheckResult struct {
	Name        string
	Issued      uint64
	Output      string
	Status      uint32
	Command     string
	Subscribers []string
	Interval    uint64
	Handler     string
	History     []string
	Flapping    bool
}

type Event struct {
	Client      Client
	Check       CheckResult
	Action      string
	Occurrences uint64
}
