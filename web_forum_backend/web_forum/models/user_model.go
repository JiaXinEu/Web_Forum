package models

import (
	"database/sql"
)

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

func GetUserByUsernameFromDB(db *sql.DB, username string) (*User, error) {
	query := "SELECT UserID, Username FROM User WHERE Username = ?"
	row := db.QueryRow(query, username)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func InsertUserIntoDB(db *sql.DB, user *User) (int, error) {
	query := "INSERT INTO User (Username) VALUES (?)"
	result, err := db.Exec(query, user.Username)
	if err != nil {
		return 0, err
	}

	userID, _ := result.LastInsertId()
	return int(userID), nil
}
