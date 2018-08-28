package types

// Message slack message
// https://api.slack.com/methods/chat.postMessage
type Message struct {
	Channel   string `json:"channel"`
	Text      string `json:"text"`
	LinkNames bool   `json:"link_names"`
}

// NewMessage New send message
func NewMessage(channel, text string) *Message {
	return &Message{
		Channel:   channel,
		Text:      text,
		LinkNames: true,
	}
}
