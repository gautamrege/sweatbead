package user

import "github.com/gautamrege/sweatbead/profilemgr/db"

type updateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type createRequest struct {
	Name string `json:"name"`
}

type findByIDResponse struct {
	User db.User `json:"user"`
}

type listResponse struct {
	Users []db.User `json:"users"`
}

func (cr createRequest) Validate() (err error) {
	if cr.Name == "" {
		return errEmptyName
	}
	return
}

func (ur updateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.Name == "" {
		return errEmptyName
	}
	return
}
