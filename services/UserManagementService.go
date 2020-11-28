package services

import (
	"context"
	"errors"
	"fiberauthenticationjwt/entities"
	"fiberauthenticationjwt/repositories"
	"log"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

//UserManagementService service related to user
type UserManagementService interface {
	CreateUserProfile(ctx context.Context, userProfile *entities.UserProfile) (*entities.UserProfile, error)
	CreateUserCredential(ctx context.Context, credential *entities.Credential) (*entities.Credential, error)
	AuthenticateUser(ctx context.Context, credential *entities.Credential) (*entities.AuthObject, error)
	GetUserInfo(ctx context.Context, profileID int64) (*entities.UserProfile, error)
	ChangePassword(ctx context.Context, credID int64, newPassword string) (*entities.Credential, error)
	ChangeProfileName(ctx context.Context, profileID int64, newProfileName string) (*entities.UserProfile, error)
}

//NewUserManagementService instance of UserManagementService
func NewUserManagementService(repo *repositories.Repository) UserManagementService {
	return &service{
		repo: repo,
	}
}

func (service *service) CreateUserProfile(ctx context.Context, userProfile *entities.UserProfile) (*entities.UserProfile, error) {
	_userProfile, err := service.repo.UserProfileRepository.CreateUserProfile(ctx, userProfile)
	if err != nil {
		return nil, err
	}
	return _userProfile, nil
}

func (service *service) CreateUserCredential(ctx context.Context, credential *entities.Credential) (*entities.Credential, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credential.Password), bcrypt.MinCost)
	if err != nil {
		return nil, errors.New("failed hashing password")
	}

	credential.Password = string(hashedPassword)
	credential.IsEnabled = true

	_credential, err := service.repo.CredentialRepository.CreateCredential(ctx, credential)

	if err != nil {
		return nil, err
	}

	return _credential, nil
}

func (service *service) AuthenticateUser(ctx context.Context, credential *entities.Credential) (*entities.AuthObject, error) {
	log.Println("Username = ", credential.Username)
	_credential, err := service.repo.CredentialRepository.FindByUsername(ctx, credential.Username)
	if err != nil {
		return nil, errors.New("Invalid username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(_credential.Password), []byte(credential.Password))
	if err != nil {
		return nil, errors.New("Invalid password")
	}

	appSecret, ok := viper.Get("APP_SECRET_KEY").(string)
	if !ok {
		return nil, errors.New("Cannot fiund secret key")
	}

	jwtID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("failed generating uuid")
	}

	userProfile, err := service.repo.UserProfileRepository.FindUserProfileByID(ctx, _credential.UserProfileID)
	if err != nil {
		return nil, errors.New("Profile not exist")
	}

	jwtClaims := &entities.JwtApplicationUserClaim{
		Username:     _credential.Username,
		CredentialID: _credential.ID,
		IsEnabled:    _credential.IsEnabled,
		UserRole:     _credential.Role,
		ProfileID:    _credential.UserProfileID,
		ProfileName:  userProfile.ProfileName,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 5).Unix(),
			Subject:   _credential.Username,
			Issuer:    "improvalab.id",
			Id:        jwtID.String(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	token, err := jwtToken.SignedString([]byte(appSecret))
	if err != nil {
		return nil, errors.New("Cannot signing token")
	}

	authObj := &entities.AuthObject{
		ProfileID: userProfile.ID,
		Token:     token,
	}

	return authObj, nil

}

func (service *service) GetUserInfo(ctx context.Context, profileID int64) (*entities.UserProfile, error) {
	userProfile, err := service.repo.UserProfileRepository.FindUserProfileByID(ctx, profileID)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}

func (service *service) ChangePassword(ctx context.Context, credID int64, newPassword string) (*entities.Credential, error) {
	credential, err := service.repo.CredentialRepository.FindCredentialByID(ctx, credID)
	if err != nil {
		return nil, errors.New("Credential not found")
	}

	if credential.IsEnabled == false {
		return nil, errors.New("User is disabled")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credential.Password), bcrypt.MaxCost)

	if err != nil {
		return nil, errors.New("Cannot hash password")
	}

	updatedCredential, err := service.repo.CredentialRepository.ChangePassword(ctx, credential.ID, string(hashedPassword))
	if err != nil {
		return nil, errors.New("cannot update password")
	}
	updatedCredential.Password = ""
	return updatedCredential, nil
}

func (service *service) ChangeProfileName(ctx context.Context, profileID int64, newProfileName string) (*entities.UserProfile, error) {
	userProfile, err := service.repo.UserProfileRepository.ChangeProfileName(ctx, profileID, newProfileName)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}
