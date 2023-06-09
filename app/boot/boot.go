package boot

import (
	"context"
	"interview/app/common"
	"interview/app/contacts"
	"interview/app/controllers"
	"interview/app/providers/db"
	"interview/app/router"
)

func initMySqlClient() error {
	config := common.GetConfig().Database
	err := db.InitDB(db.Config{
		Name:     config.Name,
		Host:     config.Host,
		Port:     config.Port,
		Username: config.Username,
		Password: config.Password,
	})

	return err
}

func initProviders() error {
	err := initMySqlClient()
	if err != nil {
		return err
	}

	return nil
}

func initRoutes() error {
	configs := common.GetConfig()
	err := router.InitializeRouter(router.CoreConfigs{
		Name: configs.Core.Name,
		Host: configs.Core.Host,
		Port: configs.Core.Port,
	})
	return err
}

func initControllers() error {
	dbclient := db.GetDb()
	repo := db.NewRepo(dbclient)
	contactsCore := contacts.NewCore(repo)
	contactsService := contacts.NewService(contactsCore)

	controllers.NewAppController()
	controllers.NewContactsController(contactsService)

	return nil
}

func Init(ctx context.Context, env string) error {
	err := common.InitConfig(env)
	if err != nil {
		return err
	}

	err = initProviders()
	if err != nil {
		return err
	}

	err = initControllers()
	if err != nil {
		return err
	}

	err = initRoutes()
	if err != nil {
		return err
	}

	return nil
}
