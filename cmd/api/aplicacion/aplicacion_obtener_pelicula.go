package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
)

type AplicacionObtenerPelicula interface {
	Handler(id int64) (modelo.Pelicula, error)
}

type ObtenerPelicula struct {
	ServicioObtenerPelicula servicio.PuertoServicioObtenerPelicula
}

func (obtenerPelicula *ObtenerPelicula) Handler(id int64) (modelo.Pelicula, error) {

	pelicula, err := obtenerPelicula.ServicioObtenerPelicula.Obtener(id)
	return pelicula, err
}
