package services

import (
	"github.com/jorgeflorescarlos/golang-microservices/mvc/domain"
	"github.com/jorgeflorescarlos/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
