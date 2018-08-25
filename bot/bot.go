package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gorilla_bot/common"
	"io/ioutil"
	"log"
	"net/http"
)

// new http client
var client = new(http.Client)

// Bot slack bot
type Bot struct {
	URLVerifyToken     string
	AuthorizationToken string
	ToChannel          string
}

// New new slack bot
func New(urlToken, authToken, toChnnel string) *Bot {
	return &Bot{
		URLVerifyToken:     urlToken,
		AuthorizationToken: authToken,
		ToChannel:          toChnnel,
	}
}

// Authenticate check token
func (b *Bot) Authenticate(urlToken, authToken string) error {
	if !(b.AuthorizationToken == authToken && b.URLVerifyToken == urlToken) {
		return errors.New("Authentication failure")
	}

	return nil
}

// SendMessage send message to user in DM
func (b *Bot) SendMessage(message string) {

	// marshal request body
	body, err := json.Marshal(NewMessage(b.ToChannel, message))

	if err != nil {
		log.Printf("[%s] create request body is failed: %s", common.TimeNow(), err)
	}

	// new request
	req, err := http.NewRequest(
		http.MethodPost,
		PostMessageURL,
		bytes.NewBuffer(body),
	)

	// set request headers
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", b.AuthorizationToken))

	log.Printf("[%s] request header: %s\n", common.TimeNow(), req.Header)
	log.Printf("[%s] request body: \n%s\n", common.TimeNow(), common.FormatStringJoin(string(body)))

	// send message
	res, err := client.Do(req)

	if err != nil {
		log.Printf("[%s] request error cause [%s]\n", common.TimeNow(), err)
		return
	}

	defer res.Body.Close()

	// output request body
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("[%s] bad response body: %s\n", common.TimeNow(), err)
		return
	}

	log.Printf("[%s] response body: \n%s\n", common.TimeNow(), common.FormatStringJoin(string(body)))
}
