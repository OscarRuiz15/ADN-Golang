package app

import (
	"ADN_Golang/cmd/api/infraestructura/contenedor"
	"ADN_Golang/cmd/api/infraestructura/controlador"
)

func mapeoUrls() {
	api := router.Group("/api")
	api.GET("/peliculas", contenedor.GetControladorPelicula().Listar)
	api.GET("/peliculas/:id", contenedor.GetControladorPelicula().Obtener)
	api.POST("/peliculas", contenedor.GetControladorPelicula().Crear)
	api.PUT("/peliculas/:id", contenedor.GetControladorPelicula().Actualizar)
	api.DELETE("/peliculas/:id", contenedor.GetControladorPelicula().Eliminar)

	api.GET("/peliculas_resty", controlador.GetPeliculasResty)
	api.GET("/productos_resty", controlador.GetProductosResty)
	api.GET("/usuarios_resty", controlador.GetUsuariosResty)
	api.POST("/usuarios_resty", controlador.PostUsuariosResty)
}
