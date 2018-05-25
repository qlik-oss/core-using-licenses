package test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	enigma "github.com/qlik-oss/enigma-go"
)

func getTestHost() string {
	testHost := os.Getenv("TEST_HOST")
	if testHost == "" {
		return "localhost"
	}
	return testHost
}

var (
	ctx     = context.Background()
	headers = make(http.Header, 1)
	Host    = getTestHost()
)

func ConnectToEngineAndReturnOnConnectedEventMessage(ctx context.Context, sessionID int, headers http.Header) (string, error) {
	headers.Set("X-Qlik-Session", fmt.Sprintf("%d", sessionID))
	global, err := enigma.Dialer{}.Dial(ctx, fmt.Sprintf("ws://%s:19076/app/engineData/", Host), headers)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	sessionMessages := global.SessionMessageChannel()
	message := "Channel Closed before reciving OnConnected Message from Engine"
	err = errors.New(message)

	for sessionEvent := range sessionMessages {
		if sessionEvent.Topic == "OnConnected" {
			var parsed map[string]string
			err := json.Unmarshal(sessionEvent.Content, &parsed)
			if err != nil {
				return "", err
			}
			return parsed["qSessionState"], nil
		}
	}
	return message, err
}
