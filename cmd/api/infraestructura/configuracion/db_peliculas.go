package configuracion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	db  *gorm.DB
	err error
)

func GetDatabaseInstance() *gorm.DB {
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
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		_ = db.Close()
		panic(err)
	}

	db.SingularTable(true)
	db.AutoMigrate(&modelo.Pelicula{})

	return db
}

func RefreshPeliculaTable() {
	db.DropTable(&modelo.Pelicula{})
	db.AutoMigrate(&modelo.Pelicula{})
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
	db.Create(&pelicula)
	return pelicula, nil
}

func SendVariousPeliculas() {
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

	for i := range peliculas {
		db.Create(&peliculas[i])
	}
}
