package db

import (
	"context"
	"sixth-io/kaeru/models"

	"github.com/google/uuid"
)

// CreateUser creates the user
func (db *Database) CreateUser(user models.User) (id string, err error) {
	id = uuid.New().String()

	_, err = db.conn.Exec(
		context.Background(), `
		insert into users(id, name, email, password_hash)
		values ($1, $2, $3, $4)
		`,
		id, user.Name, user.Email, user.PasswordHash,
	)

	return
}
