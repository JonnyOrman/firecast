package firecast

import "google.golang.org/genproto/googleapis/firestore/v1"

type FirestoreEvent struct {
	OldValue   firestore.Document `json:"oldValue"`
	Value      firestore.Document `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}
