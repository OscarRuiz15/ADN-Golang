package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/servicio"
	"log"
)

type AplicacionEliminarPelicula interface {
	Ejecutar(id int64) error
}

type EliminarPelicula struct {
	ServicioEliminarPelicula servicio.PuertoServicioEliminarPelicula
}

func (eliminarPelicula *EliminarPelicula) Ejecutar(id int64) error {

	err := eliminarPelicula.ServicioEliminarPelicula.Eliminar(id)
	if err != nil {
		log.Println("Aplicacion eliminar -> Error", err)
		return err
	}

	return err
}
