package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware validates JWT via Auth Service
func AuthMiddleware(authServiceURL string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract token from Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid authorization header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header format")
			}
			token := parts[1]

			// Validate token via Auth Service
			userID, err := validateTokenWithAuthService(authServiceURL, token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}

			// Pass user ID to context
			c.Set("user_id", userID)

			return next(c)
		}
	}
}

// validateTokenWithAuthService sends the token to Auth Service for validation
func validateTokenWithAuthService(authServiceURL, token string) (string, error) {
	// Create a request to Auth Service
	req, err := http.NewRequest(http.MethodPost, authServiceURL+"/validate", strings.NewReader(`{"token":"`+token+`"}`))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("invalid token")
	}

	// Extract user ID from Auth Service response
	var result struct {
		UserID string `json:"user_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.UserID, nil
}
