package servicio

import (
	"ADN_Golang/cmd/api/dominio/exception"
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
	"github.com/pkg/errors"
	"log"
)

type PuertoServicioActualizarPelicula interface {
	Actualizar(id int64, pelicula modelo.Pelicula) error
}

type ServicioActualizarPelicula struct {
	RepositorioPelicula puerto.RepositorioPelicula
}

func (servicioActualizarPelicula *ServicioActualizarPelicula) Actualizar(id int64, pelicula modelo.Pelicula) error {

	_, err := servicioActualizarPelicula.RepositorioPelicula.Obtener(id)
	if err != nil {
		err = exception.DataNotFound{ErrMessage: "Error al buscar pelicula"}
		log.Println("Servicio actualizar -> Error al buscar pelicula", err)
		return err
	}

	err = servicioActualizarPelicula.RepositorioPelicula.Actualizar(id, pelicula)
	if err != nil {
		err = errors.New("Servicio actualizar -> Error al actualizar pelicula")
		log.Println("Servicio actualizar -> Error al actualizar pelicula", err)
		return err
	}

	return err
}
