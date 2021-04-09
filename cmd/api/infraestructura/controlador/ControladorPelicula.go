package controlador

import (
	"ADN_Golang/cmd/api/aplicacion"
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/infraestructura/marshaller"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ControladorPelicula struct {
	AplicacionCrearPelicula      aplicacion.AplicacionCrearPelicula
	AplicacionObtenerPelicula    aplicacion.AplicacionObtenerPelicula
	AplicacionListaPelicular     aplicacion.AplicacionListaPelicular
	AplicacionEliminarPelicula   aplicacion.AplicacionEliminarPelicula
	AplicacionActualizarPelicula aplicacion.AplicacionActualizarPelicula
}

func obtenerIdPelicula(parametro string) int64 {
	id, err := strconv.ParseInt(parametro, 10, 64)
	if err != nil {
		fmt.Println("Id -> El id debe ser un número", err)
		return 0
	}
	return id
}

func (controladorPelicula *ControladorPelicula) Crear(context *gin.Context) {
	var pelicula modelo.Pelicula
	if err := context.ShouldBindJSON(&pelicula); err != nil {
		fmt.Println("Controlador Crear -> Error al parsear", err)
		return
	}

	err := controladorPelicula.AplicacionCrearPelicula.Handler(&pelicula)
	if err != nil {
		fmt.Println("Controlador Crear -> Error al crear", err)
		return
	}

	context.JSON(http.StatusCreated, pelicula)
}

func (controladorPelicula *ControladorPelicula) Obtener(context *gin.Context) {
	id := obtenerIdPelicula(context.Param("id"))
	pelicula, err := controladorPelicula.AplicacionObtenerPelicula.Handler(id)
	if err != nil {
		fmt.Println("Controlador Obtener -> Error al obtener", err)
		return
	}

	context.JSON(http.StatusOK, pelicula)
}

func (controladorPelicula *ControladorPelicula) Listar(context *gin.Context) {

	peliculas, err := controladorPelicula.AplicacionListaPelicular.Handler()
	if err != nil {
		fmt.Println("Controlador Listar -> Error al listar", err)
		return
	}

	context.JSON(http.StatusOK, marshaller.MarshallArray(peliculas))
}

func (controladorPelicula *ControladorPelicula) Eliminar(context *gin.Context) {
	id := obtenerIdPelicula(context.Param("id"))
	err := controladorPelicula.AplicacionEliminarPelicula.Handler(id)
	if err != nil {
		fmt.Println("Controlador Eliminar -> Error al eliminar", err)
		return
	}

	context.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (controladorPelicula *ControladorPelicula) Actualizar(context *gin.Context) {
	id := obtenerIdPelicula(context.Param("id"))
	var pelicula modelo.Pelicula
	if err := context.ShouldBindJSON(&pelicula); err != nil {
		fmt.Println("Controlador Actualizar -> Error al parsear", err)
		return
	}

	err := controladorPelicula.AplicacionActualizarPelicula.Handler(id, pelicula)
	if err != nil {
		fmt.Println("Controlador Actualizar -> Error al actualizar", err)
		return
	}

	context.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
