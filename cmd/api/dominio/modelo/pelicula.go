package modelo

import (
	"ADN_Golang/cmd/api/dominio/exception"
)

const (
	EL_NOMBRE_ES_OBLIGATORIO   = "El nombre de pelicula es obligatorio"
	EL_DIRECTOR_ES_OBLIGATORIO = "El director de pelicula es obligatorio"
	EL_ESCRITOR_ES_OBLIGATORIO = "El escritor de pelicula es obligatorio"
	EL_PAIS_ES_OBLIGATORIO     = "El pais de producción de pelicula es obligatorio"
	EL_IDIOMA_ES_OBLIGATORIO   = "El idioma de lanzamiento es obligatorio"
	EL_ANIO_ES_OBLIGATORIO     = "El año de lanzamiento es obligatorio"
)

type Pelicula struct {
	Id          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Director    string `json:"director"`
	Escritor    string `json:"escritor"`
	Pais        string `json:"pais"`
	Idioma      string `json:"idioma"`
	Lanzamiento int    `json:"lanzamiento"`
}

func (pelicula Pelicula) Validar() error {
	if pelicula.Nombre == "" {
		return exception.DataRequired{ErrMessage: EL_NOMBRE_ES_OBLIGATORIO}
	}

	if pelicula.Director == "" {
		return exception.DataRequired{ErrMessage: EL_DIRECTOR_ES_OBLIGATORIO}
	}

	if pelicula.Escritor == "" {
		return exception.DataRequired{ErrMessage: EL_ESCRITOR_ES_OBLIGATORIO}
	}

	if pelicula.Pais == "" {
		return exception.DataRequired{ErrMessage: EL_PAIS_ES_OBLIGATORIO}
	}

	if pelicula.Idioma == "" {
		return exception.DataRequired{ErrMessage: EL_IDIOMA_ES_OBLIGATORIO}
	}

	if pelicula.Lanzamiento == 0 {
		return exception.DataRequired{ErrMessage: EL_ANIO_ES_OBLIGATORIO}
	}

	return nil
}
