package config

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetsService struct {
	id  string
	srv *sheets.Service
}

func SetupSheets(sheetsID string, credentialPath string) (*SheetsService, error) {
	service, err := sheets.NewService(context.Background(), option.WithCredentialsFile(credentialPath))
	if err != nil {
		return nil, err
	}

	return &SheetsService{id: sheetsID, srv: service}, nil
}

func (srv *SheetsService) AppendRow(cellEnds string, values ...interface{}) error {
	inputValues := &sheets.ValueRange{
		Values: [][]interface{}{values},
	}

	resp, err := srv.srv.Spreadsheets.Values.Append(srv.id, "Sheet1!A:"+cellEnds, inputValues).ValueInputOption("RAW").Do()
	if err != nil {
		return err
	}

	println(resp.TableRange)

	return nil
}
