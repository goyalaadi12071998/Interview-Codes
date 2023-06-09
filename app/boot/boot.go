package boot

import (
	"context"
	"interview/app/common"
)

func initProviders() error {
	return nil
}

func initRoutes() error {
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

	err = initRoutes()
	if err != nil {
		return err
	}

	return nil
}
