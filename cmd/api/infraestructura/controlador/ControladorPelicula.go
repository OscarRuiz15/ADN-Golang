package controlador

import (
	"ADN_Golang/cmd/api/aplicacion"
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/infraestructura/marshaller"
	apierrors "ADN_Golang/pkg/apierror"
	"github.com/gin-gonic/gin"
	"log"
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

func obtenerIdPelicula(parametro string, context *gin.Context) int64 {
	id, err := strconv.ParseInt(parametro, 10, 64)
	if err != nil {
		log.Println("Id -> El id debe ser un número", err)
		idErr := apierrors.NewApiError("Id Invalido", err.Error(), 400)
		context.JSON(idErr.Status(), idErr)
		return 0
	}
	return id
}

func (controladorPelicula *ControladorPelicula) Crear(context *gin.Context) {
	var pelicula modelo.Pelicula
	if err := context.ShouldBindJSON(&pelicula); err != nil {
		log.Println("Controlador Crear -> Error al parsear", err)
		restErr := apierrors.NewApiError("JSON Invalido", err.Error(), 400)
		context.JSON(restErr.Status(), restErr)
		return
	}

	err := controladorPelicula.AplicacionCrearPelicula.Ejecutar(&pelicula)
	if err != nil {
		log.Println("Controlador Crear -> Error al crear", err)
		abort(context, err)
		return
	}

	context.JSON(http.StatusCreated, pelicula)
}

func (controladorPelicula *ControladorPelicula) Obtener(context *gin.Context) {
	id := obtenerIdPelicula(context.Param("id"), context)
	pelicula, err := controladorPelicula.AplicacionObtenerPelicula.Ejecutar(id)
	if err != nil {
		log.Println("Controlador Obtener -> Error al obtener", err)
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, pelicula)
}

func (controladorPelicula *ControladorPelicula) Listar(context *gin.Context) {

	peliculas, err := controladorPelicula.AplicacionListaPelicular.Ejecutar()
	if err != nil {
		log.Println("Controlador Listar -> Error al listar", err)
		err := apierrors.NewApiError("No hay peliculas", err.Error(), 404)
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, marshaller.MarshallArray(peliculas))
}

func (controladorPelicula *ControladorPelicula) Eliminar(context *gin.Context) {
	id := obtenerIdPelicula(context.Param("id"), context)
	err := controladorPelicula.AplicacionEliminarPelicula.Ejecutar(id)
	if err != nil {
		log.Println("Controlador Eliminar -> Error al eliminar", err)
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (controladorPelicula *ControladorPelicula) Actualizar(context *gin.Context) {
	id := obtenerIdPelicula(context.Param("id"), context)
	var pelicula modelo.Pelicula
	if err := context.ShouldBindJSON(&pelicula); err != nil {
		log.Println("Controlador Actualizar -> Error al parsear", err)
		restErr := apierrors.NewApiError("JSON Invalido", err.Error(), 400)
		context.JSON(restErr.Status(), restErr)
		return
	}

	err := controladorPelicula.AplicacionActualizarPelicula.Ejecutar(id, pelicula)
	if err != nil {
		log.Println("Controlador Actualizar -> Error al actualizar", err)
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func abort(ctx *gin.Context, err error) {
	ctx.Error(err)
	ctx.Abort()
}
