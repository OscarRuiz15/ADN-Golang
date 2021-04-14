package servicio

import (
	"ADN_Golang/cmd/api/dominio/exception"
	"ADN_Golang/cmd/api/dominio/puerto"
	"errors"
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
		errMsg := fmt.Sprintf("No existe la pelicula a eliminar con el id %v", id)
		err = exception.DataNotFound{ErrMessage: errMsg}
		return err
	}

	err = servicioEliminarPelicula.RepositorioPelicula.Eliminar(id)
	if err != nil {
		errMsg := fmt.Sprintf("Servicio eliminar -> Error al eliminar pelicula: %s", err)
		err = errors.New(errMsg)
		return err
	}

	return err
}
