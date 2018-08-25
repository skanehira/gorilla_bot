package bot

import (
	"bytes"
	"encoding/json"
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

// ReadMessageFromFile read send message from file
func (b *Bot) ReadMessageFromFile(file string) string {
	data, err := common.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(data)
}

// SendMessage send message to user in DM
func (b *Bot) SendMessage(message string) {

	// marshal request body
	body, err := json.Marshal(NewMessage(b.ToChannel, message))

	if err != nil {
		log.Printf("[%s] create send request body is failed: %s", common.TimeNow(), err)
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

	log.Printf("[%s] send request header: %s\n", common.TimeNow(), req.Header)
	log.Printf("[%s] send request body: \n%s\n", common.TimeNow(), common.FormatStringJoin(string(body)))

	// send message
	res, err := client.Do(req)

	if err != nil {
		log.Printf("[%s] send request is failed [%s]\n", common.TimeNow(), err)
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
