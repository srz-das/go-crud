package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/golang-crud/domain"
)

type MonoRepositoryContract interface {
	SaveOrUpdate(mahasiawa domain.Mono) (domain.Mono, error)
	FindAll() [] domain.Mono
	FindByID(id string) domain.Mono
	FindByNim(nim string) domain.Mono
	DeleteMono(Mono domain.Mono) error
}

type MonoRepository struct {
	DB *gorm.DB
}

func ProviderMonoRepository(DB *gorm.DB) MonoRepository {
	return MonoRepository{DB: DB}
}

// implementation
func (m *MonoRepository) SaveOrUpdate(Mono domain.Mono) (domain.Mono, error) {
	if err := m.DB.Create(&Mono).Error; err != nil {
		return Mono, err
	}
	return Mono, nil
}

func (m *MonoRepository) FindAll() [] domain.Mono {
	var Monos []domain.Mono

	m.DB.Find(&Monos)

	return Monos
}

func(m *MonoRepository) FindByID(id string) domain.Mono {
	var Mono domain.Mono

	m.DB.Where("id =? ", id).Find(&Mono)

	return Mono
}


func(m *MonoRepository) FindByNim(nim string) domain.Mono {
	var Mono domain.Mono

	m.DB.Where("nim =? ", nim).Find(&Mono)

	return Mono
}

func (m *MonoRepository) DeleteMono(Mono domain.Mono) error {
	if err := m.DB.Delete(&Mono).Error; err!= nil  {
		return err
	}
	return nil
}