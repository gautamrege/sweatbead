package service

import (
	"github.com/gautamrege/packt/sweatbead/proto/usermgr"
	"github.com/gautamrege/packt/sweatbead/usermgr/db"
)

type Dependencies struct {
	DB db.Storer
	// define other service dependencies

	UserMgr usermgr.UserMgrClient
}
