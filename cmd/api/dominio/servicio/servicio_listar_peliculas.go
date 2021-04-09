package servicio

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
	"fmt"
)

type PuertoServicioListarPeliculas interface {
	Listar() ([]modelo.Pelicula, error)
}

type ServicioListarPeliculas struct {
	RepositorioPelicula puerto.RepositorioPelicula
}

func (servicioListarPeliculas *ServicioListarPeliculas) Listar() ([]modelo.Pelicula, error) {

	peliculas, err := servicioListarPeliculas.RepositorioPelicula.Listar()
	if err != nil {
		fmt.Println("Servicio Eliminar Pelicula -> Error al eliminar", err)
		return nil, err
	}

	return peliculas, err
}
