package server

import (
	"github.com/gautamrege/sweatbead/eventmgr/app"
	"github.com/gautamrege/sweatbead/eventmgr/db"
	"github.com/gautamrege/sweatbead/eventmgr/user"
)

type dependencies struct {
	UserService user.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	userService := user.NewService(dbStore, logger)

	return dependencies{
		UserService: userService,
	}, nil
}
