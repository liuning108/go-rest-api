package db

import (
	"context"
	"database/sql"
	"go-rest-api/internal/comment"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {

	var cmtRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		"Select id,slug,body,author from comments where id = $1",
		uuid,
	)
	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, err
	}

	return convertCommentRowToComment(cmtRow), nil
}

func convertCommentRowToComment(cmtRow CommentRow) comment.Comment {
	return comment.Comment{
		ID:     cmtRow.ID,
		Slug:   cmtRow.Slug.String,
		Body:   cmtRow.Body.String,
		Author: cmtRow.Author.String,
	}
}
