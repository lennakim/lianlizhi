package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitDB() *sqlx.DB {
	return openDB(viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
	)
}

func openDB(host string, port int, user string, password string, dbname string) *sqlx.DB {
	config := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Printf(config, "\n")

	db, err := sqlx.Connect("postgres", config)

	if err != nil {
		panic(err)
	}

	setupDB(db)

	return db
}

func setupDB(db *sqlx.DB) {
}
