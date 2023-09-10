package repository

import (
	"context"
	"fmt"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/jmoiron/sqlx"
)

type OAuthRepositoryImpl interface {
	GetByEmail(email string, ctx context.Context) (*domain.CleanUser, error)
	UpdateProviderID(id int, providerID string, ctx context.Context) error
	BeginTx() (*sqlx.Tx, error)
}

type OAuthRepository struct {
	mysql *sqlx.DB
}

func NewOAuthRepository(mysql *sqlx.DB) OAuthRepositoryImpl {
	return &OAuthRepository{mysql: mysql}
}

func (r *OAuthRepository) GetByEmail(email string, ctx context.Context) (*domain.CleanUser, error) {
	query := fmt.Sprintf(queryGetUserByEmail, "WHERE users.email = ? LIMIT 1")

	var user *domain.User
	if err := r.mysql.QueryRowxContext(ctx, query, email).StructScan(user); err != nil {
		return nil, err
	}

	return user.Clean(), nil
}

func (r *OAuthRepository) UpdateProviderID(id int, providerID string, ctx context.Context) error {
	queryArgs := map[string]interface{}{
		"provider_id": providerID,
		"id":          id,
	}

	query, args, err := sqlx.Named(queryUpdateUserProviderID, queryArgs)
	if err != nil {
		return err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return err
	}
	query = r.mysql.Rebind(query)

	_, err = r.mysql.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *OAuthRepository) BeginTx() (*sqlx.Tx, error) {
	return r.mysql.Beginx()
}
