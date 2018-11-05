// TO-DO: For now db transactions will be handled using 'database/sql' package directly,
// 		  an ORM implementation will be replaced with current
// 		  implementation => http://gorm.io/docs/ (this the de facto standard on ORM in golang as I understand.)
package db_manager


import (
	"fmt"
	"database/sql"
	 _ "github.com/ahmeturun/pq"
)

type ConnectionStruct struct{
	Db *sql.DB
	err error
}
var Connection ConnectionStruct;
func OpenConnection(){
	connStr := "postgres://klvlelkf:DzDY7DptCu4IF4mkW1ieczwPUH23aRRw@stampy.db.elephantsql.com:5432/klvlelkf"
	db, err := sql.Open("postgres", connStr)
	Connection.Db = db
	Connection.err = err
	if err != nil {
		fmt.Println(err)
	}
}

func CloseConnection(){
	Connection.Db.Close()
}

func RunQuery(queryString string)(*sql.Rows, error){
	OpenConnection()
	return Connection.Db.Query(queryString)
	CloseConnection()
}

