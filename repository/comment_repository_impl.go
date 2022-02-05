package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

const(
	INSERT_DATA_COMMENT string = `
	INSERT INTO comments(email, comments) VALUES(?, ?)
	`

	FIND_DATA_BY_ID string = `
	SELECT id, email, comments FROM comments WHERE id = ? LIMIT 1
	`

	FIND_ALL_COMMENNT = `
	SELECT id, email, comments FROM comments
	`
)

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	result, err := repository.DB.ExecContext(ctx, INSERT_DATA_COMMENT, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	rows, err := repository.DB.QueryContext(ctx, FIND_DATA_BY_ID, id)
	comment := entity.Comment{}
	if err != nil{
		return comment, err
	}
	defer rows.Close()
	if rows.Next(){
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
	return comment, nil
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	rows, err := repository.DB.QueryContext(ctx, FIND_ALL_COMMENNT)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
