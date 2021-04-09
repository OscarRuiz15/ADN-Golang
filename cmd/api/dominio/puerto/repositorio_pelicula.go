package puerto

import "ADN_Golang/cmd/api/dominio/modelo"

type RepositorioPelicula interface {
	//Crear pelicula, recibe un modelo tipo pelicula
	Crear(pelicula *modelo.Pelicula) error

	//Actualizar pelicula, recibe un id y un modelo tipo pelicula
	Actualizar(id int64, pelicula modelo.Pelicula) error

	//Eliminar pelicula, recibe un id
	Eliminar(id int64) error

	//Obtener pelicula, recibe un id y retorna un modelo tipo pelicula
	Obtener(id int64) (modelo.Pelicula, error)

	//Obtener todas las peliculas
	Listar() ([]modelo.Pelicula, error)
}
