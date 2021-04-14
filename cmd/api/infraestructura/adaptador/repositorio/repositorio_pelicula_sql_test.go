package repositorio_test

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
	"ADN_Golang/cmd/api/infraestructura/adaptador/repositorio"
	"ADN_Golang/cmd/test/builder"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	queryInsertarPelicula   = "INSERT INTO pelicula"
	queryListarPeliculas    = "SELECT (.+) FROM pelicula "
	queryObtenerPelicula    = "SELECT (.+) FROM pelicula (.+)"
	queryExistePelicula     = "SELECT (.+) FROM pelicula (.+)"
	queryEliminarPelicula   = "DELETE FROM pelicula"
	queryActualizarPelicula = "UPDATE pelicula (.+)"
)

func setUpRepositorioPelicula() (repositorioPelicula puerto.RepositorioPelicula, mock sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	repositorioPelicula = &repositorio.RepositorioPeliculaSql{
		Db: db,
	}

	return
}

func TestCrearPeliculaExitoso(t *testing.T) {
	pelicula := builder.NewPeliculaBuilder().Build()
	repositorioPelicula, dbMock := setUpRepositorioPelicula()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(queryInsertarPelicula).WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := repositorioPelicula.Crear(&pelicula)

	assert.Nil(t, err)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}

func TestListarPeliculasExitoso(t *testing.T) {
	peliculas := []modelo.Pelicula{builder.NewPeliculaBuilder().Build()}
	pelicula := builder.NewPeliculaBuilder().Build()
	repositorioPelicula, dbMock := setUpRepositorioPelicula()

	rows := sqlmock.NewRows([]string{"id", "nombre", "director", "escritor", "pais", "idioma", "lanzamiento"}).AddRow(
		pelicula.Id,
		pelicula.Nombre,
		pelicula.Director,
		pelicula.Escritor,
		pelicula.Pais,
		pelicula.Idioma,
		pelicula.Lanzamiento)
	dbMock.ExpectQuery(queryListarPeliculas).WillReturnRows(rows)

	peliculasList, err := repositorioPelicula.Listar()

	assert.Nil(t, err)
	assert.Nil(t, dbMock.ExpectationsWereMet())
	assert.Equal(t, peliculas, peliculasList)
}

func TestObtenerPeliculaExitoso(t *testing.T) {
	repositorioPelicula, dbMock := setUpRepositorioPelicula()
	pelicula := builder.NewPeliculaBuilder().Build()

	row := sqlmock.NewRows([]string{"id", "nombre", "director", "escritor", "pais", "idioma", "lanzamiento"}).AddRow(
		pelicula.Id,
		pelicula.Nombre,
		pelicula.Director,
		pelicula.Escritor,
		pelicula.Pais,
		pelicula.Idioma,
		pelicula.Lanzamiento)
	dbMock.ExpectQuery(queryObtenerPelicula).WillReturnRows(row)

	response, err := repositorioPelicula.Obtener(pelicula.Id)

	assert.Nil(t, err)
	assert.Nil(t, dbMock.ExpectationsWereMet())
	assert.Equal(t, pelicula, response)
}

func TestExistePeliculaExitoso(t *testing.T) {
	var id int64 = 5
	var nombre = "Pelicula Test"
	repositorioPelicula, dbMock := setUpRepositorioPelicula()

	row := sqlmock.NewRows([]string{"id"}).AddRow(id)
	dbMock.ExpectQuery(queryExistePelicula).WillReturnRows(row)

	response, existe := repositorioPelicula.Existe(nombre)

	assert.True(t, existe)
	assert.Nil(t, dbMock.ExpectationsWereMet())
	assert.Equal(t, id, response)
}

func TestNoExistePeliculaExitoso(t *testing.T) {
	var id int64 = 0
	var nombre = "Pelicula Test"
	repositorioPelicula, dbMock := setUpRepositorioPelicula()

	row := sqlmock.NewRows([]string{"id"}).AddRow(id)
	dbMock.ExpectQuery(queryExistePelicula).WillReturnRows(row)

	response, existe := repositorioPelicula.Existe(nombre)

	assert.False(t, existe)
	assert.Nil(t, dbMock.ExpectationsWereMet())
	assert.Equal(t, id, response)
}

func TestEliminarPeliculaExitoso(t *testing.T) {
	var id int64 = 1
	repositorioPelicula, dbMock := setUpRepositorioPelicula()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(queryEliminarPelicula).WillReturnResult(sqlmock.NewResult(0, 1))
	dbMock.ExpectCommit()

	err := repositorioPelicula.Eliminar(id)

	assert.Nil(t, err)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}

func TestActualizarPeliculaExitoso(t *testing.T) {
	pelicula := builder.NewPeliculaBuilder().Build()
	repositorioPelicula, dbMock := setUpRepositorioPelicula()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(queryActualizarPelicula).WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := repositorioPelicula.Actualizar(pelicula.Id, pelicula)

	assert.Nil(t, err)
	assert.Nil(t, dbMock.ExpectationsWereMet())
}
