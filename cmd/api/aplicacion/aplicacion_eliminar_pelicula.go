package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/servicio"
	"fmt"
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
		fmt.Println("Aplicacion eliminar -> Error", err)
	}

	return err
}
