package config

import "database/sql"

type MySQLDatabase struct {
	Username     string `json:"db_username"`
	Password     string `json:"db_password"`
	Protocol     string `json:"db_protocol"`
	Address      string `json:"db_address"`
	DatabaseName string `json:"db_name"`
	Database     *sql.DB
}