package container

import (
	"github.com/google/wire"
	"github.com/li1553770945/personal-feedback-service/biz/infra/config"
	"github.com/li1553770945/personal-feedback-service/biz/infra/database"
	"github.com/li1553770945/personal-feedback-service/biz/internal/repo"
	"github.com/li1553770945/personal-feedback-service/biz/internal/service"
)

func GetContainer(cfg *config.Config) *Container {
	panic(wire.Build(

		//infra

		//repo
		repo.NewRepository,
		database.NewDatabase,

		//service
		service.NewFeedbackService,

		NewContainer,
	))
}
