package pubsub

// Message encapsulates the data and the topic it was published to.
type Message struct {
	Topic string
	Payload interface{}
}