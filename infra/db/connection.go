package db

import (
	"fmt"
	"practice/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DbConfig) string {

	conStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", 
	cnf.User, cnf.Password, cnf.Host, cnf.Port, cnf.Name)

	if !cnf.SslEnableMode {
		conStr += " sslmode=disable"
	}
	return conStr
}

func NewConnection(cnf *config.DbConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
