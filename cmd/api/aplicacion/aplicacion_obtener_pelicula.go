package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
	"log"
)

type AplicacionObtenerPelicula interface {
	Ejecutar(id int64) (modelo.Pelicula, error)
}

type ObtenerPelicula struct {
	ServicioObtenerPelicula servicio.PuertoServicioObtenerPelicula
}

func (obtenerPelicula *ObtenerPelicula) Ejecutar(id int64) (modelo.Pelicula, error) {

	pelicula, err := obtenerPelicula.ServicioObtenerPelicula.Obtener(id)
	if err != nil {
		log.Println("Aplicacion obtener -> Error", err)
		return modelo.Pelicula{}, err
	}

	return pelicula, err
}
