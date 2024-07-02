package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	TOPIC_RAILS_TO_GO = "rails-to-go"
	TOPIC_GO_TO_RAILS = "go-to-rails"
	PARTITION_DEFAULT = 0
)

type Messager struct {
	Conn *kafka.Conn
}

func New() *Messager {
	return &Messager{}
}

func (c *Messager) Setup(topic string) *Messager {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:29092", topic, PARTITION_DEFAULT)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	c.Conn = conn
	return c
}

func (c *Messager) Close() {
	if err := c.Conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func (c *Messager) Write(message []byte) (*Messager, error) {
	err := c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return c, err
	}

	_, err = c.Conn.Write(message)
	return c, err
}
