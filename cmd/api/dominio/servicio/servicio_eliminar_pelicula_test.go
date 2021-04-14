package servicio_test

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
	"ADN_Golang/cmd/test/builder"
	"ADN_Golang/cmd/test/mock"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEliminarPeliculaExitoso(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().Build()

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Obtener", pelicula.Id).Return(pelicula, nil)
	repositorioPelicula.On("Eliminar", pelicula.Id).Return(nil)

	servicioEliminarPelicula := servicio.ServicioEliminarPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	err := servicioEliminarPelicula.Eliminar(pelicula.Id)

	// assert
	assert.Equal(t, err, nil)
}

func TestEliminarPeliculaQueNoExisteRetornaError(t *testing.T) {
	// arrange
	var id int64 = 1
	errMsg := fmt.Sprintf("No existe la pelicula a eliminar con el id %v", id)
	errorEsperado := errors.New(errMsg)

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Obtener", id).Return(modelo.Pelicula{}, errorEsperado)

	servicioEliminarPelicula := servicio.ServicioEliminarPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	err := servicioEliminarPelicula.Eliminar(id)

	// assert
	assert.EqualError(t, err, errorEsperado.Error())
}

func TestEnviarIdAEliminarARepositorioRetornaError(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().Build()
	errorEsperado := errors.New(servicio.ErrorEliminarPelicula)

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Obtener", pelicula.Id).Return(pelicula, nil)
	repositorioPelicula.On("Eliminar", pelicula.Id).Return(errorEsperado)

	servicioEliminarPelicula := servicio.ServicioEliminarPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	err := servicioEliminarPelicula.Eliminar(pelicula.Id)

	// assert
	assert.EqualError(t, err, errorEsperado.Error())
}
