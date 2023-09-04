package domain

import (
	"errors"
	"time"
)

type Participant struct {
	ID           int       `json:"id"`
	ContestID    int       `json:"contestID"`
	Name         string    `json:"name"`
	IsVerified   bool      `json:"isVerified"`
	Origin       string    `json:"origin"`
	FormURL      string    `json:"formURL"`
	VideoURL     string    `json:"videoURL"`
	PaymentProof string    `json:"paymentProof"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type StoreParticipant struct {
	ContestID    int    `json:"contestID" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Origin       string `json:"origin" binding:"required"`
	FormURL      string `json:"-"`
	VideoURL     string `json:"videoURL" binding:"required"`
	PaymentProof string `json:"-"`
}

func (p *StoreParticipant) Validate() error {
	switch {
	case p.ContestID < 1:
		return errors.New("ID Lomba tidak valid")
	case p.Name == "":
		return errors.New("Field nama tidak boleh kosong!")
	case p.Origin == "":
		return errors.New("Field asal tidak boleh kosong!")
	case p.VideoURL == "":
		return errors.New("Field video url tidak boleh kosong!")
	default:
		return nil
	}
}
