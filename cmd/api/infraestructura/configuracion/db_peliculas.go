package configuracion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
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

func RefreshPeliculaTable() error {
	stmt, errr := db.Prepare("TRUNCATE TABLE pelicula")
	if errr != nil {
		panic(errr.Error())
	}
	_, errr = stmt.Exec()
	if errr != nil {
		log.Fatalf("Error truncating Peliculas table: %s", errr)
	}
	return errr
}

func SendOnePelicula() (modelo.Pelicula, error) {
	pelicula := modelo.Pelicula{
		Nombre:      "Nombre test",
		Director:    "Director test",
		Escritor:    "Escritor test",
		Pais:        "Pais test",
		Idioma:      "Idioma test",
		Lanzamiento: 2021,
	}
	stmt, errr := db.Prepare("INSERT INTO pelicula (nombre, director, escritor, pais, idioma, lanzamiento) VALUES (?, ?, ?, ?, ?, ?)")
	if errr != nil {
		panic(errr.Error())
	}
	insertResult, createErr := stmt.Exec(
		pelicula.Nombre,
		pelicula.Director,
		pelicula.Escritor,
		pelicula.Pais,
		pelicula.Idioma,
		pelicula.Lanzamiento)
	if createErr != nil {
		log.Fatalf("Error creating message: %s", createErr)
	}
	msgId, errr := insertResult.LastInsertId()
	if errr != nil {
		log.Fatalf("Error creating message: %s", createErr)
	}
	pelicula.Id = msgId
	return pelicula, nil
}

func SendVariousPeliculas() ([]modelo.Pelicula, error) {
	peliculas := []modelo.Pelicula{
		{
			Nombre:      "Nombre 1",
			Director:    "Director 1",
			Escritor:    "Escritor 1",
			Pais:        "Pais 1",
			Idioma:      "Idioma 1",
			Lanzamiento: 2021,
		},
		{
			Nombre:      "Nombre 2",
			Director:    "Director 2",
			Escritor:    "Escritor 2",
			Pais:        "Pais 2",
			Idioma:      "Idioma 2",
			Lanzamiento: 2022,
		},
	}
	stmt, errr := db.Prepare("INSERT INTO pelicula (nombre, director, escritor, pais, idioma, lanzamiento) VALUES (?, ?, ?, ?, ?, ?)")
	if errr != nil {
		panic(errr.Error())
	}
	for i := range peliculas {
		_, createErr := stmt.Exec(
			peliculas[i].Nombre,
			peliculas[i].Director,
			peliculas[i].Escritor,
			peliculas[i].Pais,
			peliculas[i].Idioma,
			peliculas[i].Lanzamiento)
		if createErr != nil {
			return nil, createErr
		}
	}
	get_stmt, errr := db.Prepare("SELECT id, nombre, director, escritor, pais, idioma, lanzamiento FROM pelicula")
	if errr != nil {
		return nil, errr
	}
	defer stmt.Close()

	rows, errr := get_stmt.Query()
	if errr != nil {
		return nil, errr
	}
	defer rows.Close()

	results := make([]modelo.Pelicula, 0)

	for rows.Next() {
		var pelicula modelo.Pelicula
		if getError := rows.Scan(
			&pelicula.Id,
			&pelicula.Nombre,
			&pelicula.Director,
			&pelicula.Escritor,
			&pelicula.Pais,
			&pelicula.Idioma,
			&pelicula.Lanzamiento); getError != nil {
			return nil, errr
		}
		results = append(results, pelicula)
	}
	return results, nil
}
