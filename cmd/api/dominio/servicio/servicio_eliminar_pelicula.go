package servicio

import (
	"ADN_Golang/cmd/api/dominio/puerto"
	"fmt"
)

type PuertoServicioEliminarPelicula interface {
	Eliminar(id int64) error
}

type ServicioEliminarPelicula struct {
	RepositorioPelicula puerto.RepositorioPelicula
}

func (servicioEliminarPelicula *ServicioEliminarPelicula) Eliminar(id int64) error {

	_, err := servicioEliminarPelicula.RepositorioPelicula.Obtener(id)
	if err != nil {
		fmt.Println("Servicio Eliminar Pelicula -> El registro a eliminar no existe", err)
	}

	err = servicioEliminarPelicula.RepositorioPelicula.Eliminar(id)
	return err
}
