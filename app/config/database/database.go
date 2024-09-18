package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "rsaude-db-develop.radarsaude.net.br"
	port     = 5432
	user     = "11090534957"
	password = "2lQ4P2lg8kQGZ0xNpZVsxo5ObnWBXln9a1yeRrSo9E1Rdp2X6M1nGl5ihjx2zdq4"
	dbname   = "cnes"
)

func ConnectDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
