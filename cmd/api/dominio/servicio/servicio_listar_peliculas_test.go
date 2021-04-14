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

func TestListarPeliculasExitoso(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().Build()
	peliculas := []modelo.Pelicula{pelicula}

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Listar").Return(peliculas, nil)

	servicioListarPeliculas := servicio.ServicioListarPeliculas{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	peliculasResponse, _ := servicioListarPeliculas.Listar()

	// assert
	assert.Equal(t, len(peliculas), len(peliculasResponse))
	assert.Equal(t, peliculas[0].Id, peliculasResponse[0].Id)
	assert.Equal(t, peliculas[0].Nombre, peliculasResponse[0].Nombre)
	assert.Equal(t, peliculas[0].Director, peliculasResponse[0].Director)
	assert.Equal(t, peliculas[0].Escritor, peliculasResponse[0].Escritor)
	assert.Equal(t, peliculas[0].Pais, peliculasResponse[0].Pais)
	assert.Equal(t, peliculas[0].Idioma, peliculasResponse[0].Idioma)
	assert.Equal(t, peliculas[0].Lanzamiento, peliculasResponse[0].Lanzamiento)
}

func TestRepositorioListarPeliculasRetornaError(t *testing.T) {
	// arrange
	peliculas := make([]modelo.Pelicula, 0)
	errorEsperado := errors.New(servicio.ErrorListarPeliculas)

	repositorioPelicula := new(mock.RepositorioPeliculaMock)
	repositorioPelicula.On("Listar").Return(peliculas, errorEsperado)

	servicioListarPeliculas := servicio.ServicioListarPeliculas{
		RepositorioPelicula: repositorioPelicula,
	}

	// act
	_, err := servicioListarPeliculas.Listar()

	// assert
	assert.EqualError(t, err, errorEsperado.Error())
}
