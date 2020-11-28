package services

import "fiberauthenticationjwt/repositories"

type service struct {
	repo *repositories.Repository
}

//Service object holds all services
type Service struct {
	UserManagementService UserManagementService
}

//NewService instance of Service
func NewService(repo *repositories.Repository) *Service {
	userManagementService := NewUserManagementService(repo)
	return &Service{
		UserManagementService: userManagementService,
	}
}
