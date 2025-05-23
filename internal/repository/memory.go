package repository

import (
	"fmt"
	"sync"

	"github.com/samezzz/hyperchat/internal/models"
)

var (
	userStates = make(map[string]*models.UserState)
	mutex      = &sync.Mutex{}
)

// GetUserState retrieves the user state by phone number
func GetUserState(user string) (*models.UserState, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	state, exists := userStates[user]
	return state, exists
}

// SaveUserState stores or updates the user state
func SaveUserState(user string, state *models.UserState) {
	mutex.Lock()
	defer mutex.Unlock()
	userStates[user] = state
	fmt.Printf(user)
}
