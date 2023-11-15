package repository

import (
	"context"
	"fmt"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/jmoiron/sqlx"
)

type FaqRepositoryImpl interface {
	AddFaq(ctx context.Context, req *domain.FaqReq) (int, error)
}

type FaqRepository struct {
	mysql *sqlx.DB
}

func NewFaqRepository(mysql *sqlx.DB) *FaqRepository {
	return &FaqRepository{
		mysql: mysql,
	}
}

func (r *FaqRepository) Add(ctx context.Context, req *domain.FaqReq) (int, error) {
	argKV := map[string]interface{}{
		"category": req.Category,
		"title":    req.Title,
		"question": req.Question,
		"answer":   req.Answer,
	}

	query, args, err := sqlx.Named(queryAddFaq, argKV)
	if err != nil {
		return 0, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return 0, err
	}
	query = r.mysql.Rebind(query)
	res, err := r.mysql.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *FaqRepository) GetAll(ctx context.Context) ([]*domain.Faq, error) {
	var faqs []*domain.Faq

	query := fmt.Sprintf(queryGetFaq, "")
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
		faq := &domain.Faq{}
		err := rows.StructScan(faq)
		if err != nil {
			return nil, err
		}
		faqs = append(faqs, faq)
	}
	return faqs, nil
}

func (r *FaqRepository) GetByID(ctx context.Context, id int) (*domain.Faq, error) {
	query := fmt.Sprintf(queryGetFaq, "WHERE participants.id = ?")
	var faq domain.Faq

	if err := r.mysql.QueryRowxContext(ctx, query, id).StructScan(&faq); err != nil {
		return nil, err
	}
	return &faq, nil
}

func (r *FaqRepository) Update(ctx context.Context, req domain.FaqReq, id int) (int, error) {
	argKV := map[string]interface{}{
		"category": req.Category,
		"title":    req.Title,
		"question": req.Question,
		"answer":   req.Answer,
		"id":       id,
	}

	query, args, err := sqlx.Named(queryUpdateFaq, argKV)
	if err != nil {
		return 0, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return 0, err
	}
	query = r.mysql.Rebind(query)
	res, err := r.mysql.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

func (r *FaqRepository) Delete(ctx context.Context, id int) (int, error) {
	argKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(queryDeleteFaq, argKV)
	if err != nil {
		return 0, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return 0, err
	}
	query = r.mysql.Rebind(query)
	res, err := r.mysql.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(affected), nil
}
