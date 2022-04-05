package controlplane

import (
	"context"
	"errors"
	"github.com/keptn/go-utils/pkg/api/models"
	"log"
)

var ErrEventHandleFatal = errors.New("fatal event handling error")
var ErrEventHandleIgnore = errors.New("event handling error")

// ControlPlane can be used to connect to the Keptn Control Plane
type ControlPlane struct {
	subscriptionSource   *SubscriptionSource
	eventSource          EventSource
	currentSubscriptions []models.EventSubscription
}

// New creates a new ControlPlane
// It is using a SubscriptionSource source to get information about current uniform subscriptions
// as well as an EventSource to actually receive events from Keptn
func New(subscriptionSource *SubscriptionSource, eventSource EventSource) *ControlPlane {
	return &ControlPlane{
		subscriptionSource:   subscriptionSource,
		eventSource:          eventSource,
		currentSubscriptions: []models.EventSubscription{},
	}
}

// Register is initially used to register the Keptn integration to the Control Plane
func (cp *ControlPlane) Register(ctx context.Context, integration Integration) error {
	eventUpdates := make(chan models.KeptnContextExtendedCE)
	subscriptionUpdates := make(chan []models.EventSubscription)
	if err := cp.eventSource.Start(ctx, integration.RegistrationData(), eventUpdates); err != nil {
		return err
	}
	if err := cp.subscriptionSource.Start(ctx, integration.RegistrationData(), subscriptionUpdates); err != nil {
		return err
	}
	for {
		select {
		case event := <-eventUpdates:
			err := cp.handle(ctx, event, integration)
			if errors.Is(err, ErrEventHandleFatal) {
				return err
			}
		case subscriptions := <-subscriptionUpdates:
			cp.currentSubscriptions = subscriptions
			cp.eventSource.OnSubscriptionUpdate(subjects(subscriptions))
		case <-ctx.Done():
			return nil
		}
	}
}

func (cp *ControlPlane) handle(ctx context.Context, event models.KeptnContextExtendedCE, integration Integration) error {
	subscriptionsForTopic := []models.EventSubscription{}
	for _, subscription := range cp.currentSubscriptions {
		if subscription.Event == *event.Type { // need to check against the name of the subscription because this can be a wildcard as well
			matcher := NewEventMatcherFromSubscription(subscription)
			if matcher.Matches(event) {
				subscriptionsForTopic = append(subscriptionsForTopic, subscription)
			}
		}
	}

	for range subscriptionsForTopic {
		if err := integration.OnEvent(context.WithValue(ctx, EventSenderKey, cp.eventSource.Sender()), event); err != nil {
			if errors.Is(err, ErrEventHandleFatal) {
				return err
			}
			if errors.Is(err, ErrEventHandleIgnore) {
				log.Print("error during handling of event")
			}
		}
	}
	return nil
}

func subjects(subscriptions []models.EventSubscription) []string {
	var ret []string
	for _, s := range subscriptions {
		ret = append(ret, s.Event)
	}
	return ret
}
