package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

const (
	HOST	= "localhost"
	DATABASE= "mydb"
	USER	= "postgres"
	PASSWORD= "password"
	SSLMODE = "disable"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func connectDB() (*sql.DB, error) {

	// Initialize connection string.
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", HOST, USER, PASSWORD, DATABASE, SSLMODE)

	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")
	return db,err
}

func dropTable(tb_name string) {

	db,err := connectDB()
	defer db.Close()

	// Drop previous table of same name if one exists.
	sql_statement := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tb_name)
	_, err = db.Exec(sql_statement)
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")
}

func createTable(tb_name string) {

	db,err := connectDB()
	defer db.Close()

	// Create table.
	sql_statement := fmt.Sprintf("CREATE TABLE %s (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);", tb_name)
	_, err = db.Exec(sql_statement)
	checkError(err)
	fmt.Println("Finished creating table")
}

func createData(name string, value int, tb_name string) {

	db,err := connectDB()
	defer db.Close()

	// Insert some data into table.
	sql_statement := fmt.Sprintf("INSERT INTO %s (name, quantity) VALUES ($1, $2);", tb_name)
	_, err = db.Exec(sql_statement, name, value)
	checkError(err)
	fmt.Println("Inserted 1 rows of data")
}

func readData(tb_name string) {

	db,err := connectDB()
	defer db.Close()

    	// Read rows from table.
	var id int
	var name string
	var quantity int

	sql_statement := fmt.Sprintf("SELECT * from %s;", tb_name)
	rows, err := db.Query(sql_statement)
	checkError(err)

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &quantity); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
		default:
			checkError(err)
		}
	}
}

func updateData(name string, value int, tb_name string) {

	db,err := connectDB()
	defer db.Close()

	// Modify some data in table.
	sql_statement := fmt.Sprintf("UPDATE %s SET quantity = $2 WHERE name = $1;", tb_name)
	_, err = db.Exec(sql_statement, name, value)
	checkError(err)
	fmt.Println("Updated 1 row of data")
}

func deleteData(name string, tb_name string) {

	db,err := connectDB()
	defer db.Close()

	// Delete some data from table.
	sql_statement := fmt.Sprintf("DELETE FROM %s WHERE name = $1;", tb_name)
	_, err = db.Exec(sql_statement, name)
	checkError(err)
	fmt.Println("Deleted 1 row of data")
}

func main(){
	tb_name := "test_table"
	dropTable(tb_name)
	createTable(tb_name)
	createData("banana", 25, tb_name)
	createData("orange", 32, tb_name)
	createData("apple", 64, tb_name)
	readData(tb_name)
	updateData("banaba", 60, tb_name)
	readData(tb_name)
	deleteData("orange", tb_name)
	readData(tb_name)
}
