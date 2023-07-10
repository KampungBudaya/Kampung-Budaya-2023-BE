package domain

import (
	"errors"
	"time"
)

type Participant struct {
	ID           int
	ContestID    int
	Name         string
	IsVerified   bool
	Origin       string
	PhoneNumber  string
	FormURL      string
	VideoURL     string
	PaymentProof string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type StoreParticipant struct {
	ContestID    int    `json:"contestID" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Origin       string `json:"origin" binding:"required"`
	PhoneNumber  string `json:"phoneNumber" binding:"required"`
	FormURL      string `json:"-"`
	VideoURL     string `json:"videoURL" binding:"required"`
	PaymentProof string `json:"-"`
}

type CleanParticipant struct {
	ID           int    `json:"id"`
	Origin       string `json:"origin"`
	Contest      string `json:"contest"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phoneNumber"`
	FormURL      string `json:"formURL"`
	PaymentProof string `json:"paymentProof"`
	VideoURL     string `json:"videoURL"`
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
