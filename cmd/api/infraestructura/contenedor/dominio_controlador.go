package contenedor

import "ADN_Golang/cmd/api/dominio/servicio"

func getServicioCrearPelicula() servicio.PuertoServicioCrearPelicula {
	return &servicio.ServicioCrearPelicula{
		RepositorioPelicula: getRepositorioPelicula(),
	}
}

func getServicioObtenerPelicula() servicio.PuertoServicioObtenerPelicula {
	return &servicio.ServicioObtenerPelicula{
		RepositorioPelicula: getRepositorioPelicula(),
	}
}

func getServicioListarPeliculas() servicio.PuertoServicioListarPeliculas {
	return &servicio.ServicioListarPeliculas{
		RepositorioPelicula: getRepositorioPelicula(),
	}
}

func getServicioEliminarPelicula() servicio.PuertoServicioEliminarPelicula {
	return &servicio.ServicioEliminarPelicula{
		RepositorioPelicula: getRepositorioPelicula(),
	}
}

func getServicioActualizarPelicula() servicio.PuertoServicioActualizarPelicula {
	return &servicio.ServicioActualizarPelicula{
		RepositorioPelicula: getRepositorioPelicula(),
	}
}
