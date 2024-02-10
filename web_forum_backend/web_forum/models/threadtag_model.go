// models/thread_tag.go
package models

import "database/sql"

type ThreadTag struct {
	ThreadID int
	TagID    int
}

func InsertThreadTagIntoDB(db *sql.DB, threadTag *ThreadTag) (error) {
	query := "INSERT INTO ThreadTag (ThreadID, TagID) VALUES (?, ?)"
	_, err := db.Exec(query, threadTag.ThreadID, threadTag.TagID)
	if err != nil {
		return err
	}

	return nil
}
