package configuracion

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)

const (
	DbUsername = "DbUsername"
	DbPassword = "DbPassword"
	DbHost     = "DbHost"
	DbSchema   = "DbSchema"
	DbPort     = "DbPort"
)

func GetDatabaseInstance() *sql.DB {
	username := os.Getenv(DbUsername)
	password := os.Getenv(DbPassword)
	host := os.Getenv(DbHost)
	schema := os.Getenv(DbSchema)
	port, _ := strconv.ParseInt(os.Getenv(DbPort), 10, 64)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, schema)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		_ = db.Close()
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
