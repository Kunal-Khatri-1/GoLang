package main

/*
-------------------------GO-----------------------------
 1. go get -u https://github.com/go-sql-driver/mysql
	- install drivers to connect with MySQL

------------------------MySQL---------------------------
 1. mysql --version
 2. mysql -u root -p
 3. create database goassignment;
 4. use goassignment;
 5. create table posts(
	id int,
	name varchar(255),
	text varchar(255),
	primary key (id),
 );
 6. desc posts;
*/
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id int
	Name string
	Text string
}


func main() {
	fmt.Println("Hello World")
	// connecting with MySQL
	connectDB()
}

func connectDB() {
	// connecting with MySQL
	db, err := sql.Open("mysql", "root:incorrectpassword752@/goassignment")
	// checking for any errors
	checkNilError(err)

	// defer => will run at the end of the function execution (before return)
	// closing the connection in the end
	defer db.Close()

	// After trying to connect using sql.Open, checking if connection is still alive
	// panic if there is any error
	PingDB(db)

	// executing insert query Create ✅
	insertVals(db,"6", "Post siz", "Some random text")

	// executing update query Update ✅
	updateVals(db, "4", "Some another random text")

	// exectuing fetching query Read ✅
	fetchVals(db)

	// executing delete query Delete ✅
	deleteRow(db, "1")
}

// db is of type sql.DB
// * => it should be reference and not a copy of db
func PingDB(db *sql.DB) {
	// verfiyig the connection with database is active or not
	// establishes a connection if necessary
	err := db.Ping()
	checkNilError(err)
}

func checkNilError(err error) {
	// if err is not empty => there is some error
	if err != nil {
		panic(err)
	}
}

// inserting table contents
func insertVals(db *sql.DB, id string, name string, text string) {
	// preparing to insert values
	stmt, err := db.Prepare("insert into posts(id, name, text) values(?, ?, ?)")
	checkNilError(err)
	
	// execute the query
	res, err := stmt.Exec(id, name, text)
	checkNilError(err)

	// getting the insert Id of the command
	lastId, err := res.LastInsertId()
	checkNilError(err)

	fmt.Println("Last insert id", lastId)
}

// updating table contents
func updateVals(db *sql.DB, id string, txt string) {
	// preparing to update values
	stmt, err := db.Prepare("update posts set text=? where id=?")
	checkNilError(err)
	
	// executing update query
	res, err := stmt.Exec(txt, id)
	checkNilError(err)

	// getting number of rows affected
	// if rowsAffected > 1 => table was updated
	rowsAffected, err := res.RowsAffected()
	checkNilError(err)

	fmt.Println("Number of rows affected: ", rowsAffected)
}

func fetchVals(db *sql.DB) {
	// fetching all the table contents
	rows, err := db.Query("select * from posts")
	checkNilError(err)

	// empty post variable of type Post
	var post = Post{}

	// looping over the table rows to get table contents
	// rows.Next() will prepare next row of the table to be read by scan and returns true if there is any row in the table
	for rows.Next() {
		// rows.Scan copies the columns of the corresponding rows and put them into corresponing destinations like &post.Id
		// return error in case of any error
		err = rows.Scan(&post.Id, &post.Name, &post.Text)
		checkNilError(err)

		fmt.Println("Contents of the table: \n",post)
	}
}

// Deleting Row
func deleteRow(db *sql.DB, id string) {
	// preparing to delete values
	stmt, err := db.Prepare("delete from posts where id=?")
	checkNilError(err)

	res, err := stmt.Exec(id)
	checkNilError(err)

	// rowsAffected == 1 => row was deleted
	rowsAffected, err := res.RowsAffected()
	checkNilError(err)

	fmt.Println("Number of rows affected: ", rowsAffected)
}