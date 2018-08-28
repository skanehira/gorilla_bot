package server

import (
	"encoding/json"
	"fmt"
	"gorilla_bot/bot"
	"gorilla_bot/common"
	"gorilla_bot/types"
	"io/ioutil"
	"log"
	"net/http"
)

// ErrorMessage Error Message
type ErrorMessage struct {
	Error string `json:"error"`
}

// NewErrorMessage New Error Response
func NewErrorMessage(msg string) ErrorMessage {
	return ErrorMessage{
		Error: msg,
	}
}

// NewResponse HTTP Response
func NewResponse(w http.ResponseWriter, code int, data interface{}) {
	body, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(body)
}

// Handler url_verification or someone events
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	// output request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("Bad request body [%s]", err)
		log.Printf("[%s] %s\n", common.TimeNow(), msg)
		NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
		return
	}

	log.Printf("[%s] request body: \n%s\n", common.TimeNow(), common.FormatStringJoin(string(body)))

	// get request type
	rtype := new(types.RequestType)

	if err := json.Unmarshal(body, rtype); err != nil {
		msg := fmt.Sprintf("Bad request body [%s]", err)
		log.Printf("[%s] %s\n", common.TimeNow(), msg)
		NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
		return
	}

	if !types.IsValidRequestType(rtype.Type) {
		msg := fmt.Sprintf("Bad request type [%s]", rtype.Type)
		log.Printf("[%s] %s\n", common.TimeNow(), msg)
		NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
		return
	}

	// if request type is url_verification
	if rtype.Type == types.URL_VERIFICATION {
		log.Printf("[%s] verification start from %s\n", common.TimeNow(), r.RemoteAddr)

		defer func() {
			log.Printf("[%s] verification end from %s\n", common.TimeNow(), r.RemoteAddr)
		}()

		req := new(types.ChallengeRequest)
		if err := json.Unmarshal(body, req); err != nil {
			msg := fmt.Sprintf("Bad request data %+v", r.Body)
			log.Printf("[%s] %s\n", common.TimeNow(), msg)
			NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
			return
		}

		// validation token
		if req.Token != s.URLVerifyToken {
			msg := fmt.Sprintf("Invalid token: [%s]", req.Token)
			log.Printf("[%s] %s\n", common.TimeNow(), msg)
			NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
			return
		}

		// success verification
		log.Printf("[%s] success verification from %s\n", common.TimeNow(), r.RemoteAddr)
		NewResponse(w, http.StatusOK, types.NewChallenge(req.Challenge.Challenge))
		return
	}

	// if request type is event_callback
	if rtype.Type == types.EVENT_CALLBACK {
		log.Printf("[%s] event_callback start from %s\n", common.TimeNow(), r.RemoteAddr)

		defer func() {
			log.Printf("[%s] event_callback end from %s\n", common.TimeNow(), r.RemoteAddr)
		}()

		req := types.NewRequest(types.NewMemberJoinedChannel())
		if err := json.Unmarshal(body, req); err != nil {
			msg := fmt.Sprintf("Bad request data %+v", r.Body)
			log.Printf("[%s] %s\n", common.TimeNow(), msg)
			NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
			return
		}

		// validation token
		if req.Token != s.URLVerifyToken {
			msg := fmt.Sprintf("Invalid token: [%s]", req.Token)
			log.Printf("[%s] %s\n", common.TimeNow(), msg)
			NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
			return
		}

		// send message
		bot := bot.New(s.Config, s.URLVerifyToken, s.AuthorizationToken, req.Event.ToMap()["User"].(string))
		bot.SendMessage(bot.ReadMessageFromFile(s.MessageFile))
	}
}
