package dbconfig

import (
	"database/sql"
	"fmt"
)

//Connect function for database connection
func Connect() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to database")
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

}
