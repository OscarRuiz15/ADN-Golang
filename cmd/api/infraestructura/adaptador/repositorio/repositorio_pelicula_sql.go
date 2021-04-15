package repositorio

import (
	"ADN_Golang/cmd/api/dominio/modelo"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type RepositorioPeliculaSql struct {
	Db *gorm.DB
}

func (repositorioPelicula *RepositorioPeliculaSql) Crear(pelicula *modelo.Pelicula) (err error) {
	if repositorioPelicula.Db.Create(pelicula).Error != nil {
		errMsg := fmt.Sprintf("RepositorioSQL Crear -> Ocurrió un error al guardar la pelicula con nombre %s", pelicula.Nombre)
		return errors.New(errMsg)
	}
	return err
}

func (repositorioPelicula *RepositorioPeliculaSql) Obtener(id int64) (pelicula modelo.Pelicula, err error) {
	if err = repositorioPelicula.Db.Where("id = ?", id).First(&pelicula).Error; err != nil {
		return modelo.Pelicula{}, errors.New("RepositorioSQL Obtener -> No se encuentra el registro")
	}
	return pelicula, err
}

func (repositorioPelicula *RepositorioPeliculaSql) Listar() ([]modelo.Pelicula, error) {
	var peliculas []modelo.Pelicula
	if repositorioPelicula.Db.Find(&peliculas).Error != nil {
		return nil, errors.New("RepositorioSQL Listar -> Error al realizar la consulta")
	}
	if len(peliculas) == 0 {
		return nil, errors.New("RepositorioSQL Listar -> No se encontraron peliculas")
	}
	return peliculas, nil
}

func (repositorioPelicula *RepositorioPeliculaSql) Eliminar(id int64) (err error) {
	var pelicula modelo.Pelicula
	if err = repositorioPelicula.Db.Where("id = ?", id).First(&pelicula).Error; err != nil {
		errMsg := fmt.Sprintf("RepositorioSQL Eliminar -> No se encuentra el registro a eliminar con id %v", id)
		return errors.New(errMsg)
	}
	repositorioPelicula.Db.Delete(pelicula)
	return err
}

func (repositorioPelicula *RepositorioPeliculaSql) Actualizar(id int64, pelicula modelo.Pelicula) (err error) {
	var actual modelo.Pelicula
	if repositorioPelicula.Db.Where("id = ?", id).First(&actual).RecordNotFound() {
		errMsg := fmt.Sprintf("RepositorioSQL Actualizar -> No se encuentra el registro a actualizar con id %v", id)
		return errors.New(errMsg)
	}
	if repositorioPelicula.Db.Model(&actual).Update(pelicula).Error != nil {
		return errors.New("RepositorioSQL Actualizar -> El tipo de parámetro no es correcto")
	}
	return err
}

func (repositorioPelicula *RepositorioPeliculaSql) Existe(nombre string) (int64, bool) {
	var pelicula modelo.Pelicula
	if repositorioPelicula.Db.Where("nombre = ?", nombre).Find(&pelicula).Error != nil {
		return 0, false
	}
	return pelicula.Id, true
}
