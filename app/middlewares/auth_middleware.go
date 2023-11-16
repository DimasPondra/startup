package middlewares

import (
	"net/http"
	"startup/app/helpers"
	"startup/app/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(authService services.AuthService, userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			res := helpers.ResponseAPI("Unauthorized.", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		// Split token from header
		var tokenString string

		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) == 2 {
			tokenString = splitToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			res := helpers.ResponseAPI("Unauthorized.", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			res := helpers.ResponseAPI("Unauthorized.", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			res := helpers.ResponseAPI("Unauthorized.", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}