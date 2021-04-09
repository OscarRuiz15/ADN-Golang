package marshaller

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"encoding/json"
)

type PeliculaJson struct {
	Id          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Director    string `json:"director"`
	Escritor    string `json:"escritor"`
	Pais        string `json:"pais"`
	Idioma      string `json:"idioma"`
	Lanzamiento int    `json:"lanzamiento"`
}

func Marshall(pelocula modelo.Pelicula) interface{} {
	movieJson, _ := json.Marshal(pelocula)
	var peliculaP PeliculaJson
	_ = json.Unmarshal(movieJson, &peliculaP)
	return peliculaP
}

func MarshallArray(peliculas []modelo.Pelicula) []interface{} {
	result := make([]interface{}, len(peliculas))
	for index, movie := range peliculas {
		result[index] = Marshall(movie)
	}
	return result
}
