package models

import (
	"database/sql"
)

type Tag struct {
	TagID   int
	TagName string
}

func GetTagByTagNameFromDB(db *sql.DB, tagName string) (*Tag, error) {
	query := "SELECT TagID, TagName FROM Tag WHERE TagName = ?"
	row := db.QueryRow(query, tagName)

	tag := &Tag{}
	err := row.Scan(&tag.TagID, &tag.TagName)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func GetAllTagsFromDB(db *sql.DB) ([]Tag, error) {
	query := "SELECT TagID, TagName FROM Tag"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag

	for rows.Next() {
		var tag Tag
		if err := rows.Scan(&tag.TagID, &tag.TagName); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
