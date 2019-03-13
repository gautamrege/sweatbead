package server

import (
	"github.com/gautamrege/sweatbead/profilemgr/app"
	"github.com/gautamrege/sweatbead/profilemgr/db"
	"github.com/gautamrege/sweatbead/profilemgr/user"
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
