package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Customer type
type Customer struct {
	Id   int
	Name string
}

func main() {

	// inserting a rows
	insert(Customer{101, "Neeraj"})
	insert(Customer{102, "Neeraj"})
	insert(Customer{103, "Rajat"})
	insert(Customer{104, "Suraj"})

	// updating the customer by id
	// updateById(Customer{101, "Singh"})
	updateById(Customer{104, "Kumar"})

	// delete a customer by id
	delete(103)

}

// function to get a database connection
func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@/go_db")
	if err != nil {
		fmt.Println("Error! Getting connection...")
	}
	return db
}

// function to insert a row in customer table
func insert(customer Customer) {
	db := connect()
	insert, err := db.Query("INSERT INTO customer VALUES (?, ?)", customer.Id, customer.Name)
	if err != nil {
		fmt.Println("Error! Inserting records...")
	}
	defer insert.Close()
	defer db.Close()
}

// function to update a customer record by customer id
func updateById(customer Customer) {
	db := connect()
	db.QueryRow("UPDATE customer SET name=? WHERE id=?", customer.Name, customer.Id)
}

// function to delete a customer by customer id
func delete(id int) {
	db := connect()
	db.QueryRow("DELETE FROM customer WHERE id=?", id)
}
