package servicio_test

import (
	"ADN_Golang/cmd/api/dominio/servicio"
	"ADN_Golang/cmd/test/builder"
	"ADN_Golang/cmd/test/mock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrearPeliculaExitoso(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().Build()
	var id int64 = 0

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Existe", pelicula.Nombre).Return(id, false)
	repositorioPelicula.On("Crear", &pelicula).Return(nil)

	servicioCrearPelicula := servicio.ServicioCrearPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	err := servicioCrearPelicula.Crear(&pelicula)

	// assert
	assert.Equal(t, err, nil)
}

func TestCrearPeliculaYaExistenteRetornaError(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().Build()
	var idPeliculaExistente int64 = 1

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Existe", pelicula.Nombre).Return(idPeliculaExistente, true)

	servicioCrearPelicula := servicio.ServicioCrearPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	err := servicioCrearPelicula.Crear(&pelicula)

	// assert
	assert.Equal(t, err.Error(), "La pelicula "+pelicula.Nombre+" ya está registrada")
}

func TestEnviarPeliculaARepositorioRetornaError(t *testing.T) {
	// arrange
	var id int64 = 0
	pelicula := builder.NewPeliculaBuilder().Build()
	errorEsperado := errors.New(servicio.ErrorCrearPelicula)

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Existe", pelicula.Nombre).Return(id, false)
	repositorioPelicula.On("Crear", &pelicula).Return(errorEsperado)

	servicioCrearPelicula := servicio.ServicioCrearPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	err := servicioCrearPelicula.Crear(&pelicula)

	// assert
	assert.EqualError(t, err, errorEsperado.Error())
}
