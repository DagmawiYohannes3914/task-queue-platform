package queue

import (
	"log"

	"github.com/nats-io/nats.go"
)

var NatsConn *nats.Conn

func InitNATS(natsURL string) {
	var err error
	NatsConn, err = nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	log.Println("Connected to NATS")
}

func Publish(subject string, data []byte) error {
	return NatsConn.Publish(subject, data)
}
