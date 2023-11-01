package pubsub

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type NATS struct {
	conn *nats.Conn
}

func NewNATS(url string) (*NATS, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NATS{
		conn: nc,
	}, nil
}

func (nats *NATS) Drain() {
	nats.conn.Drain()
}

func (n *NATS) Publish(collection string, embedding []float64, metadata map[string]string) error {
	event := Event{
		Collection: collection,
		Embedding:  embedding,
		Metadata:   metadata,
	}

	eventBytes, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	err = n.conn.Publish("holocron.ingest", eventBytes)
	if err != nil {
		return err
	}

	return nil
}

func (n *NATS) Pull(genericEventsChan chan Event) error {
	defer close(genericEventsChan)

	natsMessageChannel := make(chan *nats.Msg)
	subscription, err := n.conn.ChanSubscribe("holocron.ingest", natsMessageChannel)
	if err != nil {
		return err
	}

	defer subscription.Unsubscribe()
	defer close(natsMessageChannel)

	for e := range natsMessageChannel {
		gevent := Event{}
		err := json.Unmarshal(e.Data, &gevent)
		if err != nil {
			log.Printf("Could not unmarshal event: %+v", e.Data)
		}

		genericEventsChan <- gevent
	}

	return nil
}
