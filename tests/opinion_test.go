//go:build e2e
// +build e2e

package tests

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("misionimposible"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func TestPostComment(t *testing.T) {
	t.Run("puedo post opinion", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", "bearer "+createToken()).
			SetBody(`{"asunto": "asunto test post", "contenido": "contenido test post", "autor": "resty"`).
			Post("http://localhost:8080/api/v1/opinion")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("no puedo post opinion sin JWT", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"asunto": "asunto test post", "contenido": "contenido test post", "autor": "resty"`).
			Post("http://localhost:8080/api/v1/opinion")
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})

}
