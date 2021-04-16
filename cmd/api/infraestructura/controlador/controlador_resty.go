package controlador

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type usuario struct {
	Id               int    `json:"id"`
	Uid              string `json:"uid"`
	Nombre           string `json:"nombre"`
	Email            string `json:"email"`
	Foto             string `json:"foto"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
	Telefono         string `json:"telefono"`
	Genero           string `json:"genero"`
	Departamento     string `json:"departamento"`
	Municipio        string `json:"municipio"`
}

type producto struct {
	Id          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Precio      string `json:"precio"`
	Foto        string `json:"foto"`
	Lugar       int    `json:"lugar"`
}

type pelicula struct {
	Page          int         `json:"page"`
	Results       interface{} `json:"results"`
	Total_pages   int         `json:"total_pages"`
	Total_results int         `json:"total_results"`
}

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

	var peliculas []pelicula
	bytes := []byte("[" + resp.String() + "]")
	_ = json.Unmarshal(bytes, &peliculas)
	fmt.Println("La data", peliculas)
	context.JSON(http.StatusOK, peliculas)
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

	var productos []producto
	bytes := []byte(resp.String())
	_ = json.Unmarshal(bytes, &productos)
	context.JSON(http.StatusOK, productos)
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

	var usuarios []usuario
	bytes := []byte(resp.String())
	_ = json.Unmarshal(bytes, &usuarios)
	context.JSON(http.StatusOK, usuarios)
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
