package controlador

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func GetPeliculasResty(context *gin.Context) {
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"api_key": "6dccec3eb84d278599ffaa9c9620e96e",
		}).
		SetHeader("Accept", "application/json").
		Get("https://api.themoviedb.org/3/discover/movie")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)

	context.String(http.StatusOK, resp.String())
}

func GetProductosResty(context *gin.Context) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get("https://bkndtg.herokuapp.com/productos/api/")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)

	context.String(http.StatusOK, resp.String())
}

func GetUsuariosResty(context *gin.Context) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get("https://bkndtg.herokuapp.com/usuarios/api/list/")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)

	context.String(http.StatusOK, resp.String())
}

func PostUsuariosResty(context *gin.Context) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{
						"uid": "ASDASASDASDASD",
						"nombre": "Pedro Pablo",
						"email": "algo@algo.com",
						"foto": "",
						"fecha_nacimiento": "1996-08-12",
						"telefono": "123446",
						"genero": "Hombre",
						"departamento": "Antioquia",
						"municipio": "Medellín"
					}`).
		Post("https://bkndtg.herokuapp.com/usuarios/api/")

	if err != nil {
		fmt.Println("Error", err)
	}

	context.JSON(resp.StatusCode(), map[string]string{"response": "ok"})
}
