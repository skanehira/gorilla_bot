package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorilla_bot/common"
	"gorilla_bot/config"
	"gorilla_bot/types"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// new http client
var client = new(http.Client)

// Bot slack bot
type Bot struct {
	config.Config
	URLVerifyToken     string
	AuthorizationToken string
	ToChannel          string
}

// New new slack bot
func New(config config.Config, urlToken, authToken, toChnnel string) *Bot {
	return &Bot{
		Config:             config,
		URLVerifyToken:     urlToken,
		AuthorizationToken: authToken,
		ToChannel:          toChnnel,
	}
}

// ReadMessageFromFile read send message from file
func (b *Bot) ReadMessageFromFile(file string) string {

	// read mesasge template from file
	data, err := common.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// get channel list. add to message
	return b.addChannelsToMessage(string(data))
}

// SendMessage send message to user in DM
func (b *Bot) SendMessage(message string) {

	// marshal request body
	body, err := json.Marshal(types.NewMessage(b.ToChannel, message))

	if err != nil {
		log.Printf("[%s] create send request body is failed: %s", common.TimeNow(), err)
		return
	}

	// new request
	req, err := http.NewRequest(
		http.MethodPost,
		b.PostMessageURL,
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

// GetChannelList get slack workspace channel list
func (b *Bot) GetChannelList() *types.Channels {
	req, err := http.NewRequest(
		http.MethodGet,
		b.ChannelListURL,
		nil,
	)

	if err != nil {
		log.Printf("[%s] create send request body is failed: %s", common.TimeNow(), err)
		return types.NewChannels()
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", b.AuthorizationToken))

	// get channel list
	res, err := client.Do(req)
	if err != nil {
		log.Printf("[%s] send request is failed: %s", common.TimeNow(), err)
		return types.NewChannels()
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	// new channel
	if err != nil {
		log.Printf("[%s] bad response body: %s", common.TimeNow(), err)
		return types.NewChannels()
	}

	// new channel
	channels := types.NewChannels()

	if err := json.Unmarshal(body, channels); err != nil {
		log.Printf("[%s] unmarshal body is failed: %s ", common.TimeNow(), err)
		return channels
	}

	return channels
}

// add channels to message
func (b *Bot) addChannelsToMessage(oldMessage string) string {
	// channels to strings
	var channels = []string{}

	for _, channel := range b.GetChannelList().Channels {
		channels = append(channels, fmt.Sprintf("#%-30s %s", channel.Name, channel.Purpose.Value))
	}

	// channel to text
	channelText := strings.Join(channels, "\n")
	message := strings.Replace(oldMessage, "{channel_list}", channelText, -1)

	return message
}
