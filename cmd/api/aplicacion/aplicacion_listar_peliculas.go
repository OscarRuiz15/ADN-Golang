package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
	"log"
)

type AplicacionListaPelicular interface {
	Ejecutar() ([]modelo.Pelicula, error)
}

type ListarPeliculas struct {
	ServicioListarPeliculas servicio.PuertoServicioListarPeliculas
}

func (listarPeliculas *ListarPeliculas) Ejecutar() ([]modelo.Pelicula, error) {

	peliculas, err := listarPeliculas.ServicioListarPeliculas.Listar()
	if err != nil {
		log.Println("Aplicacion listar -> Error", err)
		return nil, err
	}

	return peliculas, err
}
