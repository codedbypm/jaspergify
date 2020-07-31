package model

import "time"

// Request models a request to Jaspergify a gif
type Request struct {
	GiphyIdentifier string        `firestore:"giphyId"`
	Timestamp       time.Time     `firestore:"time"`
	Status          RequestStatus `firestore:"status"`
}

// FirestoreRequest is awesome
type FirestoreRequest struct {
	GiphyIdentifier StringValue    `firestore:"giphyId"`
	Timestamp       TimestampValue `firestore:"time"`
	Status          StringValue    `firestore:"status"`
}

// StringValue is awesome
type StringValue struct {
	StringValue string `json:"stringValue"`
}

// TimestampValue is awesome
type TimestampValue struct {
	TimestampValue time.Time `json:"timestampValue"`
}

// RequestStatus models the status of a request to Jaspergify a gif
type RequestStatus string

const (
	// Received means the request has been received
	Received RequestStatus = "received"
	// Ongoing means the request started but not yet finished
	Ongoing = "ongoing"
	// Done means the request has been executed
	Done = "done"
)
