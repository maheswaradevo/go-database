package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-database/database"
	"go-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "repository@gmail.com",
		Comment: "Test Repository",
	}

	res, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 41)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}