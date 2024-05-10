package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

func RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c != nil {
				return next(c)
			}

			authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization header")
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization format")
			}
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing or malformed Authorization header"})
				return nil
			}
			// Extract the token from the header
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			// Parse the JWT token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Provide the key for verifying the token's signature
				// (replace with your actual key)
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
				return nil
			}
			// If the token is valid, proceed with the next middleware/handler

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Set("jwtClaims", claims)

			}
			return next(c)
		}
	}
	// Get the Authorization header value

	// Check if the header is empty or doesn't start with "Bearer "

}
