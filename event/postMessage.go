package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorilla_bot/common"
	"io/ioutil"
	"net/http"
)

// PostMessageURL send message url
var PostMessageURL = "https://slack.com/api/chat.postMessage"

// new http client
var client = &http.Client{}

// Message slack message
// https://api.slack.com/methods/chat.postMessage
type Message struct {
	Token   string `json:"-"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

// NewMessage New send message
func NewMessage(token, channel, text string) *Message {
	return &Message{
		Token:   token,
		Channel: channel,
		Text:    text,
	}
}

// Post send messagea to member
func (m *Message) Post() {

	body, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%s] request body: \n%s\n", common.TimeNow(), common.FormatStringJoin(string(body)))

	req, err := http.NewRequest(
		http.MethodPost,
		PostMessageURL,
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.Token))
	resp, err := client.Do(req)

	if err != nil {
		msg := fmt.Sprintf("request error cause [%s]", err)
		fmt.Printf("[%s] %s\n", common.TimeNow(), msg)
		return
	}
	defer resp.Body.Close()

	// output request body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[%s] Bad response body [%s]\n", common.TimeNow(), err)
		return
	}

	fmt.Printf("[%s] response body: \n%s\n", common.TimeNow(), common.FormatStringJoin(string(body)))
}
