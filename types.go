package sensuplugingo 

type Handler interface {
    Handle(Event) error
}

type Filter interface {
    Filter(Event) (bool, error)
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

