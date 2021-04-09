package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
	"log"
)

type AplicacionCrearPelicula interface {
	Ejecutar(pelicula *modelo.Pelicula) error
}

type CrearPelicula struct {
	ServicioCrearPelicula servicio.PuertoServicioCrearPelicula
}

func (crearPelicula *CrearPelicula) Ejecutar(pelicula *modelo.Pelicula) error {

	err := crearPelicula.ServicioCrearPelicula.Crear(pelicula)
	if err != nil {
		log.Println("Aplicacion crear -> Error", err)
		return err
	}

	return err
}
