package mapper

import (
	"github.com/teten-nugraha/golang-crud/domain"
	"github.com/teten-nugraha/golang-crud/dto"
)

// From DTO to Domain
func ToMonoDomain(dto dto.MonoDTO) domain.Mono {
	return domain.Mono{
		Nim: dto.Nim,
		Nama: dto.Nama,
		Phone: dto.Phone,
		Alamat: dto.Alamat,
		Email: dto.Email,
	}
}

func ToMonoDomainList(dtos []dto.MonoDTO) []domain.Mono {
	Monos := make([]domain.Mono, len(dtos))

	for i, itm := range dtos {
		Monos[i] = ToMonoDomain(itm)
	}
	return Monos
}

// from domain to DTO
// From DTO to Domain
func ToMonoDto(Mono domain.Mono) dto.MonoDTO {
	return dto.MonoDTO{
		Nim: Mono.Nim,
		Nama: Mono.Nama,
		Phone: Mono.Phone,
		Alamat: Mono.Alamat,
		Email: Mono.Email,
	}
}

func ToMonoDtoList(Monos []domain.Mono) []dto.MonoDTO {
	dtos := make([]dto.MonoDTO, len(Monos))

	for i, itm := range Monos {
		dtos[i] = ToMonoDto(itm)
	}

	return dtos
}
