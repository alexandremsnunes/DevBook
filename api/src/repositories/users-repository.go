package repositories

import (
	"api/src/models"
	"database/sql"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (repo UsersRepository) Create(user models.User) (uint64, error) {
	statement, erro := repo.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastInsertID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertID), nil

}
