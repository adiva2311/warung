package middlewares

import (
	"net/http"
	"strings"

	"warung/helpers"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Missing or invalid token",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return helpers.GetSecretKey(), nil
		}, jwt.WithValidMethods([]string{"HS256"}))

		if err != nil || token == nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "invalid or expired token"})
		}

		// safe conversion
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "cannot parse JWT claims"})
		}

		userID, ok := claims["id"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "invalid user_id in token"})
		}

		role, ok := claims["role"].(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "invalid role in token"})
		}

		// Simpan user_id ke context
		c.Set("user_id", userID)
		c.Set("role", role)

		return next(c)
	}
}
