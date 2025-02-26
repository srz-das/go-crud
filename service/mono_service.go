package service

import (
	"errors"
	"github.com/teten-nugraha/golang-crud/domain"
	"github.com/teten-nugraha/golang-crud/dto"
	"github.com/teten-nugraha/golang-crud/mapper"
	"github.com/teten-nugraha/golang-crud/repository"
)
// changed name from MonoRepositoryContract to MonoServiceContract
type MonoServiceContract interface {
	SaveOrUpdate(dto dto.MonoDTO) (dto.MonoDTO, error)
	FindAll() ([] dto.MonoDTO, error)
	FindByNim(nim string) dto.MonoDTO
	DeleteMono(id string) error
}

type MonoService struct {
	MonoRepository repository.MonoRepository
}

func ProviderMonoService(m repository.MonoRepository) MonoServiceContract {
	return &MonoService{
		MonoRepository: m,
	}
}

// implementation
func (m *MonoService) SaveOrUpdate(dto dto.MonoDTO) (dto.MonoDTO, error) {

	MonoEntity := mapper.ToMonoDomain(dto)

	Mono, err := m.MonoRepository.SaveOrUpdate(MonoEntity)

	return mapper.ToMonoDto(Mono), err
}

func (m *MonoService) FindAll() ([]dto.MonoDTO, error) {

	datas, err := m.MonoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return mapper.ToMonoDtoList(datas), nil
}

func (m *MonoService) FindByNim(nim string) dto.MonoDTO {
	Mono := m.MonoRepository.FindByNim(nim)

	return mapper.ToMonoDto(Mono)
}

func (m *MonoService) DeleteMono(id string) error {

	Mono, err := m.MonoRepository.FindByID(id)
	if err != nil {
		return err
	}

	if(Mono == (domain.Mono{})) {
		return errors.New("mono tidak ada")
	}

	err = m.MonoRepository.DeleteMono(Mono)
	if err != nil {
		return err
	}

	return nil
}