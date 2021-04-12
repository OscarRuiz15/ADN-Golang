package modelo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrearModeloPeliculaExitoso(t *testing.T) {
	// arrange
	pelicula := Pelicula{
		Nombre:      "Oscar",
		Director:    "Alexander",
		Escritor:    "Ruiz",
		Pais:        "Colombia",
		Idioma:      "Español",
		Lanzamiento: 2021,
	}

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, nil, err)
}

func TestCrearModeloPeliculaSinNombreError(t *testing.T) {
	// arrange
	pelicula := Pelicula{
		Director:    "Alexander",
		Escritor:    "Ruiz",
		Pais:        "Colombia",
		Idioma:      "Español",
		Lanzamiento: 2021,
	}

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), EL_NOMBRE_ES_OBLIGATORIO)
}

func TestCrearModeloPeliculaSinDirectorError(t *testing.T) {
	// arrange
	pelicula := Pelicula{
		Nombre:      "Oscar",
		Escritor:    "Ruiz",
		Pais:        "Colombia",
		Idioma:      "Español",
		Lanzamiento: 2021,
	}

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), EL_DIRECTOR_ES_OBLIGATORIO)
}

func TestCrearModeloPeliculaSinEscritorError(t *testing.T) {
	// arrange
	pelicula := Pelicula{
		Nombre:      "Oscar",
		Director:    "Alexander",
		Pais:        "Colombia",
		Idioma:      "Español",
		Lanzamiento: 2021,
	}

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), EL_ESCRITOR_ES_OBLIGATORIO)
}

func TestCrearModeloPeliculaSinPaisError(t *testing.T) {
	// arrange
	pelicula := Pelicula{
		Nombre:      "Oscar",
		Director:    "Alexander",
		Escritor:    "Ruiz",
		Idioma:      "Español",
		Lanzamiento: 2021,
	}

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), EL_PAIS_ES_OBLIGATORIO)
}

func TestCrearModeloPeliculaSinIdiomaError(t *testing.T) {
	// arrange
	pelicula := Pelicula{
		Nombre:      "Oscar",
		Director:    "Alexander",
		Escritor:    "Ruiz",
		Pais:        "Colombia",
		Lanzamiento: 2021,
	}

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), EL_IDIOMA_ES_OBLIGATORIO)
}

func TestCrearModeloPeliculaSinLanzamientoError(t *testing.T) {
	// arrange
	pelicula := Pelicula{
		Nombre:   "Oscar",
		Director: "Alexander",
		Escritor: "Ruiz",
		Pais:     "Colombia",
		Idioma:   "Español",
	}

	// act
	err := pelicula.Validar()

	// assert
	assert.Equal(t, err.Error(), EL_ANIO_ES_OBLIGATORIO)
}
