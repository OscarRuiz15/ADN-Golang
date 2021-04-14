package servicio

import (
	"ADN_Golang/cmd/api/dominio/exception"
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/dominio/puerto"
	"fmt"
)

type PuertoServicioObtenerPelicula interface {
	Obtener(id int64) (modelo.Pelicula, error)
}

type ServicioObtenerPelicula struct {
	RepositorioPelicula puerto.RepositorioPelicula
}

func (servicioObtenerPelicula *ServicioObtenerPelicula) Obtener(id int64) (modelo.Pelicula, error) {

	pelicula, err := servicioObtenerPelicula.RepositorioPelicula.Obtener(id)
	if err != nil {
		errMsg := fmt.Sprintf("No existe la pelicula con el id %v", id)
		err = exception.DataNotFound{ErrMessage: errMsg}
		return pelicula, err
	}

	return pelicula, err
}
