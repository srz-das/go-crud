package routes

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/golang-crud/handler"
)

func MonoRoute(routes *echo.Echo, api handler.MonoAPI) {

	mhs := routes.Group("/Mono")
	{
		mhs.GET("/list", api.FindAll)
		mhs.POST("/save", api.SaveOrUpdate)
		mhs.GET("/findByNIM/:nim", api.FindByNIM)
		mhs.DELETE("/remove/:id", api.DeleteMono)
	}
}
