package service

import (
	"context"
	"log"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

// Database models (use your preferred DB - PostgreSQL, MySQL, MongoDB, etc.)
type DeviceToken struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Token     string    `json:"token" db:"token"`
	Platform  string    `json:"platform" db:"platform"` // "ios", "android", "web"
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Request structures
type RegisterTokenRequest struct {
	UserID   int    `json:"user_id" validate:"required"`
	Token    string `json:"token" validate:"required"`
	Platform string `json:"platform" validate:"required,oneof=ios android web"`
}

type SendNotificationRequest struct {
	UserID int               `json:"user_id" validate:"required"`
	Title  string            `json:"title" validate:"required"`
	Body   string            `json:"body" validate:"required"`
	Data   map[string]string `json:"data,omitempty"`
}

// Firebase service
type FirebaseService struct {
	client *messaging.Client
	ctx    context.Context
}

func InitFirebase() (*FirebaseService, error) {
	ctx := context.Background()

	// Option 1: Using service account JSON file
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")

	// Option 2: Using environment variable
	// Set GOOGLE_APPLICATION_CREDENTIALS=/path/to/serviceAccountKey.json
	// opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

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

// Database interface (implement with your preferred DB)
type TokenRepository interface {
	SaveToken(token DeviceToken) error
	GetTokensByUserID(userID int) ([]DeviceToken, error)
	DeleteToken(tokenString string) error
	UpdateTokenStatus(tokenString string, isActive bool) error
	TokenExists(userID int, tokenString string) (bool, error)
}

// Example PostgreSQL implementation
type PostgresTokenRepo struct {
	// Add your DB connection here (e.g., *sql.DB or *sqlx.DB)
}

func (r *PostgresTokenRepo) SaveToken(token DeviceToken) error {
	// SQL: INSERT INTO device_tokens (user_id, token, platform, is_active, created_at, updated_at)
	// VALUES ($1, $2, $3, $4, $5, $6)
	// ON CONFLICT (user_id, token) DO UPDATE SET updated_at = $6, is_active = true
	return nil
}

func (r *PostgresTokenRepo) GetTokensByUserID(userID int) ([]DeviceToken, error) {
	// SQL: SELECT * FROM device_tokens WHERE user_id = $1 AND is_active = true
	return []DeviceToken{}, nil
}

func (r *PostgresTokenRepo) DeleteToken(tokenString string) error {
	// SQL: DELETE FROM device_tokens WHERE token = $1
	return nil
}

func (r *PostgresTokenRepo) UpdateTokenStatus(tokenString string, isActive bool) error {
	// SQL: UPDATE device_tokens SET is_active = $1 WHERE token = $2
	return nil
}

func (r *PostgresTokenRepo) TokenExists(userID int, tokenString string) (bool, error) {
	// SQL: SELECT EXISTS(SELECT 1 FROM device_tokens WHERE user_id = $1 AND token = $2)
	return false, nil
}

func main() {
	app := fiber.New()

	// Initialize Firebase
	firebaseService, err := InitFirebase()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	// Initialize your database
	tokenRepo := &PostgresTokenRepo{} // Replace with your actual DB implementation

	// API Routes

	// Register device token
	app.Post("/api/tokens/register", func(c *fiber.Ctx) error {
		var req RegisterTokenRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Check if token already exists
		exists, err := tokenRepo.TokenExists(req.UserID, req.Token)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Database error"})
		}

		deviceToken := DeviceToken{
			UserID:    req.UserID,
			Token:     req.Token,
			Platform:  req.Platform,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := tokenRepo.SaveToken(deviceToken); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save token"})
		}

		return c.JSON(fiber.Map{
			"message": "Token registered successfully",
			"exists":  exists,
		})
	})

	// Send notification to specific user
	app.Post("/api/notifications/send", func(c *fiber.Ctx) error {
		var req SendNotificationRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Get all active tokens for the user
		tokens, err := tokenRepo.GetTokensByUserID(req.UserID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to get tokens"})
		}

		if len(tokens) == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "No active tokens found for user"})
		}

		// Extract token strings
		tokenStrings := make([]string, 0, len(tokens))
		for _, t := range tokens {
			tokenStrings = append(tokenStrings, t.Token)
		}

		// Send notification
		response, err := firebaseService.SendToMultipleTokens(tokenStrings, req.Title, req.Body, req.Data)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to send notification"})
		}

		// Handle invalid tokens (tokens that failed)
		if response.FailureCount > 0 {
			for idx, resp := range response.Responses {
				if !resp.Success {
					// Mark token as inactive or delete it
					tokenRepo.UpdateTokenStatus(tokenStrings[idx], false)
				}
			}
		}

		return c.JSON(fiber.Map{
			"message":       "Notification sent",
			"success_count": response.SuccessCount,
			"failure_count": response.FailureCount,
		})
	})

	// Broadcast to all users (using topic)
	app.Post("/api/notifications/broadcast", func(c *fiber.Ctx) error {
		type BroadcastRequest struct {
			Topic string            `json:"topic" validate:"required"`
			Title string            `json:"title" validate:"required"`
			Body  string            `json:"body" validate:"required"`
			Data  map[string]string `json:"data,omitempty"`
		}

		var req BroadcastRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		response, err := firebaseService.SendToTopic(req.Topic, req.Title, req.Body, req.Data)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to send notification"})
		}

		return c.JSON(fiber.Map{
			"message":    "Broadcast sent",
			"message_id": response,
		})
	})

	// Delete token
	app.Delete("/api/tokens/:token", func(c *fiber.Ctx) error {
		token := c.Params("token")

		if err := tokenRepo.DeleteToken(token); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete token"})
		}

		return c.JSON(fiber.Map{"message": "Token deleted successfully"})
	})

	log.Fatal(app.Listen(":3000"))
}
