package configuracion

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
)

const (
	DbUsername    = "DbUsername"
	DbPassword    = "DbPassword"
	DbHost        = "DbHost"
	DbSchema      = "DbSchema"
	DbSchemaTest  = "DbSchemaTest"
	DbPort        = "DbPort"
	DbEnvironment = "DbEnvironment"
)

var (
	db  *sql.DB
	err error
)

func GetDatabaseInstance() *sql.DB {
	environment := os.Getenv(DbEnvironment)
	username := os.Getenv(DbUsername)
	password := os.Getenv(DbPassword)
	host := os.Getenv(DbHost)
	var schema string
	if environment == "test" {
		schema = os.Getenv(DbSchemaTest)
	} else {
		schema = os.Getenv(DbSchema)
	}
	port, _ := strconv.ParseInt(os.Getenv(DbPort), 10, 64)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, schema)
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		_ = db.Close()
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//if environment == "test" {
	//	_ = refreshPeliculaTable()
	//}
	return db
}

func CloseConnections(err error, tx *sql.Tx, stmt *sql.Stmt, rows *sql.Rows) {
	if tx != nil {
		switch err {
		case nil:
			_ = tx.Commit()
		default:
			_ = tx.Rollback()
		}
	}

	if stmt != nil {
		_ = stmt.Close()
	}

	if rows != nil {
		_ = rows.Close()
	}
}

func refreshPeliculaTable() error {
	stmt, errr := db.Prepare("TRUNCATE TABLE pelicula")
	if errr != nil {
		panic(errr.Error())
	}
	_, errr = stmt.Exec()
	if errr != nil {
		log.Fatalf("Error truncating Peliculas table: %s", errr)
	}
	return nil
}
