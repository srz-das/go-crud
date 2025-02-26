package injection

import (
	"github.com/google/wire"
	"github.com/teten-nugraha/golang-crud/handler"
	"github.com/teten-nugraha/golang-crud/repository"
	"github.com/teten-nugraha/golang-crud/service"
)

func initMonoAPI() handler.MonoAPI {
	wire.Build(
		repository.ProviderMonoRepository,
		service.ProviderMonoService,
		handler.ProviderMonoAPI,
	)

	return handler.MonoAPI{}
}
