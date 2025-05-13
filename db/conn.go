package db

import (
	"database/sql"
	"fmt"
	"go-api/env"
	"strconv"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	var (
		host     = env.PostgresDatabaseAddress.GetValue()
		port     = env.PostgresDatabasePort.GetValue()
		user     = env.PostgresUser.GetValue()
		password = env.PostgresPassword.GetValue()
		dbname   = env.PostgresDefaultDatabase.GetValue()
	)
	fmt.Println(port, host)

	intPort, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, intPort, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	return db, nil
}
