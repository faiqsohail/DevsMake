package persistence

import (
	"database/sql"
	"devsmake/persistence/interfaces"
	"devsmake/persistence/models"
)

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepo {
	return &AccountRepo{db}
}

var _ interfaces.AccountRepository = &AccountRepo{}

func (r *AccountRepo) GetUser(providerId uint64) (*models.User, error) {
	var user models.User
	err := r.db.
		QueryRow("SELECT id, provider_id, username, email, points, created FROM accounts where provider_id = ?", providerId).
		Scan(
			&user.ID,
			&user.ProviderID,
			&user.Username,
			&user.Email,
			&user.Points,
			&user.Created,
		)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AccountRepo) CreateUser(providerId uint64, username string) error {
	// Default current only supported provider
	provider := "github"

	_, err := r.db.
		Query("INSERT INTO accounts (provider, provider_id, username) VALUES (?, ?, ?)", provider, providerId, username)

	return err
}
