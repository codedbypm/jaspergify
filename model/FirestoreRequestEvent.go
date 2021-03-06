package model

import "time"

// FirestoreRequestEvent is the payload of a Firestore event.
type FirestoreRequestEvent struct {
	Value FirestoreRequestValue `json:"value"`
}

// FirestoreRequestValue holds Firestore fields.
type FirestoreRequestValue struct {
	CreateTime time.Time        `json:"createTime"`
	Fields     FirestoreRequest `json:"fields"`
	Name       string           `json:"name"`
	UpdateTime time.Time        `json:"updateTime"`
}

// FirestoreRequest is awesome
type FirestoreRequest struct {
	GiphyIdentifier StringValue `json:"giphyId"`
}
