package notification

import "matchmaking"

type Client struct {
	publisher matchmaking.Publisher
	topic     string
}

func NewClient(publisher matchmaking.Publisher, topic string) *Client {
	return &Client{publisher: publisher, topic: topic}
}

func (c *Client) MatchCreated(match *matchmaking.Match) error {
	message := []byte("match created: " + match.ID)
	return c.publisher.Publish(c.topic, message)
}
