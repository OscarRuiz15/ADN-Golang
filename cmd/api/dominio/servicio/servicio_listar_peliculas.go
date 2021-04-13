package servicio

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
	"github.com/pkg/errors"
	"log"
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
		err = errors.New("Error al listar peliculas")
		log.Println("Servicio Listar Peliculas -> Error al listar peliculas", err)
		return nil, err
	}

	return peliculas, err
}
