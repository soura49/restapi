package amster

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// Using the pq as the postgres driver
	_ "github.com/lib/pq"
)

// DB configuration const
const (
	Host     = "localhost"
	Port     = 5432
	User     = "container"
	Password = "test"
	Dbname   = "titanic"
)

func dbconnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

//InsertOperation is the Function to insert data to Database
func InsertOperation(uuid string, p People) (string, error) {
	db, err := dbconnection()
	if err != nil {
		return " ", err
	}
	defer db.Close()
	sqlstmt := `
	INSERT INTO people (uuid,info)
	VALUES ($1, $2)`
	fmt.Println(uuid, p)
	o, err := json.Marshal(p)
	if err != nil {
		return " ", err
	}
	_, err = db.Exec(sqlstmt, uuid, o)
	if err != nil {
		return " ", err
	}
	return "Successfully Inserted", nil
}

//SelectOperationByID reperesents the function to retrieve the data from the database based on the given uuid input
func SelectOperationByID(uuid string) (*sql.Row, error) {
	db, err := dbconnection()
	if err != nil {
		return nil, err
	}
	sqlstmt := `select info from people where uuid=$1`
	defer db.Close()
	out := db.QueryRow(sqlstmt, uuid)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//SelectOperationAll reperesents the function to retrieve all the data from the database
func SelectOperationAll() (*sql.Rows, error) {
	db, err := dbconnection()
	if err != nil {
		return nil, err
	}
	sqlstmt := `
	select * from people
	`
	defer db.Close()
	out, err := db.Query(sqlstmt)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteOperationByID represents deleting the row in the DB by the UUID
func DeleteOperationByID(uuid string) (int64, error) {
	db, err := dbconnection()
	if err != nil {
		return 0, err
	}
	sqlstmt := `
	delete from people where uuid=$1
	`
	defer db.Close()
	out, err := db.Exec(sqlstmt, uuid)
	if err != nil {
		return 0, err
	}
	numDeleted, err := out.RowsAffected()
	if err != nil {
		return 0, err
	}
	return numDeleted, nil
}
