package database

import (
	"database/sql"
	"fmt"
	"schoolsystem/config"
	"schoolsystem/utils"

	_ "github.com/go-sql-driver/mysql"
)


func HandlingDB() {
	DBPassword, DBUsername, DBProtocol, DBAddress := utils.GetEnv()
	dbName := "tst"

	NewDB := config.MySQLDatabase{
		Username: DBUsername,
		Password: DBPassword,
		Protocol: DBProtocol,
		Address: DBAddress,
		DatabaseName: dbName,
	
	}

	CurrentDB := CreateDB(dbName, NewDB)
	CreateDBTable(CurrentDB.Database)
	
	//AddDBColumn(CurrentDB.Database)
	CurrentDB.Database.Close()
}

func CreateDB(dbName string, NewDB config.MySQLDatabase) *config.MySQLDatabase{
	fmt.Println("Creating DB in progress..")
	var err error

	DataSourceNameDB := fmt.Sprintf("%s:%s@%s(%s)/", NewDB.Username, NewDB.Password, NewDB.Protocol, NewDB.Address)

	db, err := sql.Open("mysql", DataSourceNameDB)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS "+dbName)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("USE "+dbName)
	if err != nil {
       panic(err)
   	}

	Database := new(config.MySQLDatabase)
	Database.Database = db
	return Database
}

func CreateDBTable(db *sql.DB) {
	var insert *sql.Stmt
	var err error

	// column name & type
	var tableName = "students"
	var columnData = "(id int, first_name varchar(50), last_name varchar(50), email varchar(40), current_semester int)"

	var tableCommand = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s %s", tableName, columnData)

	insert, err = db.Prepare(tableCommand)

   	if err != nil {
       panic(err)
   }
   defer insert.Close()

   _, err = insert.Exec()
   if err != nil {
	  panic(err)
   }

}

/*
func AddDBColumn(db *sql.DB) {
	var insert *sql.Stmt
	var err error

	var tableName = "students"
	var columnData = "(phone_number varchar(15))"

	var tableCommand = fmt.Sprintf("ALTER TABLE %s ADD %s", tableName, columnData)


	insert, err = db.Prepare(tableCommand)

   	if err != nil {
       panic(err)
   	}
   	defer insert.Close()

   	_, err = insert.Exec()
   	if err != nil {
	  panic(err)
   	}
}
*/

// func EditDBColumn() {

// }