package servicio

import (
	"ADN_Golang/cmd/api/dominio/exception"
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
	"fmt"
	"github.com/pkg/errors"
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
		errMsg := fmt.Sprintf("No existe la pelicula a actualizar con el id %v", id)
		err = exception.DataNotFound{ErrMessage: errMsg}
		return err
	}

	err = pelicula.Validar()
	if err != nil {
		return err
	}

	idReg, existe := servicioActualizarPelicula.RepositorioPelicula.Existe(pelicula.Nombre)
	if existe && id != idReg {
		errMsg := fmt.Sprintf("La pelicula %s ya está registrada", pelicula.Nombre)
		err = exception.DataDuplicity{ErrMessage: errMsg}
		return err
	}

	err = servicioActualizarPelicula.RepositorioPelicula.Actualizar(id, pelicula)
	if err != nil {
		errMsg := fmt.Sprintf("Servicio actualizar -> Error al actualizar pelicula: %s", err)
		err = errors.New(errMsg)
		return err
	}

	return err
}
