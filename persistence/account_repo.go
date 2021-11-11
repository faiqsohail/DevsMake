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

	totalPosts, _ := r.sumPosts(int(user.ID), post)
	totalSubmissions, _ := r.sumPosts(int(user.ID), submission)
	totalComments, _ := r.sumPosts(int(user.ID), comment)

	user.TotalPosts = uint64(*totalPosts)
	user.TotalSubmissinos = uint64(*totalSubmissions)
	user.TotalComments = uint64(*totalComments)

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

		totalPosts, _ := r.sumPosts(int(user.ID), post)
		totalSubmissions, _ := r.sumPosts(int(user.ID), submission)
		totalComments, _ := r.sumPosts(int(user.ID), comment)

		user.TotalPosts = uint64(*totalPosts)
		user.TotalSubmissinos = uint64(*totalSubmissions)
		user.TotalComments = uint64(*totalComments)

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

type postType int64

const (
	post postType = iota
	submission
	comment
)

func (r *AccountRepo) sumPosts(accountId int, pt postType) (*int, error) {
	var count int

	var table string
	switch pt {
	case post:
		table = "posts"
	case submission:
		table = "posts_submissions"
	case comment:
		table = "posts_comments"
	default:
		table = "posts"
	}

	err := r.db.QueryRow("SELECT COUNT(*) FROM "+table+" WHERE author_id = ? AND deleted = 0", accountId).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}
