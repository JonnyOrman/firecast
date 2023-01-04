package firecast

import "time"

type Event struct {
	DocumentID string
	Operation  string
	Timestamp  time.Time
}

func NewEvent(
	documentID string,
	operation string,
	timestamp time.Time) *Event {
	this := new(Event)

	this.DocumentID = documentID
	this.Operation = operation
	this.Timestamp = timestamp

	return this
}
