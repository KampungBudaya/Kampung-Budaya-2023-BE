package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/contest/repository"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/config"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/email"
)

type ContestUsecaseImpl interface {
	RegisterContest(ctx context.Context, req domain.StoreParticipant, photos []multipart.File) (int, error)
	GetAllParticipants(ctx context.Context) ([]*domain.CleanParticipant, error)
	GetParticipantByID(ctx context.Context, id int) (*domain.CleanParticipant, error)
	AcceptParticipant(ctx context.Context, id int) error
	RejectParticipant(ctx context.Context, id int) error
}

type ContestUsecase struct {
	contest repository.ContestRepositoryImpl
	sheets  *config.SheetsService
}

func NewContestUsecase(contest repository.ContestRepositoryImpl, sheets *config.SheetsService) ContestUsecaseImpl {
	return &ContestUsecase{
		contest: contest,
		sheets:  sheets,
	}
}

func (uc *ContestUsecase) RegisterContest(ctx context.Context, req domain.StoreParticipant, photos []multipart.File) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	formByte, err := io.ReadAll(photos[0])
	if err != nil {
		return 0, err
	}
	linkForm, err := uploadPhotos(ctx, formByte, fmt.Sprintf("form-%s-%s-%d", req.Name, req.Institution, req.ContestID))
	if err != nil {
		return 0, err
	}

	paymentProofByte, err := io.ReadAll(photos[1])
	if err != nil {
		return 0, err
	}
	linkPaymentProof, err := uploadPhotos(ctx, paymentProofByte, fmt.Sprintf("payment-proof-%s-%s-%d", req.Name, req.Institution, req.ContestID))
	if err != nil {
		return 0, err
	}

	id, err := uc.contest.Create(ctx, &req, []string{linkForm, linkPaymentProof})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *ContestUsecase) GetAllParticipants(ctx context.Context) ([]*domain.CleanParticipant, error) {
	participants, err := uc.contest.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return participants, nil
}

func (uc *ContestUsecase) GetParticipantByID(ctx context.Context, id int) (*domain.CleanParticipant, error) {
	participant, err := uc.contest.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return participant, nil
}

func (uc *ContestUsecase) AcceptParticipant(ctx context.Context, id int) error {
	tx, err := uc.contest.BeginTx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	participant, err := uc.contest.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if participant.Status == "ACCEPTED" {
		return errors.New("PESERTA SUDAH DITERIMA")
	}

	err = uc.contest.UpdateStatus(ctx, id, "ACCEPTED")
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	if err := uc.sheets.AppendRow(
		"O",
		participant.ID,
		participant.Name,
		participant.Institution,
		participant.Category,
		participant.Birth,
		participant.Email,
		participant.PhoneNumber,
		participant.Instagram,
		participant.Line,
		participant.Form,
		participant.VideoURL,
		participant.PaymentProof,
		participant.Contest,
		"ACCEPTED",
	); err != nil {
		return err
	}

	mail := email.NewMailClient()
	mail.SetSender(os.Getenv("CONFIG_SENDER_NAME"))
	mail.SetReciever(participant.Email)
	mail.SetSubject("Announcement Email")

	var groupLink string
	if participant.Contest == "Tari" {
		groupLink = os.Getenv("GROUP_LINK_TARI")
	} else if participant.Contest == "Musik" {
		groupLink = os.Getenv("GROUP_LINK_MUSIK")
	} else if participant.Contest == "Busana Kreasi" {
		groupLink = os.Getenv("GROUP_LINK_BUSANA_KREASI")
	} else {
		groupLink = os.Getenv("GROUP_LINK_STAND_BAZAR")
	}

	mail.SetBodyHTML(participant.Name, participant.Contest, groupLink)
	if err := mail.SendMail(); err != nil {
		return err
	}
	return nil
}

func (uc *ContestUsecase) RejectParticipant(ctx context.Context, id int) error {
	tx, err := uc.contest.BeginTx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	participant, err := uc.contest.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if participant.Status == "REJECTED" {
		return errors.New("PESERTA SUDAH DITOLAK")
	}

	err = uc.contest.UpdateStatus(ctx, id, "REJECTED")
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func uploadPhotos(ctx context.Context, file []byte, fileName string) (string, error) {
	fb, err := config.InitFirebase()
	if err != nil {
		return "", err
	}
	link, err := fb.UploadFile(ctx, file, fileName)
	if err != nil {
		return "", err
	}
	return link, nil
}
