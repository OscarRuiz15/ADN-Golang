package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
	"log"
)

type AplicacionActualizarPelicula interface {
	Ejecutar(id int64, pelicula modelo.Pelicula) error
}

type ActualizarPelicula struct {
	ServicioActualizarPelicula servicio.PuertoServicioActualizarPelicula
}

func (actualizarPelicula *ActualizarPelicula) Ejecutar(id int64, pelicula modelo.Pelicula) error {

	err := actualizarPelicula.ServicioActualizarPelicula.Actualizar(id, pelicula)
	if err != nil {
		log.Println("Aplicacion actualizar -> Error", err)
		return err
	}

	return err
}
