package pubsub

// Subscriber represents a client listening to topic
type Subscriber struct {
	topic string
	channel chan *Message
}

// Messages returns a read-only channel to consume messages.
func (s *Subscriber) Messages() <-chan *Message {
	return s.channel
}

// Topic returns the topic this subscriber listening to
func (s *Subscriber) Topic() string {
	return s.topic
}

// close safely shuts down the Subscriber's channel.
func (s *Subscriber) close() {
	close(s.channel)
}