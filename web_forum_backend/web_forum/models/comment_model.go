// models/comment.go
package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	CommentID int
	Content   string
	Timestamp time.Time
	UserID    int
	ThreadID  int
}

func GetCommentsByThreadIDFromDB(db *sql.DB, threadID int) ([]Comment, error) {
	query := "SELECT CommentID, Content, Timestamp, UserID, ThreadID FROM Comments WHERE UserID = ?"
	rows, err := db.Query(query, threadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.CommentID, &comment.Content, &comment.Timestamp, &comment.UserID, &comment.ThreadID); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func InsertCommentIntoDB(db *sql.DB, comment *Comment) (int, error) {
    query := "INSERT INTO Comment (Content, UserID, ThreadID) VALUES (?, ?, 0)"
    result, err := db.Exec(query, comment.Content, comment.UserID, comment.ThreadID)
    if err != nil {
        return 0, err
    }

    commentID, _ := result.LastInsertId()
    return int(commentID), nil
}