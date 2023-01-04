package firecast

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
)

func Broadcast(
	ctx context.Context,
	e FirestoreEvent,
	documentID string,
	projectID string,
	topicID string,
	operation string) error {
	client, _ := pubsub.NewClient(ctx, projectID)

	defer client.Close()

	event := NewEvent(
		documentID,
		operation,
		e.Value.CreateTime.AsTime(),
	)

	data, _ := json.Marshal(event)

	topic := client.Topic(topicID)

	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	_, _ = result.Get(ctx)

	return nil
}
