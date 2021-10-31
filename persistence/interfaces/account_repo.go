package interfaces

import (
	"devsmake/persistence/models"
)

type AccountRepository interface {
	GetUser(uint64, bool) (*models.User, error)
	GetUsers(uint64, uint64, string) (models.Users, error)
	CreateUser(uint64, string) error
}
