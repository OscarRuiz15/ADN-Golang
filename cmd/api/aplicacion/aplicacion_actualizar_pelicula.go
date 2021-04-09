package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
)

type AplicacionActualizarPelicula interface {
	Handler(id int64, pelicula modelo.Pelicula) error
}

type ActualizarPelicula struct {
	ServicioActualizarPelicula servicio.PuertoServicioActualizarPelicula
}

func (actualizarPelicula *ActualizarPelicula) Handler(id int64, pelicula modelo.Pelicula) error {

	err := actualizarPelicula.ServicioActualizarPelicula.Actualizar(id, pelicula)
	return err
}
