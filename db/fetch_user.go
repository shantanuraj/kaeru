package db

import (
	"context"
	"sixth-io/kaeru/models"
)

// FetchUserForEmail fetches the user for given email
func (db *Database) FetchUserForEmail(email string) (user models.User, err error) {

	err = db.conn.QueryRow(
		context.Background(), `
		select id, name, email, password_hash, salt from users
		where email=$1
		`,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Salt)

	return
}
