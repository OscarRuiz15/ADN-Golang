package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
)

type AplicacionCrearPelicula interface {
	Handler(pelicula *modelo.Pelicula) error
}

type CrearPelicula struct {
	ServicioCrearPelicula servicio.PuertoServicioCrearPelicula
}

func (crearPelicula *CrearPelicula) Handler(pelicula *modelo.Pelicula) error {

	err := crearPelicula.ServicioCrearPelicula.Crear(pelicula)
	return err
}
