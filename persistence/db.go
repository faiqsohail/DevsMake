package persistence

import (
	"database/sql"
	"devsmake/persistence/interfaces"
	"os"

	// Import MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

type Repositories struct {
	AccountRepo interfaces.AccountRepository
	PostRepos   interfaces.PostRepositories
	db          *sql.DB
}

func NewRepository() (*Repositories, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":3306)/"+os.Getenv("DB_NAME"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Repositories{
		AccountRepo: NewAccountRepository(db),
		PostRepos:   NewPostRepositories(db),
		db:          db,
	}, nil
}

func (r *Repositories) Ping() error {
	return r.db.Ping()
}

func (r *Repositories) Close() error {
	return r.db.Close()
}
