package handlers

import (
	"github.com/google/uuid"
)

var (
	cookieTime = 604800 // 3 days
)

func generateSessionToken() string {
	return uuid.NewString()
}
