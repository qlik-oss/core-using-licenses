package test

import (
	"testing"

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
	message, err := ConnectToEngineAndReturnOnConnectedEventMessage(ctx, numberOfAllowedConnectionsWithNoLicenseService+1, headers)
	assert.Equal(t, "SESSION_ERROR_LIMIT_EXCEEDED", message)
	assert.Nil(t, err, "This session should not be allowed to be open")

}
