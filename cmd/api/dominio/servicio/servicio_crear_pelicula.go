package servicio

import (
	"ADN_Golang/cmd/api/dominio/exception"
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
	"errors"
	"fmt"
)

type PuertoServicioCrearPelicula interface {
	Crear(pelicula *modelo.Pelicula) error
}

type ServicioCrearPelicula struct {
	RepositorioPelicula puerto.RepositorioPelicula
}

func (servicioCrearPelicula *ServicioCrearPelicula) Crear(pelicula *modelo.Pelicula) error {
	err := pelicula.Validar()
	if err != nil {
		return err
	}

	_, existe := servicioCrearPelicula.RepositorioPelicula.Existe(pelicula.Nombre)
	if existe {
		errMsg := fmt.Sprintf("La pelicula %s ya está registrada", pelicula.Nombre)
		err = exception.DataDuplicity{ErrMessage: errMsg}
		return err
	}

	err = servicioCrearPelicula.RepositorioPelicula.Crear(pelicula)
	if err != nil {
		errMsg := fmt.Sprintf("Servicio crear -> Error al crear pelicula: %s", err)
		err = errors.New(errMsg)
		return err
	}

	return err
}
