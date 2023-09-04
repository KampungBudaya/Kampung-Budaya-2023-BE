package domain

import (
	"errors"
	"time"
)

type Contest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type StoreContest struct {
	Name string `json:"name" binding:"required"`
}

func (c *StoreContest) Validate() error {
	if c.Name == "" {
		return errors.New("Field nama tidak boleh kosong!")
	}

	return nil
}
