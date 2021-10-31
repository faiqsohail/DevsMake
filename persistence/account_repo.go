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

func (r *AccountRepo) GetUser(id uint64, providerId bool) (*models.User, error) {
	var user models.User

	var row *sql.Row
	if providerId {
		row = r.db.
			QueryRow("SELECT id, provider_id, username, email, points, created FROM accounts where provider_id = ?", id)
	} else {
		row = r.db.
			QueryRow("SELECT id, provider_id, username, email, points, created FROM accounts where id = ?", id)
	}

	err := row.Scan(
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

func (r *AccountRepo) GetUsers(limit uint64, offset uint64, sort string) (models.Users, error) {
	var users = models.Users{}

	results, err := r.db.
		Query("SELECT id, provider_id, username, email, points, created FROM accounts ORDER BY ? DESC LIMIT ?, ?", sort, offset, limit)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var user models.User

		results.Scan(
			&user.ID,
			&user.ProviderID,
			&user.Username,
			&user.Email,
			&user.Points,
			&user.Created,
		)

		users = append(users, user)
	}
	return users, nil
}

func (r *AccountRepo) CreateUser(providerId uint64, username string) error {
	// Default current only supported provider
	provider := "github"

	_, err := r.db.
		Query("INSERT INTO accounts (provider, provider_id, username) VALUES (?, ?, ?)", provider, providerId, username)

	return err
}
