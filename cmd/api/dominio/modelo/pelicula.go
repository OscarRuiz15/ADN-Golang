package modelo

import (
	"ADN_Golang/cmd/api/dominio/exception"
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
		return exception.DataRequired{ErrMessage: "El nombre de pelicula es obligatorio"}
	}

	if pelicula.Director == "" {
		return exception.DataRequired{ErrMessage: "El director de pelicula es obligatorio"}
	}

	if pelicula.Escritor == "" {
		return exception.DataRequired{ErrMessage: "El escritor de pelicula es obligatorio"}
	}

	if pelicula.Pais == "" {
		return exception.DataRequired{ErrMessage: "El pais de producción de pelicula es obligatorio"}
	}

	if pelicula.Idioma == "" {
		return exception.DataRequired{ErrMessage: "El idioma de lanzamiento es obligatorio"}
	}

	if pelicula.Lanzamiento == 0 {
		return exception.DataRequired{ErrMessage: "El año de lanzamiento es obligatorio"}
	}

	return nil
}
