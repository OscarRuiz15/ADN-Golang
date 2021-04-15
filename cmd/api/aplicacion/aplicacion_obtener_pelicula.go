package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
)

type AplicacionObtenerPelicula interface {
	Ejecutar(id int64) (modelo.Pelicula, error)
}

type ObtenerPelicula struct {
	ServicioObtenerPelicula servicio.PuertoServicioObtenerPelicula
}

func (obtenerPelicula *ObtenerPelicula) Ejecutar(id int64) (modelo.Pelicula, error) {

	return obtenerPelicula.ServicioObtenerPelicula.Obtener(id)
}
