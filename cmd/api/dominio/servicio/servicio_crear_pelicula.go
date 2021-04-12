package servicio

import (
	"ADN_Golang/cmd/api/dominio/exception"
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
	"errors"
	"log"
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
		err = exception.DataDuplicity{ErrMessage: "La pelicula " + pelicula.Nombre + " ya está registrada"}
		log.Println("Servicio crear -> La pelicula"+pelicula.Nombre+" ya está registrada", err)
		return err
	}

	err = servicioCrearPelicula.RepositorioPelicula.Crear(pelicula)
	if err != nil {
		err = errors.New("Error al crear pelicula")
		log.Println("Servicio crear -> Error al crear pelicula", err)
		return err
	}

	return err
}
