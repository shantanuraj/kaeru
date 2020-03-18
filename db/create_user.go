package db

import (
	"context"
	"sixth-io/kaeru/hash"
	"sixth-io/kaeru/models"

	"github.com/google/uuid"
)

// CreateUser creates the user
func (db *Database) CreateUser(user models.User) (err error) {
	id := uuid.New().String()
	hash, salt, err := hash.Password(user.Password)

	if err != nil {
		return err
	}

	_, err = db.conn.Exec(context.Background(), `
	insert into users(id, name, email, password_hash, salt)
	values ($1, $2, $3, $4, $5)
	`,
		id, user.Name, user.Email, hash, salt)

	return err
}