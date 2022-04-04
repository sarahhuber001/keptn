package nats

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/keptn/go-utils/pkg/api/models"
	"github.com/nats-io/nats.go"
	"os"
)

const (
	envVarNatsURL        = "NATS_URL"
	envVarNatsURLDefault = "nats://keptn-nats"
)

var (
	ErrSubAlreadySubscribed   = errors.New("already subscribed")
	ErrSubNilMessageProcessor = errors.New("message processor is nil")
	ErrSubEmptySubject        = errors.New("empty subject")
	ErrPubEventTypeMissing    = errors.New("event is missing the event type")
)

// ProcessEventFn is used to process a received keptn event
type ProcessEventFn func(event models.KeptnContextExtendedCE) error

// NatsConnector can be used to subscribe to certain events
// on the NATS event system
type NatsConnector struct {
	conn          *nats.Conn
	subscriptions map[string]*nats.Subscription
}

// Connect connects a NatsConnector to NATS.
// Note that this will automatically and indefinitely try to reconnect
// as soon as it looses connection
func Connect(connectURL string) (*NatsConnector, error) {
	conn, err := nats.Connect(connectURL, nats.MaxReconnects(-1))
	if err != nil {
		return nil, fmt.Errorf("could not connect to NATS: %w", err)
	}
	return &NatsConnector{
		conn:          conn,
		subscriptions: make(map[string]*nats.Subscription),
	}, nil
}

// ConnectFromEnv connects a NatsConnector to NATS.
// The URL is read from the environment variable "NATS_URL"
// If the URL is not set via the environment variable "NATS_URL",
// it falls back to the default URL "nats://keptn-nats"
func ConnectFromEnv() (*NatsConnector, error) {
	natsURL := os.Getenv(envVarNatsURL)
	if natsURL == "" {
		natsURL = envVarNatsURLDefault
	}
	return Connect(natsURL)
}

func (nc *NatsConnector) UnsubscribeAll() error {
	for _, s := range nc.subscriptions {
		if err := s.Unsubscribe(); err != nil {
			return fmt.Errorf("unable to unsubscribe from subject %s: %w", s.Subject, err)
		}
	}
	nc.subscriptions = make(map[string]*nats.Subscription)
	return nil
}

// Subscribe adds a subscription to a specific subject to the NatsConnector.
// It takes the subject as string (usually the event type) and a function fn
// being called when an event is received
func (nc *NatsConnector) Subscribe(subject string, fn ProcessEventFn) error {
	return nc.QueueSubscribe("", subject, fn)
}

func (nc *NatsConnector) QueueSubscribe(queueGroup string, subject string, fn ProcessEventFn) error {
	if subject == "" {
		return ErrSubEmptySubject
	}
	if fn == nil {
		return ErrSubNilMessageProcessor
	}
	return nc.queueSubscribe(subject, queueGroup, fn)
}

func (nc *NatsConnector) SubscribeMultiple(subjects []string, fn ProcessEventFn) error {
	return nc.QueueSubscribeMultiple("", subjects, fn)
}

func (nc *NatsConnector) QueueSubscribeMultiple(queueGroup string, subjects []string, fn ProcessEventFn) error {
	if fn == nil {
		return ErrSubNilMessageProcessor
	}

	for _, sub := range subjects {
		if err := nc.queueSubscribe(sub, queueGroup, fn); err != nil {
			return fmt.Errorf("could not subscribe to subject %s: %w", sub, err)
		}
	}
	return nil
}

func (nc *NatsConnector) Publish(event models.KeptnContextExtendedCE) error {
	if event.Type == nil || *event.Type == "" {
		return ErrPubEventTypeMissing
	}
	serializedEvent, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("could not publish event: %w", err)
	}
	return nc.conn.Publish(*event.Type, serializedEvent)
}

func (nc *NatsConnector) queueSubscribe(subject string, queueGroup string, fn ProcessEventFn) error {
	sub, err := nc.conn.QueueSubscribe(subject, queueGroup, func(m *nats.Msg) {
		event := &models.KeptnContextExtendedCE{}
		if err := json.Unmarshal(m.Data, event); err != nil {
			fmt.Printf("could not unmarshal message %s: %v\n", string(m.Data), err)
			return
		}
		err := fn(*event)
		if err != nil {
			fmt.Printf("Could not process message %s: %v\n", string(m.Data), err)
		}
	})
	if err != nil {
		return fmt.Errorf("could not subscribe to subject %s: %w", subject, err)
	}
	if nc.subscriptions == nil {
		nc.subscriptions = make(map[string]*nats.Subscription)
	}

	if _, ok := nc.subscriptions[subject]; ok {
		return ErrSubAlreadySubscribed
	}

	nc.subscriptions[subject] = sub
	return nil
}
