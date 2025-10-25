package firebase

import (
	"context"
	"dubai-auto/internal/config"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

// Firebase service
type FirebaseService struct {
	client *messaging.Client
	ctx    context.Context
}

func InitFirebase(cfg *config.Config) (*FirebaseService, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(cfg.FIREBASE_ACCOUNT_FILE)
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(ctx)

	if err != nil {
		return nil, err
	}

	return &FirebaseService{
		client: client,
		ctx:    ctx,
	}, nil
}

// Send notification to a specific device token
func (fs *FirebaseService) SendToToken(token, title, body string, data map[string]string) (string, error) {
	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				Sound: "default",
			},
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Sound: "default",
				},
			},
		},
	}

	response, err := fs.client.Send(fs.ctx, message)

	if err != nil {
		return "", err
	}

	return response, nil
}

// Send notification to multiple tokens
func (fs *FirebaseService) SendToMultipleTokens(tokens []string, title, body string, data map[string]string) (*messaging.BatchResponse, error) {
	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
	}

	response, err := fs.client.SendEachForMulticast(fs.ctx, message)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Send to topic (for broadcast notifications)
func (fs *FirebaseService) SendToTopic(topic, title, body string, data map[string]string) (string, error) {
	message := &messaging.Message{
		Topic: topic,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
	}

	response, err := fs.client.Send(fs.ctx, message)

	if err != nil {
		return "", err
	}

	return response, nil
}

// Subscribe tokens to a topic
func (fs *FirebaseService) SubscribeToTopic(tokens []string, topic string) error {
	_, err := fs.client.SubscribeToTopic(fs.ctx, tokens, topic)
	return err
}
