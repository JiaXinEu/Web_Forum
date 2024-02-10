// models/thread.go
package models

import (
	"database/sql"
	"time"
)

type Thread struct {
	ThreadID    int       `json:"thread_id"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
	UserID      int       `json:"user_id"`
	UpvoteCount int       `json:"upvote_count"`
}

func GetThreadsByUserIDFromDB(db *sql.DB, userID int) ([]Thread, error) {
	query := "SELECT ThreadID, Content, Timestamp, UserID, UpvoteCount FROM Thread WHERE UserID = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []Thread

	for rows.Next() {
		var thread Thread
		if err := rows.Scan(&thread.ThreadID, &thread.Content, &thread.Timestamp, &thread.UserID, &thread.UpvoteCount); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}

	return threads, nil
}

func GetAllThreadsFromDB(db *sql.DB) ([]Thread, error) {
	query := "SELECT ThreadID, Content, Timestamp, UserID, UpvoteCount FROM Thread"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []Thread

	for rows.Next() {
		var thread Thread
		if err := rows.Scan(&thread.ThreadID, &thread.Content, &thread.Timestamp, &thread.UserID, &thread.UpvoteCount); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}

	return threads, nil
}

func GetThreadsByTagIDFromDB(db *sql.DB, tagID int) ([]Thread, error) {
	query := "SELECT ThreadID, Content, Timestamp, UserID, UpvoteCount FROM Thread WHERE ThreadID IN (SELECT ThreadID FROM ThreadTag WHERE TagID = ?)"
	rows, err := db.Query(query, tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []Thread

	for rows.Next() {
		var thread Thread
		if err := rows.Scan(&thread.ThreadID, &thread.Content, &thread.Timestamp, &thread.UserID, &thread.UpvoteCount); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}

	return threads, nil
}

func InsertThreadIntoDB(db *sql.DB, thread *Thread) (int, error) {
	query := "INSERT INTO Thread (Content, UserID, UpvoteCount) VALUES (?, ?, 0)"
	result, err := db.Exec(query, thread.Content, thread.UserID)
	if err != nil {
		return 0, err
	}

	threadID, _ := result.LastInsertId()
	return int(threadID), nil
}

func IncrementDbUpvoteCount(db *sql.DB, threadID int) error {
	query := "UPDATE Thread SET UpvoteCount = UpvoteCount + 1 WHERE ThreadID = ?"
	_, err := db.Exec(query, threadID)
	return err
}

func DecrementDbUpvoteCount(db *sql.DB, threadID int) error {
	query := "UPDATE Thread SET UpvoteCount = UpvoteCount - 1 WHERE ThreadID = ?"
	_, err := db.Exec(query, threadID)
	return err
}

func EditContentInDb(db *sql.DB, threadID int, content string) error {
	query := "UPDATE Thread SET Content = ? WHERE ThreadID = ?"
	_, err := db.Exec(query, content, threadID)
	return err
}
