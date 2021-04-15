package controlador_test

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/infraestructura/configuracion"
	"ADN_Golang/cmd/api/infraestructura/contenedor"
	"ADN_Golang/cmd/test/builder"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

const (
	urlPeliculas = "/peliculas"
)

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("./../../../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	os.Exit(m.Run())
}

func TestCrear(t *testing.T) {
	gin.SetMode(gin.TestMode)

	configuracion.GetDatabaseInstance()
	_ = configuracion.RefreshPeliculaTable()

	samples := []struct {
		inputJSON   string
		statusCode  int
		nombre      string
		director    string
		escritor    string
		pais        string
		idioma      string
		lanzamiento int
		errMessage  string
	}{
		{
			inputJSON:   builder.NewPeliculaBuilder().BuildString(),
			statusCode:  201,
			nombre:      "Oscar",
			director:    "Alexander",
			escritor:    "Ruiz",
			pais:        "Colombia",
			idioma:      "Español",
			lanzamiento: 2021,
			errMessage:  "",
		},
		{
			inputJSON: `{
							"nombre": "COVID",
							"director": "Alfonso",
							"escritor": "Yo",
							"pais": "Colombia",
							"idioma": "Ingles",
							"lanzamiento": "2018"
						}`,
			statusCode: 422,
			errMessage: "JSON Invalido",
		},
	}

	for _, v := range samples {
		r := gin.Default()
		r.POST(urlPeliculas, contenedor.GetControladorPelicula().Crear)
		req, err := http.NewRequest(http.MethodPost, urlPeliculas, bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		fmt.Println("Response data: ", responseMap)
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["nombre"], v.nombre)
			assert.Equal(t, responseMap["director"], v.director)
			assert.Equal(t, responseMap["escritor"], v.escritor)
			assert.Equal(t, responseMap["pais"], v.pais)
			assert.Equal(t, responseMap["idioma"], v.idioma)
		}
		if v.statusCode == 400 || v.statusCode == 422 || v.statusCode == 500 && v.errMessage != "" {
			assert.Equal(t, responseMap["message"], v.errMessage)
		}
	}
}

func TestObtener(t *testing.T) {
	gin.SetMode(gin.TestMode)

	configuracion.GetDatabaseInstance()
	_ = configuracion.RefreshPeliculaTable()
	pelicula, _ := configuracion.SendOnePelicula()

	samples := []struct {
		id          string
		statusCode  int
		nombre      string
		director    string
		escritor    string
		pais        string
		idioma      string
		lanzamiento int
		errMessage  string
	}{
		{
			id:          strconv.Itoa(int(pelicula.Id)),
			statusCode:  200,
			nombre:      pelicula.Nombre,
			director:    pelicula.Director,
			escritor:    pelicula.Escritor,
			pais:        pelicula.Pais,
			idioma:      pelicula.Idioma,
			lanzamiento: pelicula.Lanzamiento,
			errMessage:  "",
		},
	}

	for _, v := range samples {
		r := gin.Default()
		r.GET(urlPeliculas+"/:id", contenedor.GetControladorPelicula().Obtener)
		req, err := http.NewRequest(http.MethodGet, urlPeliculas+"/"+v.id, nil)
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, responseMap["nombre"], v.nombre)
			assert.Equal(t, responseMap["director"], v.director)
			assert.Equal(t, responseMap["escritor"], v.escritor)
			assert.Equal(t, responseMap["pais"], v.pais)
			assert.Equal(t, responseMap["idioma"], v.idioma)
		}
		if v.statusCode == 400 || v.statusCode == 422 && v.errMessage != "" {
			assert.Equal(t, responseMap["message"], v.errMessage)
		}
	}
}

func TestActualizar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	configuracion.GetDatabaseInstance()
	_ = configuracion.RefreshPeliculaTable()
	pelicula, _ := configuracion.SendOnePelicula()

	samples := []struct {
		id         string
		inputJSON  string
		statusCode int
		status     string
		errMessage string
	}{
		{
			id:         strconv.Itoa(int(pelicula.Id)),
			inputJSON:  builder.NewPeliculaBuilder().ConNombre("Nombre update").ConDirector("Director update").ConIdioma("Idioma update").BuildString(),
			statusCode: 200,
			status:     "ok",
			errMessage: "",
		},
		{
			id: strconv.Itoa(int(pelicula.Id)),
			inputJSON: `{
							"nombre": "Nombre update",
							"director": "Director update",
							"escritor": "Escritor update",
							"pais": "Pais update",
							"idioma": "Idioma update",
							"lanzamiento": "2018"
						}`,
			statusCode: 422,
			errMessage: "JSON Invalido",
		},
	}

	for _, v := range samples {
		r := gin.Default()
		r.PUT(urlPeliculas+"/:id", contenedor.GetControladorPelicula().Actualizar)
		req, err := http.NewRequest(http.MethodPut, urlPeliculas+"/"+v.id, bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, responseMap["status"], v.status)
		}
		if v.statusCode == 400 || v.statusCode == 422 && v.errMessage != "" {
			assert.Equal(t, responseMap["message"], v.errMessage)
		}
	}
}

func TestListar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	configuracion.GetDatabaseInstance()
	_ = configuracion.RefreshPeliculaTable()
	_, err := configuracion.SendVariousPeliculas()

	r := gin.Default()
	r.GET(urlPeliculas, contenedor.GetControladorPelicula().Listar)

	req, err := http.NewRequest(http.MethodGet, urlPeliculas, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var peliculas []modelo.Pelicula

	err = json.Unmarshal(rr.Body.Bytes(), &peliculas)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(peliculas), 2)
}

func TestEliminar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	configuracion.GetDatabaseInstance()
	_ = configuracion.RefreshPeliculaTable()
	pelicula, _ := configuracion.SendOnePelicula()

	samples := []struct {
		id         string
		statusCode int
		status     string
		errMessage string
	}{
		{
			id:         strconv.Itoa(int(pelicula.Id)),
			statusCode: 200,
			status:     "ok",
			errMessage: "",
		},
	}

	for _, v := range samples {
		r := gin.Default()
		r.DELETE(urlPeliculas+"/:id", contenedor.GetControladorPelicula().Eliminar)
		req, err := http.NewRequest(http.MethodDelete, urlPeliculas+"/"+v.id, nil)
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, responseMap["status"], v.status)
		}
		if v.statusCode == 400 || v.statusCode == 404 || v.statusCode == 422 && v.errMessage != "" {
			assert.Equal(t, responseMap["message"], v.errMessage)
		}
	}
}
