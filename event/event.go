package event

// ValidRequestType valid Slack Request type
// https://api.slack.com/events-api
// Callback field overview
var ValidRequestType = map[string]struct{}{
	"url_verification": struct{}{},
	"event_callback":   struct{}{},
}

// Event Slack Event type
type Event interface {
	Name() string
}

// RequestType Slac Request Type [url_verification,event_callback...]
type RequestType struct {
	Type string `json:"type"`
}

// Request Slack event request
type Request struct {
	Token       string   `json:"token"`
	TeamID      string   `json:"team_id"`
	APIAppID    string   `json:"api_app_id"`
	Event       Event    `json:"event"`
	Type        string   `json:"type"`
	EventID     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
	AuthedUsers []string `json:"authed_users"`
}

// IsValidRequestType check request type
func IsValidRequestType(requestType string) bool {
	if _, ok := ValidRequestType[requestType]; !ok {
		return false
	}
	return true
}

// NewRequest New Slack Request
func NewRequest(e Event) *Request {
	return &Request{
		Event: e,
	}
}
