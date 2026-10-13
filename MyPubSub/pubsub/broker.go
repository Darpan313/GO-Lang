package pubsub

import "sync"

// Broker manages topics, subscribers, and message routing
type Broker struct {
	topics 	map[string]map[*Subscriber]struct{}
	mu 		sync.RWMutex
}

// NewBroker initializes new pub/sub broker
func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string]map[*Subscriber]struct{}),
	}
}

func (b *Broker) Subscribe(topic string, bufferSize int) *Subscriber{
	sub := &Subscriber{
		topic: topic,
		channel: make(chan *Message, bufferSize),
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.topics[topic] == nil {
		b.topics[topic] = make(map[*Subscriber]struct{})
	}

	b.topics[topic][sub] = struct{}{}

	return sub
}

func (b *Broker) UnSubscribe(sub *Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()

	topic := sub.topic

	if topicSubs, exists := b.topics[topic]; exists {
		if _, ok := topicSubs[sub]; ok {
			delete(topicSubs, sub)
			sub.close()

			if len(topicSubs) == 0 {
				delete(b.topics, topic)
			}
		}
	}
}

func (b *Broker) Publish(topic string, payload interface{}) {
	b.mu.RLock()
	subsMap, exists := b.topics[topic]

	if !exists {
		b.mu.RUnlock()
		return
	}

	// Copy the subscriber pointers to avoid holding the lock during channel sends.
	// This prevents a slow fan-out from blocking new subscriptions.
	subs := make([]*Subscriber, 0, len(subsMap))
	for sub := range subsMap {
		subs = append(subs, sub)
	}
	b.mu.RUnlock()

	msg := &Message{Topic: topic, Payload: payload}

	for _, sub := range subs {
		select {
		case sub.channel<-msg:
		default:
		}
	}
}