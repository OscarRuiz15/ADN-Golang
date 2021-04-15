package repositorio

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"ADN_Golang/cmd/api/infraestructura/configuracion"
	"database/sql"
	"errors"
	"fmt"
)

const (
	QUERY_CREAR_PELICULA      = "INSERT INTO pelicula (nombre, director, escritor, pais, idioma, lanzamiento) VALUES (?, ?, ?, ?, ?, ?)"
	QUERY_LISTAR_PELICULAS    = "SELECT id, nombre, director, escritor, pais, idioma, lanzamiento FROM pelicula"
	QUERY_EXISTE_PELICULA     = "SELECT id FROM pelicula WHERE nombre LIKE ?"
	QUERY_ACTUALIZAR_PELICULA = "UPDATE pelicula SET nombre=?, director=?, escritor=?, pais=?, idioma=?, lanzamiento=? WHERE id = ?;"
	QUERY_OBTENER_PELICULA    = "SELECT id, nombre, director, escritor, pais, idioma, lanzamiento FROM pelicula WHERE id = ?"
	QUERY_ELIMINAR_PELICULA   = "DELETE FROM pelicula WHERE id = ?"
)

type RepositorioPeliculaSql struct {
	Db *sql.DB
}

// Begin & Exec
func (repositorioPelicula *RepositorioPeliculaSql) Crear(pelicula *modelo.Pelicula) (err error) {
	var tx *sql.Tx

	defer func() {
		configuracion.CloseConnections(err, tx, nil, nil)
	}()

	tx, err = repositorioPelicula.Db.Begin()
	if err != nil {
		errMsg := fmt.Sprintf("RepositorioSQL Crear -> Ocurrió un error al guardar la pelicula con nombre %s", pelicula.Nombre)
		//log.Println(errMsg, err)
		return errors.New(errMsg)
	}

	result, err := repositorioPelicula.Db.Exec(QUERY_CREAR_PELICULA,
		pelicula.Nombre,
		pelicula.Director,
		pelicula.Escritor,
		pelicula.Pais,
		pelicula.Idioma,
		pelicula.Lanzamiento)
	if err != nil {
		//log.Println("RepositorioSQL Crear -> El tipo de parámetro no es correcto", err)
		return errors.New("RepositorioSQL Crear -> El tipo de parámetro no es correcto")
	}

	pelicula.Id, _ = result.LastInsertId()
	return err
}

// QueryRow
func (repositorioPelicula *RepositorioPeliculaSql) Obtener(id int64) (pelicula modelo.Pelicula, err error) {
	row := repositorioPelicula.Db.QueryRow(QUERY_OBTENER_PELICULA, id)
	err = row.Scan(&pelicula.Id, &pelicula.Nombre, &pelicula.Director, &pelicula.Escritor, &pelicula.Pais, &pelicula.Idioma, &pelicula.Lanzamiento)
	if err != nil {
		//log.Println("RepositorioSQL Obtener -> Error al ejecutar instancia SQL", err)
		return modelo.Pelicula{}, errors.New("RepositorioSQL Obtener -> Error al ejecutar instancia SQL")
	}
	return pelicula, err
}

// Query
func (repositorioPelicula *RepositorioPeliculaSql) Listar() ([]modelo.Pelicula, error) {
	rows, err := repositorioPelicula.Db.Query(QUERY_LISTAR_PELICULAS)
	if err != nil {
		//log.Println("RepositorioSQL Listar -> Error de sintaxis de consulta", err)
		return nil, errors.New("RepositorioSQL Listar -> Error de sintaxis de consulta")
	}

	defer rows.Close()

	peliculas := make([]modelo.Pelicula, 0)
	for rows.Next() {
		var pelicula modelo.Pelicula
		if err := rows.Scan(
			&pelicula.Id,
			&pelicula.Nombre,
			&pelicula.Director,
			&pelicula.Escritor,
			&pelicula.Pais,
			&pelicula.Idioma,
			&pelicula.Lanzamiento,
		); err != nil {
			//log.Println("RepositorioSQL Listar -> Error al escanear la información", err)
			return nil, errors.New("RepositorioSQL Listar -> error al escanear la información")
		}
		peliculas = append(peliculas, pelicula)
	}

	if len(peliculas) == 0 {
		//log.Println("RepositorioSQL Listar -> No retorna peliculas la consulta", err)
		return nil, errors.New("RepositorioSQL Listar -> No retorna peliculas la consulta")
	}

	return peliculas, nil
}

// Begin & Exec
func (repositorioPelicula *RepositorioPeliculaSql) Eliminar(id int64) (err error) {
	var tx *sql.Tx

	defer func() {
		configuracion.CloseConnections(err, tx, nil, nil)
	}()

	tx, err = repositorioPelicula.Db.Begin()
	if err != nil {
		errMsg := fmt.Sprintf("RepositorioSQL Eliminar -> Ocurrió un error al eliminar la pelicula con id %v", id)
		//log.Println(errMsg, err)
		return errors.New(errMsg)
	}

	_, err = repositorioPelicula.Db.Exec(QUERY_ELIMINAR_PELICULA, id)
	if err != nil {
		//log.Println("RepositorioSQL Eliminar -> El tipo de parámetro no es correcto", err)
		return errors.New("RepositorioSQL Eliminar -> El tipo de parámetro no es correcto")
	}

	return err
}

// Begin & Exec
func (repositorioPelicula *RepositorioPeliculaSql) Actualizar(id int64, pelicula modelo.Pelicula) (err error) {
	var tx *sql.Tx

	defer func() {
		configuracion.CloseConnections(err, tx, nil, nil)
	}()

	tx, err = repositorioPelicula.Db.Begin()
	if err != nil {
		errMsg := fmt.Sprintf("RepositorioSQL Actualizar -> Ocurrió un error al actualizar la pelicula con id %v", id)
		//log.Println(errMsg, err)
		return errors.New(errMsg)
	}

	_, err = repositorioPelicula.Db.Exec(QUERY_ACTUALIZAR_PELICULA,
		pelicula.Nombre,
		pelicula.Director,
		pelicula.Escritor,
		pelicula.Pais,
		pelicula.Idioma,
		pelicula.Lanzamiento,
		id)
	if err != nil {
		//log.Println("RepositorioSQL Actualizar -> El tipo de parámetro no es correcto", err)
		return errors.New("RepositorioSQL Actualizar -> El tipo de parámetro no es correcto")
	}

	return err
}

// QueryRow
func (repositorioPelicula *RepositorioPeliculaSql) Existe(nombre string) (int64, bool) {
	row := repositorioPelicula.Db.QueryRow(QUERY_EXISTE_PELICULA, "%"+nombre+"%")
	var id int64
	err := row.Scan(&id)
	if err != nil {
		//log.Println("RepositorioSQL Existe -> Error al ejecutar instancia SQL", err)
		return 0, false
	}
	if id > 0 {
		return id, true
	}
	return 0, false
}
