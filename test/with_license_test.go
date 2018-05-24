package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatMoreThanFiveSessionsWorkWithALicense(t *testing.T) {

	for i := 0; i < 10; i++ {
		message, err := ConnectToEngineAndReturnOnConnectedEventMessage(ctx, i, headers)
		assert.Equal(t, "SESSION_CREATED", message)
		assert.Nil(t, err, "Connecting to engine should not give an error")
	}

}
