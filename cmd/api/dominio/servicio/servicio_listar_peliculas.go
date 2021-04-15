package servicio

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
)

const ErrorListarPeliculas = "Servicio listar -> Error al listar peliculas desde el repositorio."

type PuertoServicioListarPeliculas interface {
	Listar() ([]modelo.Pelicula, error)
}

type ServicioListarPeliculas struct {
	RepositorioPelicula puerto.RepositorioPelicula
}

func (servicioListarPeliculas *ServicioListarPeliculas) Listar() ([]modelo.Pelicula, error) {

	return servicioListarPeliculas.RepositorioPelicula.Listar()
}
