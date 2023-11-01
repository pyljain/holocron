package pubsub

type Publisher interface {
	Publish(collection string, embedding []float64, metadata map[string]string) error
}

type Subscriber interface {
	Pull(chan Event) error
}

type Event struct {
	Collection string            `json:"collection"`
	Embedding  []float64         `json:"embedding"`
	Metadata   map[string]string `json:"metadata"`
}
