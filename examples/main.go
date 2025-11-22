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
	simulateComplexData()
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

func simulateComplexData() {
	// Test arrays, nested objects, null values, and booleans
	slog.Info("User profile fetched",
		"user_id", "user_999",
		"tags", []string{"premium", "verified", "developer"},
		"permissions", []string{"read", "write", "admin"},
		"scores", []int{85, 92, 78, 95},
	)

	slog.Debug("Cache status check",
		"redis_connected", true,
		"memcached_connected", false,
		"fallback_enabled", true,
		"primary_cache", "redis",
		"backup_cache", nil,
	)

	slog.Info("Order details",
		"order_id", "ord_12345",
		"items", []map[string]any{
			{"name": "Laptop", "price": 999.99, "quantity": 1, "in_stock": true},
			{"name": "Mouse", "price": 29.99, "quantity": 2, "in_stock": true},
			{"name": "Monitor", "price": 349.99, "quantity": 1, "in_stock": false},
		},
		"discount_applied", false,
		"total", 1409.96,
		"shipping_address", nil,
	)

	slog.Warn("Feature flags evaluated",
		"user_id", "user_555",
		"flags", map[string]any{
			"new_ui":           true,
			"beta_features":    false,
			"dark_mode":        true,
			"analytics":        true,
			"experimental_api": nil,
		},
		"evaluation_ms", 12,
	)

	slog.Info("User profile fetched",
		"user_id", "user_999",
		"tags", []string{"premium", "verified", "developer"},
		"permissions", []string{"read", "write", "admin"},
		"scores", []int{85, 92, 78, 95},
		"recent_logins", []int64{1699564800, 1699651200, 1699737600},
	)

	slog.Debug("Feature toggles",
		"enabled_features", []string{"dark_mode", "notifications", "analytics"},
		"disabled_features", []string{},
		"beta_flags", []bool{true, false, true, true},
	)

	slog.Error("Batch job failed",
		"job_id", "batch_777",
		"total_items", 1000,
		"processed", 847,
		"failed", 153,
		"success_rate", 0.847,
		"errors", []map[string]any{
			{"item_id": "item_1", "error": "invalid format", "retry": true},
			{"item_id": "item_2", "error": "timeout", "retry": false},
			{"item_id": "item_3", "error": "not found", "retry": nil},
		},
		"will_retry", true,
		"next_attempt", nil,
	)
}
