package bot

// ChannelListURL slack Channel URL
var ChannelListURL = "https://slack.com/api/channels.list"

// Channels slack api will return this Channel
type Channels struct {
	OK       bool      `json:"ok"`
	Channels []Channel `json:"channels"`
}

// Channel channel's info
type Channel struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Purpose Purpose `json:"purpose"`
}

// Purpose channel's purpose
type Purpose struct {
	Value string `json:"value"`
}

// NewChannels create new channel object
func NewChannels() *Channels {
	return new(Channels)
}
