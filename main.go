package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gorilla_bot/common"
	"gorilla_bot/event"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	URLVerifyToken     string
	AuthorizationToken string
)

// GetEnv get verify token and type from enviroments
func GetEnv() {
	URLVerifyToken = os.Getenv("URLVerificationToken")
	AuthorizationToken = os.Getenv("AuthorizationToken")

	if URLVerifyToken == "" {
		panic("please set URLVerificationToken")
	}

	if AuthorizationToken == "" {
		panic("please set AuthorizationToken")
	}
}

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

// WelcomHandler url_verification or someone events
func WelcomHandler(w http.ResponseWriter, r *http.Request) {
	// output request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("Bad request body [%s]", err)
		fmt.Printf("[%s] %s\n", common.TimeNow(), msg)
		NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
		return
	}

	fmt.Printf("[%s] request body: \n%s\n", common.TimeNow(), common.FormatStringJoin(string(body)))

	// get request type
	rtype := new(event.RequestType)

	if err := json.Unmarshal(body, rtype); err != nil {
		msg := fmt.Sprintf("Bad request body [%s]", err)
		fmt.Printf("[%s] %s\n", common.TimeNow(), msg)
		NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
		return
	}

	if !event.IsValidRequestType(rtype.Type) {
		msg := fmt.Sprintf("Bad request type [%s]", rtype.Type)
		fmt.Printf("[%s] %s\n", common.TimeNow(), msg)
		NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
		return
	}

	// if request type is url_verification
	if rtype.Type == event.URL_VERIFICATION {
		fmt.Printf("[%s] verification start from %s\n", common.TimeNow(), r.RemoteAddr)

		defer func() {
			fmt.Printf("[%s] verification end from %s\n", common.TimeNow(), r.RemoteAddr)
		}()

		req := new(event.ChallengeRequest)
		if err := json.Unmarshal(body, req); err != nil {
			msg := fmt.Sprintf("Bad request data %+v", r.Body)
			fmt.Printf("[%s] %s\n", common.TimeNow(), msg)
			NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
			return
		}

		// output unmarshaled data
		fmt.Printf("[%s] request=%+v\n", common.TimeNow(), req)

		// validation token
		if req.Token != URLVerifyToken {
			msg := fmt.Sprintf("Invalid token: [%s]", req.Token)
			fmt.Printf("[%s] %s\n", common.TimeNow(), msg)
			NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
			return
		}

		// success verification
		fmt.Printf("[%s] success verification from %s\n", common.TimeNow(), r.RemoteAddr)
		NewResponse(w, http.StatusOK, event.NewChallenge(req.Challenge.Challenge))
		return
	}

	// if request type is event_callback
	if rtype.Type == event.EVENT_CALLBACK {
		fmt.Printf("[%s] event_callback start from %s\n", common.TimeNow(), r.RemoteAddr)

		defer func() {
			fmt.Printf("[%s] event_callback end from %s\n", common.TimeNow(), r.RemoteAddr)
		}()

		req := event.NewRequest(event.NewMemberJoinedChannel())
		if err := json.Unmarshal(body, req); err != nil {
			msg := fmt.Sprintf("Bad request data %+v", r.Body)
			fmt.Printf("[%s] %s\n", common.TimeNow(), msg)
			NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
			return
		}

		// output unmarshaled data
		fmt.Printf("[%s] request=%+v\n", common.TimeNow(), req)
		fmt.Printf("[%s] event=%+v\n", common.TimeNow(), req.Event)

		// validation token
		if req.Token != URLVerifyToken {
			msg := fmt.Sprintf("Invalid token: [%s]", req.Token)
			fmt.Printf("[%s] %s\n", common.TimeNow(), msg)
			NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
			return
		}

		// send message
		message := event.NewMessage(AuthorizationToken, req.Event.Name(), "welcom gorilla-lab")
		message.Post()
		return
	}
}

// NewResponse HTTP Response
func NewResponse(w http.ResponseWriter, code int, data interface{}) {
	body, _ := json.Marshal(data)

	w.WriteHeader(code)
	w.Write(body)
}

func main() {
	// get enviroment token and type
	GetEnv()

	// parse args
	flag.Parse()
	if len(flag.Args()) < 1 {
		panic("please specified server port")
	}

	port := flag.Arg(0)

	// regist handler
	http.HandleFunc("/slack/gorilla", WelcomHandler)

	// server start
	fmt.Printf("[%s] start server in port: %s\n", common.TimeNow(), port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
