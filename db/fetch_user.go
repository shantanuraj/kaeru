package db

import (
	"context"
	"sixth-io/kaeru/models"
)

// FetchUserForEmail fetches the user for given email
func (db *Database) FetchUserForEmail(email string) (user models.User, err error) {

	err = db.conn.QueryRow(
		context.Background(), `
		select id, name, email, password_hash from users
		where email=$1
		`,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)

	return
}

// FetchUserForID fetches the user for given id
func (db *Database) FetchUserForID(id string) (user models.User, err error) {

	err = db.conn.QueryRow(
		context.Background(), `
		select id, name, email, password_hash from users
		where id=$1
		`,
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)

	return
}
