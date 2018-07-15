package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	MissingHeaderError = errors.New("Missing `Authorization` header.")
)

type Context struct {
	ID       uint64
	Username string
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Check the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates token with secret key
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse token
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return ctx, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return ctx, err
	}

	ctx.ID = uint64(claims["id"].(float64))
	ctx.Username = claims["username"].(string)

	return ctx, nil
}

// ParseRequest gets token from header

func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// Load JWT secret from config
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, MissingHeaderError
	}

	var t string
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	// Load JWT from config if secret is not specified
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	// Define token content
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(), // JWT effective time
		"iat":      time.Now().Unix(), // JWT issuance time
	})

	tokenString, err = token.SignedString([]byte(secret))

	return tokenString, err
}
