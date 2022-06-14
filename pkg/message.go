package pkg

// Message was never implemented
type Message struct {
	timeScheduled int64
	id            int64
	data          interface{}
	state         int64
	tryCount      uint64
	failAfter     int64
}
