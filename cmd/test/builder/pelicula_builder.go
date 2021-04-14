package builder

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"fmt"
)

type PeliculaBuilder struct {
	Id          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Director    string `json:"director"`
	Escritor    string `json:"escritor"`
	Pais        string `json:"pais"`
	Idioma      string `json:"idioma"`
	Lanzamiento int    `json:"lanzamiento"`
}

func (peliculaBuilder *PeliculaBuilder) ConId(id int64) *PeliculaBuilder {
	peliculaBuilder.Id = id
	return peliculaBuilder
}

func (peliculaBuilder *PeliculaBuilder) ConNombre(nombre string) *PeliculaBuilder {
	peliculaBuilder.Nombre = nombre
	return peliculaBuilder
}

func (peliculaBuilder *PeliculaBuilder) ConDirector(director string) *PeliculaBuilder {
	peliculaBuilder.Director = director
	return peliculaBuilder
}

func (peliculaBuilder *PeliculaBuilder) ConEscritor(escritor string) *PeliculaBuilder {
	peliculaBuilder.Escritor = escritor
	return peliculaBuilder
}

func (peliculaBuilder *PeliculaBuilder) ConPais(pais string) *PeliculaBuilder {
	peliculaBuilder.Pais = pais
	return peliculaBuilder
}

func (peliculaBuilder *PeliculaBuilder) ConIdioma(idioma string) *PeliculaBuilder {
	peliculaBuilder.Idioma = idioma
	return peliculaBuilder
}

func (peliculaBuilder *PeliculaBuilder) ConLanzamiento(lanzamiento int) *PeliculaBuilder {
	peliculaBuilder.Lanzamiento = lanzamiento
	return peliculaBuilder
}

func NewPeliculaBuilder() *PeliculaBuilder {
	return &PeliculaBuilder{
		Id:          1,
		Nombre:      "Oscar",
		Director:    "Alexander",
		Escritor:    "Ruiz",
		Pais:        "Colombia",
		Idioma:      "Español",
		Lanzamiento: 2021,
	}
}

func (peliculaBuilder *PeliculaBuilder) Build() modelo.Pelicula {
	return modelo.Pelicula{
		Id:          peliculaBuilder.Id,
		Nombre:      peliculaBuilder.Nombre,
		Director:    peliculaBuilder.Director,
		Escritor:    peliculaBuilder.Escritor,
		Pais:        peliculaBuilder.Pais,
		Idioma:      peliculaBuilder.Idioma,
		Lanzamiento: peliculaBuilder.Lanzamiento,
	}
}

func (peliculaBuilder *PeliculaBuilder) BuildString() string {
	return fmt.Sprintf(
		`{"nombre":"%s","director":"%s","escritor":"%s","pais":"%s","idioma":"%s","lanzamiento":%v}`,
		peliculaBuilder.Nombre,
		peliculaBuilder.Director,
		peliculaBuilder.Escritor,
		peliculaBuilder.Pais,
		peliculaBuilder.Idioma,
		peliculaBuilder.Lanzamiento,
	)
}
