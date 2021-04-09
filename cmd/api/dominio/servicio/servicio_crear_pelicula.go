package servicio

import (
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

	err := servicioCrearPelicula.RepositorioPelicula.Crear(pelicula)
	if err != nil {
		err = errors.New("Error al crear pelicula")
		log.Println("Servicio crear -> Error al crear pelicula", err)
		return err
	}

	return err
}
