package repository

import "bitbucket.org/janpavtel/site/internal/models"

type DatabaseRepo interface{
	LoadAllUsers() ([]models.User, error)

	LoadUserById(id int) (models.User, error)

	AddUser(user models.User) (int, error)
}