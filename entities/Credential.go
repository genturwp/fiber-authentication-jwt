package entities

import (
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
)

//Credential object of credentials
type Credential struct {
	ID            int64      `json:"id,omitempty"`
	Username      string     `json:"username,omitempty"`
	Password      string     `json:"password,omitempty"`
	IsEnabled     bool       `json:"isEnabled"`
	UserProfileID int64      `json:"userProfileId,omitempty"`
	Role          string     `json:"role,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty"`
}

//JwtApplicationUserClaim object jwtclaim
type JwtApplicationUserClaim struct {
	Username     string `json:"username,omitempty"`
	CredentialID int64  `json:"credentialId,omitempty"`
	IsEnabled    bool   `json:"isEnabled,omitempty"`
	UserRole     string `json:"userRole,omitempty"`
	ProfileID    int64  `json:"profileId,omitempty"`
	ProfileName  string `json:"profileName,omitempty"`
	jwt.StandardClaims
}

//AuthObject authentication object
type AuthObject struct {
	ProfileID int64  `json:"profileId,omitempty"`
	Token     string `json:"token,omitempty"`
}
