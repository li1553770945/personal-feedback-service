// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/li1553770945/personal-feedback-service/biz/infra/config"
	"github.com/li1553770945/personal-feedback-service/biz/infra/database"
	"github.com/li1553770945/personal-feedback-service/biz/internal/repo"
	"github.com/li1553770945/personal-feedback-service/biz/internal/service"
)

// Injectors from wire.go:

func GetContainer(env string) *Container {
	configConfig := config.InitConfig(env)
	db := database.NewDatabase(configConfig)
	iRepository := repo.NewRepository(db)
	iProjectService := project.NewProjectService(iRepository)
	container := NewContainer(configConfig, iProjectService)
	return container
}
