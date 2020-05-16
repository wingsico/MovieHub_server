package token

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/wingsico/movie_server/conf"
	"github.com/wingsico/movie_server/errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Context struct {
	ID       uint64
	Username string
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil

	} else {
		return ctx, err
	}
}

func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	secret := os.Getenv("secret")

	if len(header) == 0 {
		return &Context{}, errors.ErrMissingAuthHeader
	}

	var t string
	_, _ = fmt.Sscanf(header, "Bearer %s", &t)

	return Parse(t, secret)
}

func Sign(c Context, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = os.Getenv("secret")
	}
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(24 * time.Hour * time.Duration(1)).Unix(),
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
