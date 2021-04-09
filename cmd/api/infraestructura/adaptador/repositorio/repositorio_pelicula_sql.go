package repositorio

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"database/sql"
	"fmt"
)

const (
	queryInsertarPelicula     = "INSERT INTO pelicula (nombre, director, escritor, pais, idioma, lanzamiento) VALUES (?,?,?,?,?,?)"
	queryObtenerPeliculaPorId = "SELECT id, nombre, director, escritor, pais, idioma, lanzamiento FROM pelicula WHERE id = ?"
	queryObtenerPeliculas     = "SELECT id, nombre, director, escritor, pais, idioma, lanzamiento FROM pelicula "
	queryEliminarPeliculas    = "DELETE FROM pelicula WHERE id = ?"
	queryActualizarPelicula   = "UPDATE pelicula SET nombre=?, director=?, escritor=?, pais=?, idioma=?, lanzamiento=? WHERE id=?;"
)

type RepositorioPeliculaSql struct {
	Db *sql.DB
}

func (repositorioPelicula *RepositorioPeliculaSql) Crear(pelicula *modelo.Pelicula) error {
	stmt, err := repositorioPelicula.Db.Prepare(queryInsertarPelicula)
	if err != nil {
		fmt.Println("RepositorioSQL Crear -> Error al preparar instancia SQL", err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(pelicula.Nombre, pelicula.Director, pelicula.Escritor, pelicula.Pais, pelicula.Idioma, pelicula.Lanzamiento)
	if err != nil {
		fmt.Println("RepositorioSQL Crear -> Error al ejecutar SQL", err)
		return err
	}

	pelicula.Id, err = result.LastInsertId()
	if err != nil {
		fmt.Println("RepositorioSQL Crear -> Error al obtener ultimo id", err)
		return err
	}

	return nil
}

func (repositorioPelicula *RepositorioPeliculaSql) Obtener(id int64) (modelo.Pelicula, error) {
	stmt, err := repositorioPelicula.Db.Prepare(queryObtenerPeliculaPorId)
	if err != nil {
		fmt.Println("RepositorioSQL Obtener -> Error al preparar instancia SQL", err)
		return modelo.Pelicula{}, err
	}

	defer stmt.Close()

	var pelicula modelo.Pelicula
	result := stmt.QueryRow(id)
	err = result.Scan(&pelicula.Id, &pelicula.Nombre, &pelicula.Director, &pelicula.Escritor, &pelicula.Pais, &pelicula.Idioma, &pelicula.Lanzamiento)
	if err != nil {
		fmt.Println("RepositorioSQL Obtener -> Error al ejecutar instancia SQL", err)
		return modelo.Pelicula{}, err
	}

	return pelicula, nil
}

func (repositorioPelicula *RepositorioPeliculaSql) Listar() ([]modelo.Pelicula, error) {
	rows, err := repositorioPelicula.Db.Query(queryObtenerPeliculas)
	if err != nil {
		fmt.Println("RepositorioSQL Listar -> Error al preparar instancia SQL", err)
		return nil, err
	}

	defer rows.Close()

	peliculas := make([]modelo.Pelicula, 0)
	for rows.Next() {
		var pelicula modelo.Pelicula
		if err := rows.Scan(&pelicula.Id, &pelicula.Nombre, &pelicula.Director, &pelicula.Escritor, &pelicula.Pais, &pelicula.Idioma, &pelicula.Lanzamiento); err != nil {
			fmt.Println("RepositorioSQL Listar -> Error al recorrer filas", err)
			return nil, err
		}
		peliculas = append(peliculas, pelicula)
	}

	if len(peliculas) == 0 {
		fmt.Println("RepositorioSQL Listar -> No hay usuarios", err)
		return nil, err
	}

	return peliculas, nil
}

func (repositorioPelicula *RepositorioPeliculaSql) Eliminar(id int64) error {
	stmt, err := repositorioPelicula.Db.Prepare(queryEliminarPeliculas)
	if err != nil {
		fmt.Println("RepositorioSQL Eliminar -> Error al preparar instancia SQL", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Println("RepositorioSQL Eliminar -> Error al ejecutar SQL", err)
		return err
	}

	return nil
}

func (repositorioPelicula *RepositorioPeliculaSql) Actualizar(id int64, pelicula modelo.Pelicula) error {
	stmt, err := repositorioPelicula.Db.Prepare(queryActualizarPelicula)
	if err != nil {
		fmt.Println("RepositorioSQL Actualizar -> Error al preparar instancia SQL", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(pelicula.Nombre, pelicula.Director, pelicula.Escritor, pelicula.Pais, pelicula.Idioma, pelicula.Lanzamiento, id)
	if err != nil {
		fmt.Println("RepositorioSQL Actualizar -> Error al ejecutar SQL", err)
		return err
	}

	return nil
}
