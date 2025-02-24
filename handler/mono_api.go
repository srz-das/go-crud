package handler

import (
	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
	"github.com/teten-nugraha/golang-crud/dto"
	_ "github.com/teten-nugraha/golang-crud/dto"
	"github.com/teten-nugraha/golang-crud/service"
	"net/http"
)

type MonoAPI struct {
	MonoService service.MonoService
}

func ProviderMonoAPI(k service.MonoService) MonoAPI {
	return MonoAPI{MonoService: k}
}

// implementasi
func (m *MonoAPI) FindAll(e echo.Context) error {

	Monos := m.MonoService.FindAll()

	if len(Monos) == 0 {
		return SuccessResponse(e, http.StatusNoContent, Monos)
	}

	return SuccessResponse(e, http.StatusOK, Monos)
}

func (m *MonoAPI) SaveOrUpdate(e echo.Context) error {
	var newDto dto.MonoDTO

	newDto.Nim = e.FormValue("Nim")
	newDto.Nama = e.FormValue("Nama")
	newDto.Phone = e.FormValue("Phone")
	newDto.Alamat = e.FormValue("Alamat")
	newDto.Email = e.FormValue("Email")

	res, err := m.MonoService.SaveOrUpdate(newDto)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

func (m *MonoAPI) FindByNIM(e echo.Context) error {
	nim := e.Param("nim")

	Mono := m.MonoService.FindByNim(nim)

	return SuccessResponse(e, http.StatusOK, Mono)
}

func (m *MonoAPI) DeleteMono(e echo.Context) error {
	id := e.Param("id")

	err := m.MonoService.DeleteMono(id)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, "Delete Success")
}
