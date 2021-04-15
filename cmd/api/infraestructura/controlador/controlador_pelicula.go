package controlador

import (
	"ADN_Golang/cmd/api/aplicacion"
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/infraestructura/exception"
	"ADN_Golang/cmd/api/infraestructura/marshaller"
	apierrors "ADN_Golang/pkg/apierror"
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

func obtenerIdPelicula(parametro string) (int64, error) {
	id, err := strconv.ParseInt(parametro, 10, 64)
	if err != nil {
		err = exception.InternalServerError{ErrMessage: "El id debe ser un número"}
		return 0, err
	}
	return id, nil
}

func (controladorPelicula *ControladorPelicula) Crear(context *gin.Context) {
	var pelicula modelo.Pelicula
	if err := context.ShouldBindJSON(&pelicula); err != nil {
		restErr := apierrors.NewApiError("JSON Invalido", err.Error(), 422)
		context.JSON(restErr.Status(), restErr)
		return
	}

	err := controladorPelicula.AplicacionCrearPelicula.Ejecutar(&pelicula)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusCreated, pelicula)
}

func (controladorPelicula *ControladorPelicula) Obtener(context *gin.Context) {
	id, err := obtenerIdPelicula(context.Param("id"))
	if err != nil {
		abort(context, err)
		return
	}

	pelicula, err := controladorPelicula.AplicacionObtenerPelicula.Ejecutar(id)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, pelicula)
}

func (controladorPelicula *ControladorPelicula) Listar(context *gin.Context) {

	peliculas, err := controladorPelicula.AplicacionListaPelicular.Ejecutar()
	if err != nil {
		err := apierrors.NewApiError("No hay peliculas", err.Error(), 404)
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, marshaller.MarshallArray(peliculas))
}

func (controladorPelicula *ControladorPelicula) Eliminar(context *gin.Context) {
	id, err := obtenerIdPelicula(context.Param("id"))
	if err != nil {
		abort(context, err)
		return
	}

	err = controladorPelicula.AplicacionEliminarPelicula.Ejecutar(id)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (controladorPelicula *ControladorPelicula) Actualizar(context *gin.Context) {
	id, err := obtenerIdPelicula(context.Param("id"))
	if err != nil {
		abort(context, err)
		return
	}

	var pelicula modelo.Pelicula
	if err = context.ShouldBindJSON(&pelicula); err != nil {
		restErr := apierrors.NewApiError("JSON Invalido", err.Error(), 422)
		context.JSON(restErr.Status(), restErr)
		return
	}

	err = controladorPelicula.AplicacionActualizarPelicula.Ejecutar(id, pelicula)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func abort(ctx *gin.Context, err error) {
	ctx.Error(err)
	ctx.Abort()
}
