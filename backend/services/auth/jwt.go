package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jonathanmeij/go-reservation/configs"
	"github.com/jonathanmeij/go-reservation/types"
	"github.com/jonathanmeij/go-reservation/utils"
)

type contextKey string

const UserKey contextKey = "userID"

func jwtAuthHelper(w http.ResponseWriter, r *http.Request, store types.UserStore) (*types.UserEntity, bool) {
	tokenString := utils.GetTokenFromRequest(r)

	token, err := validateJWT(tokenString, []byte(configs.Envs.JWTSecret))
	if err != nil {
		log.Printf("failed to validate token: %v", err)
		permissionDenied(w)
		return nil, false
	}

	if !token.Valid {
		log.Println("invalid token")
		permissionDenied(w)
		return nil, false
	}

	claims := token.Claims.(jwt.MapClaims)
	str := claims["userID"].(string)

	userID, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("failed to convert userID to int: %v", err)
		permissionDenied(w)
		return nil, false
	}

	u, err := store.GetUserByID(userID)
	if err != nil {
		log.Printf("failed to get user by id: %v", err)
		permissionDenied(w)
		return nil, false
	}

	return u, true
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, ok := jwtAuthHelper(w, r, store)
		if !ok {
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func WithJWTAuthRole(handlerFunc http.HandlerFunc, store types.UserStore, role string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, ok := jwtAuthHelper(w, r, store)
		if !ok {
			return
		}

		if u.Role != role {
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func CreateJWT(userID uint, secret []byte) (string, error) {
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func validateJWT(tokenString string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}
