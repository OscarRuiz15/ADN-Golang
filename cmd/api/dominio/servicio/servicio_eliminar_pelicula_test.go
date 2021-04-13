package servicio_test

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
	"ADN_Golang/cmd/test/builder"
	"ADN_Golang/cmd/test/mock"
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
	errorEsperado := errors.New("El registro a eliminar no existe")

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
	errorEsperado := errors.New("Servicio Eliminar Pelicula -> Error al eliminar")

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
