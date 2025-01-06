package middleware

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(authServiceURL string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid authorization header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header format")
			}
			token := parts[1]

			log.Printf("Validating token: %s", token)

			userID, email, err := validateTokenWithAuthService(authServiceURL, token)
			if err != nil {
				log.Printf("Token validation failed: %v", err)
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}

			log.Printf("Token valid. user_id: %s, email: %s", userID, email)

			c.Set("user_id", userID)
			c.Set("email", email)

			return next(c)
		}
	}
}

func validateTokenWithAuthService(authServiceURL, token string) (string, string, error) {
	req, err := http.NewRequest(http.MethodPost, authServiceURL+"/validate", strings.NewReader(`{"token":"`+token+`"}`))
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return "", "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Request to auth service failed: %v", err)
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Auth service returned status: %d", resp.StatusCode)
		return "", "", errors.New("invalid token")
	}

	var result struct {
		UserID string `json:"user_id"`
		Email  string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Failed to decode response: %v", err)
		return "", "", err
	}

	return result.UserID, result.Email, nil
}
