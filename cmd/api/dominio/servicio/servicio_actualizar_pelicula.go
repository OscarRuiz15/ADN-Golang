package servicio

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
	"fmt"
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
		fmt.Println("Servicio actualizar -> Error al buscar pelicula", err)
		return err
	}

	err = servicioActualizarPelicula.RepositorioPelicula.Actualizar(id, pelicula)
	return err
}
