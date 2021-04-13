package modelo_test

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/test/builder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrearModeloPeliculaExitoso(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().Build()

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, nil, err)
}

func TestCrearModeloPeliculaSinNombreError(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().ConNombre("").Build()

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), modelo.EL_NOMBRE_ES_OBLIGATORIO)
}

func TestCrearModeloPeliculaSinDirectorError(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().ConDirector("").Build()

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), modelo.EL_DIRECTOR_ES_OBLIGATORIO)
}

func TestCrearModeloPeliculaSinEscritorError(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().ConEscritor("").Build()

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), modelo.EL_ESCRITOR_ES_OBLIGATORIO)
}

func TestCrearModeloPeliculaSinPaisError(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().ConPais("").Build()

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), modelo.EL_PAIS_ES_OBLIGATORIO)
}

func TestCrearModeloPeliculaSinIdiomaError(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().ConIdioma("").Build()

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), modelo.EL_IDIOMA_ES_OBLIGATORIO)
}

func TestCrearModeloPeliculaSinLanzamientoError(t *testing.T) {
	// arrange
	pelicula := builder.NewPeliculaBuilder().ConLanzamiento(0).Build()

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), modelo.EL_ANIO_ES_OBLIGATORIO)
}
