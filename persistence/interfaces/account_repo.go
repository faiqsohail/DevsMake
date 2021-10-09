package interfaces

import (
	"devsmake/persistence/models"
)

type AccountRepository interface {
	GetUser(uint64) (*models.User, error)
	CreateUser(uint64, string) error
}
