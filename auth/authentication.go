package auth

import (
	"errors"
	"fmt"
	db2 "github.com/Swan/Nameless/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetInGameUser Retrieves an online user by their in-game session in Redis.
func GetInGameUser(c *gin.Context) (db2.User, error) {
	authHeader, err := getAuthenticationToken(c)

	if err != nil {
		return db2.User{}, err
	}

	err = verifyUserAgent(c)

	if err != nil {
		return db2.User{}, err
	}

	userId, err := getUserIdFromToken(authHeader)

	if err != nil {
		return db2.User{}, err
	}

	user, err := db2.GetUserById(userId)

	if err != nil {
		return db2.User{}, err
	}

	if !user.Allowed {
		return db2.User{}, err
	}

	return user, nil
}

// Fetches the authentication token from a request's headers if one exists
func getAuthenticationToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("auth")

	if authHeader == "" {
		return "", errors.New("request does not contain an `auth` token")
	}

	return authHeader, nil
}

// Checks if the request has a valid user agent User-Agent must be "Quaver"
func verifyUserAgent(c *gin.Context) error {
	if c.GetHeader("User-Agent") != "Quaver" {
		return errors.New("failed to authenticate client details")
	}

	return nil
}

// Fetches the user's id from redis with their authentication token
func getUserIdFromToken(token string) (int, error) {
	key := fmt.Sprintf("quaver:server:session:%v", token)
	userId, err := db2.Redis.Get(db2.RedisCtx, key).Result()

	if err != nil {
		return -1, err
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		return -1, err
	}

	return id, nil
}
