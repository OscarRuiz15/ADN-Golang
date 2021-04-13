package mock

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"github.com/stretchr/testify/mock"
)

type RepositorioPeliculaMock struct {
	mock.Mock
}

func (mock *RepositorioPeliculaMock) Existe(nombre string) (int64, bool) {
	args := mock.Called(nombre)
	return args.Get(0).(int64), args.Bool(1)
}

func (mock *RepositorioPeliculaMock) Actualizar(id int64, pelicula modelo.Pelicula) error {
	args := mock.Called(id, pelicula)
	return args.Error(0)
}

func (mock *RepositorioPeliculaMock) Crear(pelicula *modelo.Pelicula) error {
	args := mock.Called(pelicula)
	return args.Error(0)
}

func (mock *RepositorioPeliculaMock) Eliminar(id int64) error {
	args := mock.Called(id)
	return args.Error(0)
}

func (mock *RepositorioPeliculaMock) Listar() ([]modelo.Pelicula, error) {
	args := mock.Called()
	return args.Get(0).([]modelo.Pelicula), args.Error(1)
}

func (mock *RepositorioPeliculaMock) Obtener(id int64) (modelo.Pelicula, error) {
	args := mock.Called(id)
	return args.Get(0).(modelo.Pelicula), args.Error(1)
}
