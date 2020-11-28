package repositories

import (
	"context"
	"fiberauthenticationjwt/entities"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

//UserProfileRepository data repository of user_profiles
type UserProfileRepository interface {
	FindUserProfileByID(ctx context.Context, id int64) (*entities.UserProfile, error)
	CreateUserProfile(ctx context.Context, userProfile *entities.UserProfile) (*entities.UserProfile, error)
	ChangeProfileName(ctx context.Context, id int64, profileName string) (*entities.UserProfile, error)
}

//NewUserProfileRepository instance of UserProfileRepository
func NewUserProfileRepository(db *pgxpool.Pool) UserProfileRepository {
	return &repo{
		DB: db,
	}
}

func (repo *repo) FindUserProfileByID(ctx context.Context, id int64) (*entities.UserProfile, error) {
	query := `
		SELECT id, profile_name, phone_number, email, gender, created_at, updated_at, deleted_at
		FROM user_profiles WHERE id = $1 AND deleted_at IS NULL
	`
	var (
		_ID          pgtype.Int8
		_ProfileName pgtype.Varchar
		_PhoneNumber pgtype.Varchar
		_Email       pgtype.Varchar
		_Gender      pgtype.Varchar
		_CreatedAt   pgtype.Timestamp
		_UpdatedAt   pgtype.Timestamp
		_DeletedAt   pgtype.Timestamp
	)
	err := repo.DB.QueryRow(ctx, query, id).
		Scan(&_ID, &_ProfileName, &_PhoneNumber, &_Email, &_Gender, &_CreatedAt, &_UpdatedAt, &_DeletedAt)
	if err != nil {
		return nil, err
	}

	userProfile := &entities.UserProfile{
		ID:          _ID.Int,
		ProfileName: _ProfileName.String,
		PhoneNumber: _PhoneNumber.String,
		Email:       _Email.String,
		Gender:      _Gender.String,
		CreatedAt:   &_CreatedAt.Time,
		UpdatedAt:   &_UpdatedAt.Time,
		DeletedAt:   &_DeletedAt.Time,
	}

	return userProfile, nil
}

func (repo *repo) CreateUserProfile(ctx context.Context, userProfile *entities.UserProfile) (*entities.UserProfile, error) {
	query := `
		INSERT INTO user_profiles(profile_name, phone_number, email, gender)
		VALUES($1, $2, $3, $4)
		RETURNING id
	`
	var _ID pgtype.Int8
	err := repo.DB.QueryRow(ctx, query, userProfile.ProfileName, userProfile.PhoneNumber, userProfile.Email, userProfile.Gender).
		Scan(&_ID)
	if err != nil {
		return nil, err
	}

	_userProfile, err := repo.FindUserProfileByID(ctx, _ID.Int)
	if err != nil {
		return nil, err
	}

	return _userProfile, nil
}
func (repo *repo) ChangeProfileName(ctx context.Context, id int64, profileName string) (*entities.UserProfile, error) {
	query := `
		UPDATE user_profiles SET profile_name=$1 WHERE id=$2 AND deleted_at IS NULL
		RETURNING id
	`
	var _ID pgtype.Int8
	err := repo.DB.QueryRow(ctx, query, profileName, id).
		Scan(&_ID)
	if err != nil {
		return nil, err
	}

	userProfile, err := repo.FindUserProfileByID(ctx, _ID.Int)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}
