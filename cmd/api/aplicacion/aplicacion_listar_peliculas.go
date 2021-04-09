package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
)

type AplicacionListaPelicular interface {
	Handler() ([]modelo.Pelicula, error)
}

type ListarPeliculas struct {
	ServicioListarPeliculas servicio.PuertoServicioListarPeliculas
}

func (listarPeliculas *ListarPeliculas) Handler() ([]modelo.Pelicula, error) {

	peliculas, err := listarPeliculas.ServicioListarPeliculas.Listar()
	return peliculas, err
}
