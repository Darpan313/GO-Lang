package pubsub

import "testing"

func TestPubSub_FanOut(t *testing.T) {
	broker := NewBroker()
	topic := "events"

	// Create 2 subscriber with ample buffer size
	sub1 := broker.Subscribe(topic, 5)
	sub2 := broker.Subscribe(topic, 5)

	broker.Publish(topic, "hello world!")

	msg1 := <-sub1.channel
	if msg1.Payload != "hello world!" {
		t.Errorf("Sub1 expected 'hello world', got %v", msg1.Payload)
	}

	msg2 := <-sub2.channel
	if msg2.Payload != "hello world!" {
		t.Errorf("Sub2 expected 'hello world', got %v", msg2.Payload)
	}
}

func TestPubSub_SlowSubscriber(t *testing.T) {
	
}

func TestPubSub_Unsubscribe(t *testing.T) {
	
}

func TestPubSub_Concurrency(t *testing.T) {
	
}