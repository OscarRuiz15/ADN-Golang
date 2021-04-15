package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
)

type AplicacionActualizarPelicula interface {
	Ejecutar(id int64, pelicula modelo.Pelicula) error
}

type ActualizarPelicula struct {
	ServicioActualizarPelicula servicio.PuertoServicioActualizarPelicula
}

func (actualizarPelicula *ActualizarPelicula) Ejecutar(id int64, pelicula modelo.Pelicula) error {

	return actualizarPelicula.ServicioActualizarPelicula.Actualizar(id, pelicula)
}
