package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("your-secret-key")

// CustomClaims is a custom structure for JWT claims
type CustomClaims struct {
	Username string `json:"username"`
	UserID   string `json:"userID"`
	jwt.StandardClaims
}

func ExtractUserInfoFromToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		fmt.Print("Username:", claims.Username)
		fmt.Print("UserID:", claims.UserID)
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Pengecekan otorisasi
		username := claims.(CustomClaims).Username
		userID := claims.(CustomClaims).UserID
		// Lakukan pengecekan otorisasi sesuai kebutuhan aplikasi

		// Set Username dan UserID dalam context untuk digunakan oleh handler selanjutnya
		c.Set("username", username)
		c.Set("userID", userID)

		c.Next()
	}
}

// CreateToken generates a new JWT token
func CreateToken(username string, userID string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &CustomClaims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
