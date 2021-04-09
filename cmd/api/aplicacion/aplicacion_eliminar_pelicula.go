package aplicacion

import "ADN_Golang/cmd/api/dominio/servicio"

type AplicacionEliminarPelicula interface {
	Handler(id int64) error
}

type EliminarPelicula struct {
	ServicioEliminarPelicula servicio.PuertoServicioEliminarPelicula
}

func (eliminarPelicula *EliminarPelicula) Handler(id int64) error {

	err := eliminarPelicula.ServicioEliminarPelicula.Eliminar(id)
	return err
}
