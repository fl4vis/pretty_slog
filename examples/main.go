package main

import (
	"log/slog"
	"time"

	"github.com/fl4vis/pretty_slog"
)

func main() {
	logger := slog.New(pretty_slog.NewHandler(&slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
	}))

	// Set this logger as the default for the entire application.
	// After this, any package using slog.Info(), slog.Error(), etc
	// will automatically use our pretty_slog handler instead of the standard one.
	// This means you don't need to pass the logger around everywhere.
	slog.SetDefault(logger)

	// Simulate API request lifecycle
	simulateAPIRequest()
	simulateDatabaseOperation()
	simulateAuthFlow()
	simulateErrorScenario()

}

func simulateAPIRequest() {
	requestID := "req_abc123"
	userID := "user_456"

	slog.Debug("Incoming request",
		"request_id", requestID,
		"method", "POST",
		"path", "/api/v1/users",
		"ip", "192.168.1.100",
	)

	slog.Info("Request processed successfully",
		"request_id", requestID,
		"user_id", userID,
		"status", 200,
		"duration_ms", 45,
	)
}

func simulateDatabaseOperation() {
	slog.Debug("Executing database query",
		"query", "SELECT * FROM users WHERE active = ?",
		"params", []any{true},
		"connection", "postgres://localhost:5432",
	)

	slog.Info("Database query completed",
		"rows_affected", 150,
		"duration_ms", 23,
		"cache_hit", false,
	)
}

func simulateAuthFlow() {
	slog.Info("User authentication attempt",
		"username", "john.doe",
		"ip", "203.0.113.45",
		"user_agent", "Mozilla/5.0",
	)

	slog.Warn("Rate limit approaching",
		"user_id", "user_789",
		"requests_count", 95,
		"limit", 100,
		"window", "1m",
		"remaining", 5,
	)
}

func simulateErrorScenario() {
	// Simulate various error conditions
	
	slog.Warn("Deprecated API version used",
		"version", "v1",
		"endpoint", "/api/v1/legacy",
		"user_id", "user_321",
		"migration_deadline", time.Now().Add(30*24*time.Hour).Format(time.RFC3339),
	)

	slog.Error("Database connection failed",
		"error", "connection timeout",
		"host", "db.example.com:5432",
		"retry_count", 3,
		"max_retries", 5,
	)

	slog.Error("Payment processing failed",
		"transaction_id", "txn_xyz789",
		"amount", 99.99,
		"currency", "USD",
		"error", "insufficient funds",
		"user_id", "user_654",
	)

	slog.Error("External API call failed",
		"service", "stripe",
		"endpoint", "https://api.stripe.com/v1/charges",
		"status_code", 503,
		"error", "service unavailable",
		"retry_after", 30,
	)
}
