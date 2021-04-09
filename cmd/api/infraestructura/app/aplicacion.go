package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

var router = gin.Default()

func IniciarAplicacion() {
	err := godotenv.Load(".env")
	if err != nil {
		//Error al cargar archivo de ambiente
	}

	mapeoUrls()

	if err := router.Run(":8080"); err != nil {
		//Error arrancando servidor
	} else {
		log.Println("Aplicación iniciada")
	}
}
