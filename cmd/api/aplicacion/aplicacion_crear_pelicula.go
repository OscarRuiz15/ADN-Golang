package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
)

type AplicacionCrearPelicula interface {
	Ejecutar(pelicula *modelo.Pelicula) error
}

type CrearPelicula struct {
	ServicioCrearPelicula servicio.PuertoServicioCrearPelicula
}

func (crearPelicula *CrearPelicula) Ejecutar(pelicula *modelo.Pelicula) error {

	return crearPelicula.ServicioCrearPelicula.Crear(pelicula)
}
