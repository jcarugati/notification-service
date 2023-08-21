package main

import "fmt"

// LocalGateway represents a local gateway to send notifications.
type LocalGateway struct{}

// Send outputs a notification for a given payload to the console.
func (lg *LocalGateway) Send(payload *Payload) error {
	fmt.Println(fmt.Sprintf("Sending notification to %s with message %s", payload.UserID, payload.Message))
	return nil
}

// Payload represents the data structure for a notification payload.
type Payload struct {
	UserID  string
	Message string
}

// NewLocalGateway initializes a new LocalGateway.
func NewLocalGateway() *LocalGateway {
	return &LocalGateway{}
}
