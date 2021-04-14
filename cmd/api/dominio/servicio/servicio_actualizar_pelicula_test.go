package servicio_test

import (
	"ADN_Golang/cmd/api/dominio/exception"
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
	"ADN_Golang/cmd/test/builder"
	"ADN_Golang/cmd/test/mock"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActualizarPeliculaExitoso(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().Build()
	var id int64 = 0

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Obtener", pelicula.Id).Return(pelicula, nil)
	repositorioPelicula.On("Existe", pelicula.Nombre).Return(id, false)
	repositorioPelicula.On("Actualizar", pelicula.Id, pelicula).Return(nil)

	servicioActualizarPelicula := servicio.ServicioActualizarPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	err := servicioActualizarPelicula.Actualizar(pelicula.Id, pelicula)

	// assert
	assert.Equal(t, err, nil)
}

func TestPeliculaAActualizarNoExisteRetornaError(t *testing.T) {
	// arrange
	var id int64 = 99
	pelicula := builder.NewPeliculaBuilder().Build()
	errMsg := fmt.Sprintf("No existe la pelicula a actualizar con el id %v", id)
	errorEsperado := exception.DataNotFound{ErrMessage: errMsg}

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Obtener", id).Return(modelo.Pelicula{}, errorEsperado)

	servicioActualizarPelicula := servicio.ServicioActualizarPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	err := servicioActualizarPelicula.Actualizar(id, pelicula)

	// assert
	assert.EqualError(t, err, errorEsperado.Error())
}

func TestActualizarPeliculaConNombreYaExistenteRetornaError(t *testing.T) {
	// arrange
	var idPeliculaExistente int64 = 2
	pelicula := builder.NewPeliculaBuilder().Build()
	errorEsperado := exception.DataNotFound{ErrMessage: "La pelicula " + pelicula.Nombre + " ya está registrada"}

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Obtener", pelicula.Id).Return(pelicula, nil)
	repositorioPelicula.On("Existe", pelicula.Nombre).Return(idPeliculaExistente, true)

	servicioActualizarPelicula := servicio.ServicioActualizarPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	err := servicioActualizarPelicula.Actualizar(pelicula.Id, pelicula)

	// assert
	assert.EqualError(t, err, errorEsperado.Error())
}

func TestActualizarPeliculaRetornaErrorRepositorio(t *testing.T) {
	// arrange
	var id int64 = 0
	pelicula := builder.NewPeliculaBuilder().Build()
	errorEsperado := errors.New(servicio.ErrorActualizarPelicula)

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Obtener", pelicula.Id).Return(pelicula, nil)
	repositorioPelicula.On("Existe", pelicula.Nombre).Return(id, false)
	repositorioPelicula.On("Actualizar", pelicula.Id, pelicula).Return(errorEsperado)

	servicioActualizarPelicula := servicio.ServicioActualizarPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	err := servicioActualizarPelicula.Actualizar(pelicula.Id, pelicula)

	// assert
	assert.EqualError(t, err, errorEsperado.Error())
}
