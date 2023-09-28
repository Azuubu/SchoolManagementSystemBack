package database

import (
	"database/sql"
	"fmt"
	"schoolsystem/types"
	"schoolsystem/utils"

	_ "github.com/go-sql-driver/mysql"
)

type Storage interface {
	CreateUser(*types.User) error
	UpdateUser(*types.User) error
	GetUserByID(int) (*types.User, error)
	GetUsers() ([]*types.User, error)  // ? idk
}

type MySQLDatabase struct {
	Db    *sql.DB
}

func NewMySQLDB() (*MySQLDatabase, error) {
	DBPassword, DBUsername, DBProtocol, DBAddress := utils.GetEnv()
	dbName := "testdb1"

	DataSourceNameDB := fmt.Sprintf("%s:%s@%s(%s)/", DBUsername, DBPassword, DBProtocol, DBAddress)

	db, err := sql.Open("mysql", DataSourceNameDB)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS "+dbName)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("USE "+dbName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLDatabase{
		Db: db,
	}, nil
}

func (s *MySQLDatabase) Init() error {
	return s.createUserTable()
}

func (s *MySQLDatabase) createUserTable() error {
	query := "CREATE TABLE IF NOT EXISTS students (id int, first_name varchar(50), last_name varchar(50), email varchar(40), date_of_birth date, gpa decimal(2, 1))"
	_,err := s.Db.Exec(query)
	
   return err
}

func (s *MySQLDatabase) CreateUser(*types.User) error {
	return nil
}

func (s *MySQLDatabase) UpdateUser(*types.User) error {
	return nil

}

func (s *MySQLDatabase) GetUsers() ([]*types.User, error) {
	return nil, nil

}

func (s *MySQLDatabase) GetUserByID(id int) (*types.User, error) {
	return nil, nil

}