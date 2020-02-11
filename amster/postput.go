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

//InsertOperation is the Function to insert data to Database
func InsertOperation(uuid string, p People) (string, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return " ", nil
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
