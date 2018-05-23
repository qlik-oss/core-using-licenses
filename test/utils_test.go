package test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	enigma "github.com/qlik-oss/enigma-go"
)

var (
	ctx     = context.Background()
	headers = make(http.Header, 1)
)

func ConnectToEngineAndReturnOnConnectedEventMessage(ctx context.Context, sessionID int, headers http.Header) (string, error) {
	headers.Set("X-Qlik-Session", fmt.Sprintf("%d", sessionID))
	global, _ := enigma.Dialer{}.Dial(ctx, "ws://localhost:19076/app/engineData/", headers)

	sessionMessages := global.SessionMessageChannel()
	message := "Channel Closed before reciving OnConnected Message from Engine"
	err := errors.New(message)

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
