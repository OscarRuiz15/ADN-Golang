package servicio

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
	"fmt"
	"github.com/pkg/errors"
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
		errMsg := fmt.Sprintf("Servicio listar -> Error al listar peliculas: %s", err)
		err = errors.New(errMsg)
		return nil, err
	}

	return peliculas, err
}
