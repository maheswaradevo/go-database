package database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	insertData := "INSERT INTO customer(id, name) VALUES('reza', 'Reza')"
	_, err := db.ExecContext(ctx, insertData)
	if err != nil{
		panic(err)
	}
	fmt.Println("Success insert data customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	queryData := "SELECT id, name FROM customer"
	
	rows, err := db.QueryContext(ctx, queryData)
	if err != nil {
		panic(err)
	}

	for rows.Next(){
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
	}
	defer rows.Close()
}