package repository

import (
	"context"
	"fmt"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/jmoiron/sqlx"
)

type ContestRepositoryImpl interface {
	Create(ctx context.Context, req *domain.StoreParticipant, linkPhotos []string) (int, error)
	GetAll(ctx context.Context) ([]*domain.CleanParticipant, error)
	GetByID(ctx context.Context, id int) (*domain.CleanParticipant, error)
	UpdateStatus(ctx context.Context, id int) error
	BeginTx() (*sqlx.Tx, error)
}

type ContestRepository struct {
	mysql *sqlx.DB
}

func NewContestRepository(mysql *sqlx.DB) ContestRepositoryImpl {
	return &ContestRepository{
		mysql: mysql,
	}
}

func (r *ContestRepository) Create(ctx context.Context, req *domain.StoreParticipant, linkPhotos []string) (int, error) {
	argKV := map[string]interface{}{
		"contest_id":    req.ContestID,
		"name":          req.Name,
		"origin":        req.Origin,
		"phone_number":  req.PhoneNumber,
		"video_url":     req.VideoURL,
		"form_url":      linkPhotos[0],
		"payment_proof": linkPhotos[1],
	}

	query, args, err := sqlx.Named(queryRegisterCompetition, argKV)
	if err != nil {
		return 0, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return 0, err
	}
	query = r.mysql.Rebind(query)
	result, err := r.mysql.ExecContext(ctx, query, args...)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *ContestRepository) GetByID(ctx context.Context, id int) (*domain.CleanParticipant, error) {
	query := fmt.Sprintf(queryGetParticipants, "WHERE participants.id = ?")
	var participant domain.ParticipantDB

	if err := r.mysql.QueryRowxContext(ctx, query, id).StructScan(&participant); err != nil {
		return nil, err
	}
	return participant.Clean(), nil
}

func (r *ContestRepository) GetAll(ctx context.Context) ([]*domain.CleanParticipant, error) {
	var participants []*domain.CleanParticipant

	query := fmt.Sprintf(queryGetParticipants, "")
	query, args, err := sqlx.Named(query, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = r.mysql.Rebind(query)

	rows, err := r.mysql.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var participant domain.ParticipantDB
		err := rows.StructScan(&participant)
		if err != nil {
			return nil, err
		}
		participants = append(participants, participant.Clean())
	}
	return participants, nil
}

func (r *ContestRepository) UpdateStatus(ctx context.Context, id int) error {
	argKV := map[string]interface{}{
		"is_verified": true,
		"id":          id,
	}

	query, args, err := sqlx.Named(queryUpdateParticipant, argKV)
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

func (r *ContestRepository) BeginTx() (*sqlx.Tx, error) {
	return r.mysql.Beginx()
}
