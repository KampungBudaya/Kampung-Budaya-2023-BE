package domain

import (
	"time"
)

type User struct {
	ID         int       `db:"id"`
	Roles      *string   `db:"roles"`
	Provider   string    `db:"provider"`
	ProviderID *string   `db:"provider_id"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type StoreUser struct {
	Provider string `json:"provider" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type CleanUser struct {
	ID           int       `json:"id"`
	Roles        string    `json:"roles"`
	Provider     string    `json:"provider"`
	ProviderID   string    `json:"providerID"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	RegisteredAt time.Time `json:"registeredAt"`
	LastModified time.Time `json:"lastModified"`
}

func (u *User) Clean() *CleanUser {
	user := &CleanUser{
		ID:           u.ID,
		Provider:     u.Provider,
		Name:         u.Name,
		Email:        u.Email,
		RegisteredAt: u.CreatedAt,
		LastModified: u.UpdatedAt,
	}

	if u.Roles != nil {
		user.Roles = *u.Roles
	}

	if u.ProviderID != nil {
		user.ProviderID = *u.ProviderID
	}

	return user
}
