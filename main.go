package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	VerifyToken string
	VerifyType  string
)

// GetEnv get verify token and type from enviroments
func GetEnv() {
	VerifyToken = os.Getenv("GorillaToken")
	VerifyType = os.Getenv("GorillaType")

	if VerifyToken == "" {
		panic("please set env GorillaToken")
	}
	if VerifyType == "" {
		panic("please set env GorillaType")
	}
}

type ChallengeRequest struct {
	Token string `json:"token"`
	Type  string `json:"type"`
	Challenge
}

type Challenge struct {
	Challenge string `json:"challenge"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}

func NewChallenge(challenge string) Challenge {
	return Challenge{
		Challenge: challenge,
	}
}

func NewErrorMessage(msg string) ErrorMessage {
	return ErrorMessage{
		Error: msg,
	}
}

func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func WelcomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] verification start from %s\n", TimeNow(), r.RemoteAddr)

	defer func() {
		fmt.Printf("[%s] verification end from %s\n", TimeNow(), r.RemoteAddr)
	}()

	// get request body
	req := new(ChallengeRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		msg := fmt.Sprintf("Bad request data %+v", r.Body)
		fmt.Printf("[%s] %s", TimeNow(), msg)
		NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
		return
	}

	// validation
	if req.Token != VerifyToken {
		msg := fmt.Sprintf("Invalid token: %s", req.Token)
		fmt.Printf("[%s] %s\n", TimeNow(), msg)
		NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
		return
	}

	if req.Type != VerifyType {
		msg := fmt.Sprintf("Invalid type: %s", req.Type)
		fmt.Printf("[%s] %s\n", TimeNow(), msg)
		NewResponse(w, http.StatusBadRequest, NewErrorMessage(msg))
		return
	}

	// success verification
	fmt.Printf("[%s] success verification from %s\n", TimeNow(), r.RemoteAddr)
	NewResponse(w, http.StatusOK, NewChallenge(req.Challenge.Challenge))
}

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
	fmt.Printf("[%s] start server in port: %s\n", TimeNow(), port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
