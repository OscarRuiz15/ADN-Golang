package repositorio

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"database/sql"
	"fmt"
	"io/ioutil"
)

const (
	SQL_LISTAR_PELICULAS    = "listar_peliculas.sql"
	SQL_CREAR_PELICULA      = "crear_pelicula.sql"
	SQL_OBTENER_PELICULA    = "obtener_pelicula.sql"
	SQL_ACTUALIZAR_PELICULA = "actualizar_pelicula.sql"
	SQL_ELIMINAR_PELICULA   = "eliminar_pelicula.sql"
)

type RepositorioPeliculaSql struct {
	Db *sql.DB
}

func (repositorioPelicula *RepositorioPeliculaSql) Crear(pelicula *modelo.Pelicula) error {
	query := leerArchivoSql(SQL_CREAR_PELICULA)
	stmt, err := repositorioPelicula.Db.Prepare(query)
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
	query := leerArchivoSql(SQL_OBTENER_PELICULA)
	stmt, err := repositorioPelicula.Db.Prepare(query)
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
	query := leerArchivoSql(SQL_LISTAR_PELICULAS)
	rows, err := repositorioPelicula.Db.Query(query)
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
		fmt.Println("RepositorioSQL Listar -> No retorna peliculas la consulta", err)
		return nil, err
	}

	return peliculas, nil
}

func (repositorioPelicula *RepositorioPeliculaSql) Eliminar(id int64) error {
	query := leerArchivoSql(SQL_ELIMINAR_PELICULA)
	stmt, err := repositorioPelicula.Db.Prepare(query)
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
	query := leerArchivoSql(SQL_ACTUALIZAR_PELICULA)
	stmt, err := repositorioPelicula.Db.Prepare(query)
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

func leerArchivoSql(archivoSql string) string {
	ruta := fmt.Sprintf("./cmd/api/infraestructura/resources/sql/%s", archivoSql)
	file, err := ioutil.ReadFile(ruta)
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(file)
}
