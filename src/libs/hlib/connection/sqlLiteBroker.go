package connection

import (
	"database/sql"
	"fmt"
	_ "libs/github.com/mattn/go-sqlite3"
)

type SqlLiteBroker struct{}

/**
     * Método para realizar una instancia de conexión con sqlLite.
     *
     * @return    conexión SqlLite.
     * @return    error, en alguna parte del proceso(error).
**/

func (s *SqlLiteBroker) ConnectSqlLite() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./bdStandar/securityLite.db")

	if err != nil {
		fmt.Println("error ", err)
		return nil, err
	}

	return db, nil
}
