package test

import (
	"fmt"
	"testing"

	"github.com/qlik-oss/enigma-go"
	"github.com/stretchr/testify/assert"
)

func TestThatFiveButNotMoreConcurentSessionsWork(t *testing.T) {
	var numberOfAllowedConnectionsWithNoLicenseService = 5

	for i := 0; i < numberOfAllowedConnectionsWithNoLicenseService; i++ {

		message, err := ConnectToEngineAndReturnOnConnectedEventMessage(ctx, i, headers)
		assert.Equal(t, "SESSION_CREATED", message)
		assert.Nil(t, err, "Connecting to engine should not give an error")

	}

	//One more connection than allowed, this should not be allowed
	headers.Set("X-Qlik-Session", fmt.Sprintf("%d", numberOfAllowedConnectionsWithNoLicenseService+1))
	global, _ := enigma.Dialer{}.Dial(ctx, "ws://localhost:19076/app/engineData/", headers)
	//try to create a session app to see that the WS has been closed (In the future engine should send the SESSION_ERROR_NO_LICENSE event here.)
	_, err := global.CreateSessionApp(ctx)
	assert.NotNil(t, err, "This session should not be allowed to be open")

}
