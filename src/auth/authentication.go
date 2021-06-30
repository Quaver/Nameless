package auth

import (
	"errors"
	"fmt"
	"github.com/Swan/Nameless/src/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetInGameUser Retrieves an online user by their in-game session in Redis.
func GetInGameUser(c *gin.Context) (db.User, error) {
	authHeader, err := getAuthenticationToken(c)

	if err != nil {
		return db.User{}, err
	}

	err = verifyUserAgent(c)

	if err != nil {
		return db.User{}, err
	}

	userId, err := getUserIdFromToken(authHeader)

	if err != nil {
		return db.User{}, err
	}

	user, err := db.GetUserById(userId)

	if err != nil {
		return db.User{}, err
	}

	if !user.Allowed {
		return db.User{}, err
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
	userId, err := db.Redis.Get(db.RedisCtx, key).Result()

	if err != nil {
		return -1, errors.New("an online session with that token could not be found")
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		return -1, errors.New("failed to parse user id from online session")
	}

	return id, nil
}
