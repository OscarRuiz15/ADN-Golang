package modelo

type Pelicula struct {
	Id          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Director    string `json:"director"`
	Escritor    string `json:"escritor"`
	Pais        string `json:"pais"`
	Idioma      string `json:"idioma"`
	Lanzamiento int    `json:"lanzamiento"`
}
