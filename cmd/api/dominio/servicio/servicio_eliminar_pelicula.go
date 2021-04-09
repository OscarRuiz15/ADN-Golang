package servicio

import (
	"ADN_Golang/cmd/api/dominio/exception"
	"ADN_Golang/cmd/api/dominio/puerto"
	"errors"
	"log"
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
		err = exception.DataNotFound{ErrMessage: "El registro a eliminar no existe"}
		log.Println("Servicio Eliminar Pelicula -> El registro a eliminar no existe", err)
		return err
	}

	err = servicioEliminarPelicula.RepositorioPelicula.Eliminar(id)
	if err != nil {
		err = errors.New("Servicio Eliminar Pelicula -> Error al eliminar")
		log.Println("Servicio Eliminar Pelicula -> Error al eliminar", err)
		return err
	}

	return err
}
