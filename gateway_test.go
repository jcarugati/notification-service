package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLocalGateway(t *testing.T) {
	gateway := NewLocalGateway()
	assert.NotNil(t, gateway)
}

func TestLocalGateway_Send(t *testing.T) {
	gateway := NewLocalGateway()
	err := gateway.Send(&Payload{
		UserID:  "user_id",
		Message: "message",
	})
	assert.Nil(t, err)
}
