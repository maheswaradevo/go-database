package database

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	insertData := "INSERT INTO customer(id, name) VALUES('devo', 'Devo')"
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

func TestQuerySqlComplex(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	queryData := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer"

	rows, err := db.QueryContext(ctx, queryData)
	if err != nil {
		panic(err)
	}

	for rows.Next(){
		var id, name, email string
		var balance int32
		var rating float64
		var birth_date, created_at time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}
		log.Println("=====================")
		log.Println("ID : ", id)
		log.Println("Name : ", name)
		log.Println("Email : ", email)
		log.Println("Balance : ", balance)
		log.Println("Rating : ", rating)
		log.Println("Created At : ", created_at)
		log.Println("Birth Date : ", birth_date)
		log.Println("Married : ", married)
	}
}