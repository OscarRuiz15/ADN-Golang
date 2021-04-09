package contenedor

import (
	"ADN_Golang/cmd/api/dominio/puerto"
	"ADN_Golang/cmd/api/infraestructura/adaptador/repositorio"
	"ADN_Golang/cmd/api/infraestructura/configuracion"
	"ADN_Golang/cmd/api/infraestructura/controlador"
)

func GetControladorPelicula() *controlador.ControladorPelicula {
	return &controlador.ControladorPelicula{
		AplicacionCrearPelicula:      getAplicacionCrearPelicula(),
		AplicacionObtenerPelicula:    getAplicacionObtenerPelicula(),
		AplicacionListaPelicular:     getAplicacionListarPeliculas(),
		AplicacionEliminarPelicula:   getAplicacionEliminarPelicula(),
		AplicacionActualizarPelicula: getAplicacionActualizarPelicula(),
	}
}

func getRepositorioPelicula() puerto.RepositorioPelicula {
	return &repositorio.RepositorioPeliculaSql{
		Db: configuracion.GetDatabaseInstance(),
	}
}
