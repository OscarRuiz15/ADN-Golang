package contenedor

import "ADN_Golang/cmd/api/aplicacion"

func getAplicacionCrearPelicula() aplicacion.AplicacionCrearPelicula {
	return &aplicacion.CrearPelicula{
		ServicioCrearPelicula: getServicioCrearPelicula(),
	}
}

func getAplicacionObtenerPelicula() aplicacion.AplicacionObtenerPelicula {
	return &aplicacion.ObtenerPelicula{
		ServicioObtenerPelicula: getServicioObtenerPelicula(),
	}
}

func getAplicacionListarPeliculas() aplicacion.AplicacionListaPelicular {
	return &aplicacion.ListarPeliculas{
		ServicioListarPeliculas: getServicioListarPeliculas(),
	}
}

func getAplicacionEliminarPelicula() aplicacion.AplicacionEliminarPelicula {
	return &aplicacion.EliminarPelicula{
		ServicioEliminarPelicula: getServicioEliminarPelicula(),
	}
}

func getAplicacionActualizarPelicula() aplicacion.AplicacionActualizarPelicula {
	return &aplicacion.ActualizarPelicula{
		ServicioActualizarPelicula: getServicioActualizarPelicula(),
	}
}
