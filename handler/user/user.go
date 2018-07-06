package user

import (
	"apiserver/model"
	"apiserver/package/errors"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

func (r *CreateRequest) checkParams() error {
	if r.Username == "" {
		return errors.New(errors.ValidationError, nil).Add("Username is empty")
	}
	if r.Password == "" {
		return errors.New(errors.ValidationError, nil).Add("Password is empty")
	}

	return nil
}
