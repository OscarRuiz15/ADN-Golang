package app

import "ADN_Golang/cmd/api/infraestructura/contenedor"

func mapeoUrls() {
	api := router.Group("/api")
	api.GET("/peliculas", contenedor.GetControladorPelicula().Listar)
	api.GET("/peliculas/:id", contenedor.GetControladorPelicula().Obtener)
	api.POST("/peliculas", contenedor.GetControladorPelicula().Crear)
	api.PUT("/peliculas/:id", contenedor.GetControladorPelicula().Actualizar)
	api.DELETE("/peliculas/:id", contenedor.GetControladorPelicula().Eliminar)
}
