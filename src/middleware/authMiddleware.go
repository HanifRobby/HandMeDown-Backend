package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("your-secret-key")

// CustomClaims is a custom structure for JWT claims
type CustomClaims struct {
	Username string `json:"username"`
	UserID   uint   `json:"userID"`
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
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Memeriksa apakah header "Authorization" memiliki skema "Bearer"
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token format"})
			c.Abort()
			return
		}

		// Mendapatkan token setelah skema "Bearer"
		tokenString := authHeader[7:]

		// Ekstrak informasi pengguna dari token
		claims, err := ExtractUserInfoFromToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set claims dalam konteks untuk digunakan oleh handler selanjutnya jika diperlukan
		c.Set("claims", claims)
		fmt.Println("UserID in AuthorizationMiddleware:", claims.UserID)
		c.Set("userID", claims.UserID)

		c.Next()
	}
}

// CreateToken generates a new JWT token
func CreateToken(username string, userID uint) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
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
