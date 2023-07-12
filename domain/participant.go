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
	VideoURL     string
	PaymentProof string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type StoreParticipant struct {
	ContestID    int
	Name         string
	Origin       string
	VideoURL     string
	PaymentProof string
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
	case p.PaymentProof == "":
		return errors.New("Field bukti pembayaran tidak boleh kosong!")
	default:
		return nil
	}
}
