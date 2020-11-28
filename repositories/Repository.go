package repositories

import "github.com/jackc/pgx/v4/pgxpool"

type repo struct {
	DB *pgxpool.Pool
}

//Repository repository object
type Repository struct {
	CredentialRepository  CredentialRepository
	UserProfileRepository UserProfileRepository
}

//NewRepository instance of Repository
func NewRepository(db *pgxpool.Pool) *Repository {
	credentialRepository := NewCredentialRepository(db)
	userProfileRepository := NewUserProfileRepository(db)
	return &Repository{
		CredentialRepository:  credentialRepository,
		UserProfileRepository: userProfileRepository,
	}
}
