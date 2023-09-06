package repository

import "context"

type OAuthRepositoryImpl interface {
	GetByID(ctx context.Context, id string)
}
