package domain

import (
	"errors"
)

type ParticipantDB struct {
	ID           int     `db:"id"`
	Name         string  `db:"name"`
	Category     string  `db:"category"`
	Status       string  `db:"status"`
	Contest      string  `db:"contest_name"`
	Origin       string  `db:"origin"`
	PhoneNumber  string  `db:"phone_number"`
	FormURL      *string `db:"form_url"`
	VideoURL     *string `db:"video_url"`
	PaymentProof *string `db:"payment_proof"`
}

type StoreParticipant struct {
	ContestID    int    `json:"contestID" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Category     string `json:"category" binding:"required"`
	Origin       string `json:"origin" binding:"required"`
	PhoneNumber  string `json:"phoneNumber" binding:"required"`
	FormURL      string `json:"-"`
	VideoURL     string `json:"videoURL" binding:"required"`
	PaymentProof string `json:"-"`
}

type CleanParticipant struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	Status       string `json:"status"`
	Contest      string `json:"contest"`
	Origin       string `json:"origin"`
	PhoneNumber  string `json:"phoneNumber"`
	FormURL      string `json:"formURL"`
	VideoURL     string `json:"videoURL"`
	PaymentProof string `json:"paymentProof"`
}

func (p *StoreParticipant) Validate() error {
	switch {
	case p.ContestID < 1 || p.ContestID > 5:
		return errors.New("ID LOMBA TIDAK VALID")
	case p.Name == "":
		return errors.New("FIELD NAMA TIDAK BOLEH KOSONG")
	case p.Origin == "":
		return errors.New("FIELD ASAL TIDAK BOLEH KOSONG")
	case p.PhoneNumber == "":
		return errors.New("FIELD NOMOR TELEPON TIDAK BOLEH KOSONG")
	case p.Category == "":
		return errors.New("FIELD CATEGORY TIDAK BOLEH KOSONG")
	default:
		return nil
	}
}

func (p *ParticipantDB) Clean() *CleanParticipant {
	cleanParticipant := &CleanParticipant{
		ID:          p.ID,
		Name:        p.Name,
		Category:    p.Category,
		Status:      p.Status,
		Contest:     p.Contest,
		Origin:      p.Origin,
		PhoneNumber: p.PhoneNumber,
	}
	if p.VideoURL != nil {
		cleanParticipant.VideoURL = *p.VideoURL
	}
	if p.FormURL != nil {
		cleanParticipant.FormURL = *p.FormURL
	}
	if p.PaymentProof != nil {
		cleanParticipant.PaymentProof = *p.PaymentProof
	}
	return cleanParticipant
}
