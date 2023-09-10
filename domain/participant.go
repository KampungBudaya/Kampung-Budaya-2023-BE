package domain

import (
	"errors"
	"regexp"
)

type ParticipantDB struct {
	ID           int     `db:"id"`
	Name         string  `db:"name"`
	Birth        string  `db:"birth"`
	Category     string  `db:"category"`
	Status       string  `db:"status"`
	Contest      string  `db:"contest_name"`
	Institution  string  `db:"institution"`
	Email        string  `db:"email"`
	Instagram    string  `db:"instagram"`
	Line         string  `db:"line"`
	PhoneNumber  string  `db:"phone_number"`
	Form         string  `db:"form"`
	VideoURL     *string `db:"video_url"`
	PaymentProof string  `db:"payment_proof"`
}

type StoreParticipant struct {
	ContestID    int    `json:"contestID" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Birth        string `json:"birth" binding:"required"`
	Category     string `json:"category" binding:"required"`
	Institution  string `json:"institution" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Instagram    string `json:"instagram" binding:"required"`
	Line         string `json:"line" binding:"required"`
	PhoneNumber  string `json:"phoneNumber" binding:"required"`
	FormURL      string `json:"-"`
	VideoURL     string `json:"videoURL" binding:"required"`
	PaymentProof string `json:"-"`
}

type CleanParticipant struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Birth        string `json:"birth"`
	Category     string `json:"category"`
	Status       string `json:"status"`
	Contest      string `json:"contest"`
	Institution  string `json:"institution"`
	Email        string `json:"email"`
	Instagram    string `json:"instagram"`
	Line         string `json:"line"`
	PhoneNumber  string `json:"phoneNumber"`
	Form         string `json:"form"`
	VideoURL     string `json:"videoURL"`
	PaymentProof string `json:"paymentProof"`
}

func (p *StoreParticipant) Validate() error {
	switch {
	case p.ContestID < 1 || p.ContestID > 4:
		return errors.New("ID LOMBA TIDAK VALID")
	case p.Name == "":
		return errors.New("FIELD NAMA TIDAK BOLEH KOSONG")
	case p.Birth == "":
		return errors.New("FIELD TANGGAL LAHIR TIDAK BOLEH KOSONG")
	case p.Institution == "":
		return errors.New("FIELD INSTITUSI TIDAK BOLEH KOSONG")
	case p.Email == "" || validateEmail(p.Email):
		return errors.New("FIELD EMAIL TIDAK VALID")
	case p.Instagram == "":
		return errors.New("FIELD INSTAGRAM TIDAK BOLEH KOSONG")
	case p.Line == "":
		return errors.New("FIELD LINE TIDAK BOLEH KOSONG")
	case p.PhoneNumber == "" || validatePhoneNumber(p.PhoneNumber):
		return errors.New("FIELD NOMOR TELEPON TIDAK VALID")
	case p.Category == "":
		return errors.New("FIELD CATEGORY TIDAK BOLEH KOSONG")
	case p.ContestID == 1 || p.ContestID == 2 || p.ContestID == 3:
		if (p.VideoURL == "") || (p.Category == "UMUM" && p.VideoURL == "") {
			return errors.New("FIELD LINK VIDEO TIDAK BOLEH KOSONG")
		}
	case p.ContestID == 4:
		if p.Category == "UMUM" {
			return errors.New("MAAF LOMBA INI HANYA BISA DIIKUTI OLEH FORDA")
		}
	default:
		return nil
	}
	return nil
}

func (p *ParticipantDB) Clean() *CleanParticipant {
	cleanParticipant := &CleanParticipant{
		ID:           p.ID,
		Name:         p.Name,
		Birth:        p.Birth,
		Category:     p.Category,
		Status:       p.Status,
		Contest:      p.Contest,
		Institution:  p.Institution,
		Email:        p.Email,
		Instagram:    p.Instagram,
		Line:         p.Line,
		PhoneNumber:  p.PhoneNumber,
		Form:         p.Form,
		PaymentProof: p.PaymentProof,
	}
	if p.VideoURL != nil {
		cleanParticipant.VideoURL = *p.VideoURL
	}
	return cleanParticipant
}

func validatePhoneNumber(phoneNumber string) bool {
	if len(phoneNumber) > 14 || len(phoneNumber) < 12 || phoneNumber[:1] != "0" {
		return true
	}

	for _, char := range phoneNumber {
		if char < '0' || char > '9' || char == '+' {
			return true
		}
	}
	return false
}

func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return !emailRegex.MatchString(email)
}
