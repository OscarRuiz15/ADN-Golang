package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/servicio"
)

type AplicacionEliminarPelicula interface {
	Ejecutar(id int64) error
}

type EliminarPelicula struct {
	ServicioEliminarPelicula servicio.PuertoServicioEliminarPelicula
}

func (eliminarPelicula *EliminarPelicula) Ejecutar(id int64) error {

	return eliminarPelicula.ServicioEliminarPelicula.Eliminar(id)
}
