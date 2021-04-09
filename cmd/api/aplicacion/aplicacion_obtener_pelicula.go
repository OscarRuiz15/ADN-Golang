package aplicacion

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/servicio"
	"fmt"
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
		fmt.Println("Aplicacion obtener -> Error", err)
	}

	return pelicula, err
}
