package domain

import (
	"errors"
	"time"
)

type Contest struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StoreContest struct {
	Name string `json:"name" binding:"required"`
}

type CleanContest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *StoreContest) Validate() error {
	if c.Name == "" {
		return errors.New("FIELD NAMA TIDAK BOLEH KOSONG")
	}

	return nil
}
