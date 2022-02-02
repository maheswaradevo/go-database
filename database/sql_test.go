package database

import (
	"context"
	"fmt"
	"log"
	"strconv"
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

//TestQuerySqlComplex -> this function including more complex query using QueryContext
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

//TestQuerySQLParameter is using a parameter compared with using 
//an string concate
func TestQuerySQLParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "salah"

	queryData := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, queryData, username, password)

	if err != nil {
		panic(err)
	}
	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

// TestPrepareStatement including PrepareStatement that can be used when we want to insert or query
// multiple data at one time. So, the Exec or Query don't have to get data pool everytime.
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	insertData := "INSERT INTO comments(email, comments) VALUES(?, ?)"

	statement, err := db.PrepareContext(ctx, insertData)
	defer statement.Close()

	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++{
		email := "devo" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen ke" + strconv.Itoa(i)

		res, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses masukkan komen dengan ID : ", id)
	}

}