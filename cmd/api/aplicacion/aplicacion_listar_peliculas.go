package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
)

type AplicacionListaPelicular interface {
	Ejecutar() ([]modelo.Pelicula, error)
}

type ListarPeliculas struct {
	ServicioListarPeliculas servicio.PuertoServicioListarPeliculas
}

func (listarPeliculas *ListarPeliculas) Ejecutar() ([]modelo.Pelicula, error) {

	return listarPeliculas.ServicioListarPeliculas.Listar()
}
