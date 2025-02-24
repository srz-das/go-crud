package injection

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/golang-crud/handler"
	"github.com/teten-nugraha/golang-crud/repository"
	"github.com/teten-nugraha/golang-crud/service"
)

func initMonoAPI(db *gorm.DB) handler.MonoAPI {
	

	return handler.MonoAPI{}
}