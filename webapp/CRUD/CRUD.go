package CRUD

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"time"
)

const (
	db_HOST	= "database"
	db_DATABASE = "mydb"
	db_USER	= "postgres"
	db_PASSWORD = "password"
	db_SSLMODE = "disable"
)

const db_TABLE = `(
  id	SERIAL PRIMARY KEY,
  name	text	NOT NULL,
  email	text	NOT NULL,
  created_at	timestamptz	NOT NULL,
  updated_at	timestamptz	NOT NULL
)`

type User struct {
	Id	int
  Name	string
  Email	string
  Created_at	time.Time
  Updated_at	time.Time
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func connectDB() (*sql.DB, error) {

	// Initialize connection string.
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", db_HOST, db_USER, db_PASSWORD, db_DATABASE, db_SSLMODE)

	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")

	sql_statement := fmt.Sprintf("SET TIMEZONE TO 'Japan';")
	_, err = db.Exec(sql_statement)
	checkError(err)
	fmt.Println("Finished setting timezone to Japan")

	return db,err
}

func WaitDB() {

	// Initialize connection string.
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", db_HOST, db_USER, db_PASSWORD, db_DATABASE, db_SSLMODE)

	// Initialize connection object.
	for{
		db, err := sql.Open("postgres", connectionString)
		err = db.Ping()
		if err == nil{
			break
		}
	}
	fmt.Println("Database is up")
}

func DropTable(tb_name string) {

	db,err := connectDB()
	defer db.Close()

	// Drop previous table of same name if one exists.
	sql_statement := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tb_name)
	_, err = db.Exec(sql_statement)
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")
}

func CreateTable(tb_name string) {

	db,err := connectDB()
	defer db.Close()

	// Create table.
	sql_statement := fmt.Sprintf("CREATE TABLE %s %s;", tb_name, db_TABLE)
	_, err = db.Exec(sql_statement)
	checkError(err)
	fmt.Println("Finished creating table")
}

func CreateData(name string, email string, tb_name string) (User){

	db,err := connectDB()
	defer db.Close()

	// Insert some data into table.
	sql_statement := fmt.Sprintf("INSERT INTO %s (name, email, created_at, updated_at) VALUES ($1, $2, current_timestamp, current_timestamp) RETURNING *;", tb_name)
	rows, err := db.Query(sql_statement, name, email)
	checkError(err)
	fmt.Println("Inserted 1 rows of data")

	var user User
	for rows.Next() {
		switch err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Created_at, &user.Updated_at); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %s, %d, %d)\n", user.Id, user.Name, user.Email, user.Created_at, user.Updated_at)
		default:
			checkError(err)
		}
	}
	return user

}

func ReadData(id int, tb_name string) (User) {

	db,err := connectDB()
	defer db.Close()

	// Read rows from table.

	sql_statement := fmt.Sprintf("SELECT * from %s WHERE id = $1 LIMIT 1;", tb_name)
	rows, err := db.Query(sql_statement, id)
	checkError(err)

	var user User
	for rows.Next() {
		switch err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Created_at, &user.Updated_at); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %s, %d, %d)\n", user.Id, user.Name, user.Email, user.Created_at, user.Updated_at)
		default:
			checkError(err)
		}
	}
	return user
}

//I wanna use overload
func ReadDataAll(tb_name string)  ([]User) {

	db,err := connectDB()
	defer db.Close()

  // Read rows from table.
	var users []User//= make([]User, 1)

	sql_statement := fmt.Sprintf("SELECT * from %s;", tb_name)
	rows, err := db.Query(sql_statement)
	checkError(err)

	for rows.Next() {
		var user User
		switch err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Created_at, &user.Updated_at); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %s, %d, %d)\n", user.Id, user.Name, user.Email, user.Created_at, user.Updated_at)
		default:
			checkError(err)
		}
		users = append(users, user)
	}
	return users
}

func UpdateData(id int, name string, email string, tb_name string) (User){

	db,err := connectDB()
	defer db.Close()

	// Modify some data in table.
	sql_statement := fmt.Sprintf("UPDATE %s SET name = $2, email = $3, updated_at = current_timestamp WHERE id = $1 RETURNING *;", tb_name)
	rows, err := db.Query(sql_statement, id, name, email)
	checkError(err)
	fmt.Println("Updated 1 row of data")

	var user User
	for rows.Next() {
		switch err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Created_at, &user.Updated_at); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %s, %d, %d)\n", user.Id, user.Name, user.Email, user.Created_at, user.Updated_at)
		default:
			checkError(err)
		}
	}
	return user

}

func DeleteData(id int, tb_name string) {

	db,err := connectDB()
	defer db.Close()

	// Delete some data from table.
	sql_statement := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", tb_name)
	_, err = db.Exec(sql_statement, id)
	checkError(err)
	fmt.Println("Deleted 1 row of data")
}
