package repositories

import (
	"context"
	"fiberauthenticationjwt/entities"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

//CredentialRepository data repository of credentials
type CredentialRepository interface {
	FindByUsername(ctx context.Context, username string) (*entities.Credential, error)
	FindCredentialByID(ctx context.Context, id int64) (*entities.Credential, error)
	CreateCredential(ctx context.Context, credential *entities.Credential) (*entities.Credential, error)
	ChangePassword(ctx context.Context, id int64, newPassword string) (*entities.Credential, error)
}

//NewCredentialRepository instance of CredentialRepository
func NewCredentialRepository(db *pgxpool.Pool) CredentialRepository {
	return &repo{
		DB: db,
	}
}

func (repo *repo) FindCredentialByID(ctx context.Context, id int64) (*entities.Credential, error) {
	query := `
		SELECT id, username, is_enabled, user_profile_id, created_at
		FROM credentials
		WHERE id = $1 AND deleted_at IS NULL
	`
	var (
		_ID            pgtype.Int8
		_Username      pgtype.Varchar
		_IsEnabled     pgtype.Bool
		_UserProfileID pgtype.Int8
		_CreatedAt     pgtype.Timestamp
	)
	err := repo.DB.QueryRow(ctx, query, id).
		Scan(&_ID, &_Username, &_IsEnabled, &_UserProfileID, &_CreatedAt)

	if err != nil {
		return nil, err
	}

	credential := &entities.Credential{
		ID:            _ID.Int,
		Username:      _Username.String,
		UserProfileID: _UserProfileID.Int,
		CreatedAt:     &_CreatedAt.Time,
	}

	return credential, nil
}

func (repo *repo) CreateCredential(ctx context.Context, credential *entities.Credential) (*entities.Credential, error) {
	query := `
		INSERT INTO credentials(username, password, is_enabled, user_profile_id)
		VALUES($1, $2, $3, $4)
		RETURNING id
	`
	var _ID pgtype.Int8
	err := repo.DB.QueryRow(ctx, query, credential.Username, credential.Password, credential.IsEnabled, credential.UserProfileID).
		Scan(&_ID)
	if err != nil {
		return nil, err
	}

	_credential, err := repo.FindCredentialByID(ctx, _ID.Int)
	if err != nil {
		return nil, err
	}

	return _credential, nil
}

func (repo *repo) FindByUsername(ctx context.Context, username string) (*entities.Credential, error) {
	query := `
		SELECT id, username, password, is_enabled, user_profile_id, created_at, updated_at, deleted_at
		FROM credentials
		WHERE username = $1 AND deleted_at IS NULL
	`
	var (
		_ID            pgtype.Int8
		_Username      pgtype.Varchar
		_Password      pgtype.Varchar
		_IsEnabled     pgtype.Bool
		_UserProfileID pgtype.Int8
		_CreatedAt     pgtype.Timestamp
		_UpdatedAt     pgtype.Timestamp
		_DeletedAt     pgtype.Timestamp
	)
	err := repo.DB.QueryRow(ctx, query, username).
		Scan(&_ID, &_Username, &_Password, &_IsEnabled, &_UserProfileID, &_CreatedAt, &_UpdatedAt, &_DeletedAt)

	if err != nil {
		return nil, err
	}

	credential := &entities.Credential{
		ID:            _ID.Int,
		Username:      _Username.String,
		Password:      _Password.String,
		UserProfileID: _UserProfileID.Int,
		CreatedAt:     &_CreatedAt.Time,
	}

	return credential, nil
}

func (repo *repo) ChangePassword(ctx context.Context, id int64, newPassword string) (*entities.Credential, error) {
	query := `
		UPDATE credentials SET password=$1 WHERE id=$2 AND deleted_at IS NULL
		RETURNING id
	`
	var _ID pgtype.Int8
	err := repo.DB.QueryRow(ctx, query, newPassword, id).
		Scan(&_ID)
	if err != nil {
		return nil, err
	}

	credential, err := repo.FindCredentialByID(ctx, _ID.Int)
	if err != nil {
		return nil, err
	}

	return credential, nil
}
