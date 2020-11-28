package entities

import "time"

//UserProfile object of user_profiles
type UserProfile struct {
	ID          int64      `json:"id,omitempty"`
	ProfileName string     `json:"profileName,omitempty"`
	PhoneNumber string     `json:"phoneNumber,omitempty"`
	Email       string     `json:"email,omitempty"`
	Gender      string     `json:"gender,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}
