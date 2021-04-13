package servicio_test

import (
	"ADN_Golang/cmd/api/dominio/exception"
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
	"ADN_Golang/cmd/test/builder"
	"ADN_Golang/cmd/test/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObtenerPeliculaExitoso(t *testing.T) {
	// arrange
	var id int64 = 1
	pelicula := builder.NewPeliculaBuilder().Build()

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Obtener", id).Return(pelicula, nil)

	servicioObtenerPelicula := servicio.ServicioObtenerPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	peliculaResponse, _ := servicioObtenerPelicula.Obtener(id)

	// assert
	assert.Equal(t, pelicula.Id, peliculaResponse.Id)
	assert.Equal(t, pelicula.Nombre, peliculaResponse.Nombre)
	assert.Equal(t, pelicula.Director, peliculaResponse.Director)
	assert.Equal(t, pelicula.Escritor, peliculaResponse.Escritor)
	assert.Equal(t, pelicula.Pais, peliculaResponse.Pais)
	assert.Equal(t, pelicula.Idioma, peliculaResponse.Idioma)
	assert.Equal(t, pelicula.Lanzamiento, peliculaResponse.Lanzamiento)
}

func TestEnviarIdDePeliculaAObtenerNoExistente(t *testing.T) {
	// arrange
	var id int64 = 1
	errorEsperado := exception.DataNotFound{ErrMessage: "No existe la pelicula con el id"}

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Obtener", id).Return(modelo.Pelicula{}, errorEsperado)

	servicioObtenerPelicula := servicio.ServicioObtenerPelicula{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	_, err := servicioObtenerPelicula.Obtener(id)

	// assert
	assert.EqualError(t, err, errorEsperado.Error())
}
