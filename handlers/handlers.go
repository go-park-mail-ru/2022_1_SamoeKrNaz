package handlers

import (
	"github.com/google/uuid"
	"sync"
)

var (
	cookieTime = 604800 // 3 days
	lockUser   = sync.RWMutex{}
	lockSess   = sync.RWMutex{}
)

func generateSessionToken() string {
	return uuid.NewString()
}
