package presenter

import "github.com/ngochoangminh/golang_backend/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}