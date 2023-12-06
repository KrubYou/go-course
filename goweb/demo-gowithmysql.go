package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func createTable(db *sql.DB) {
	query := `create table user( id int(11) NOT NULL AUTO_INCREMENT,
		username varchar(50) NOT NULL,
		password varchar(50) NOT NULL,
		create_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id));`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func insertData(db *sql.DB) {
	var username, password string
	fmt.Scan(&username, &password)
	createdAt := time.Now()

	result, err := db.Exec("insert into user(username,password,create_at) values(?,?,?)", username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	fmt.Println(id)

}

func query(db *sql.DB) {
	var (
		id         int
		coursename string
		price      float64
		instructor string
	)

	for {

		var inputID int
		fmt.Scan(&inputID)

		query := "select id,coursename,price,instructor from onlinecourse where id = ?"
		if err := db.QueryRow(query, inputID).Scan(&id, &coursename, &price, &instructor); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, coursename, price, instructor)
	}

}

func deleteData(db *sql.DB) {
	var deleteID int
	fmt.Scan(&deleteID)
	_, err := db.Exec("delete from user where id = ?", deleteID)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/coursedb")
	if err != nil {
		fmt.Println("Connect Error")
	} else {
		fmt.Println("Connect Success")
	}

	// createTable(db)
	//insertData(db)
	deleteData(db)
}
