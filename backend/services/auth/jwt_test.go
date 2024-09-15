package auth

import (
	"testing"

	"github.com/golang-jwt/jwt"
)

func TestCreateJWT(t *testing.T) {
	t.Run("CreateAndValidateCorrectJWT", func(t *testing.T) {
		secret := "secret"
		userId := 123
		jwtToken, err := CreateJWT(userId, []byte(secret))
		if err != nil {
			t.Errorf("CreateJWT() failed: %v", err)
		}

		// Validate the JWT
		token, err := validateJWT(jwtToken, []byte(secret))
		if err != nil {
			t.Errorf("ValidateJWT() failed: %v", err)
		}

		// Check the claims
		if !token.Valid {
			t.Errorf("ValidateJWT() failed: claims are nil")
		}

		// Check the userID
		claims := token.Claims.(jwt.MapClaims)
		userID, ok := claims["userID"].(string)
		if !ok {
			t.Errorf("ValidateJWT() failed: userID is missing")
		}

		if userID != "123" {
			t.Errorf("ValidateJWT() failed: invalid userID")
		}
	})

	t.Run("CreateAndValidateIncorrectSecret", func(t *testing.T) {
		secret := "secret"
		userId := 123
		jwt, err := CreateJWT(userId, []byte(secret))
		if err != nil {
			t.Errorf("CreateJWT() failed: %v", err)
		}

		// Validate the JWT
		_, err = validateJWT(jwt, []byte("wrong_secret"))
		if err == nil {
			t.Errorf("ValidateJWT() failed: expected an error")
		}
	})

	t.Run("CreateAndValidateIncorrectJWT", func(t *testing.T) {
		secret := "secret"
		userId := 123
		_, err := CreateJWT(userId, []byte(secret))
		if err != nil {
			t.Errorf("CreateJWT() failed: %v", err)
		}

		// Validate the JWT
		_, err = validateJWT("wrong_jwt", []byte(secret))
		if err == nil {
			t.Errorf("ValidateJWT() failed: expected an error")
		}
	})

}
