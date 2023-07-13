package domain

import (
	"errors"
	"time"
)

type Contest struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StoreContest struct {
	Name string
}

func (c *StoreContest) Validate() (*Contest, error) {
	if c.Name == "" {
		return nil, errors.New("Field nama tidak boleh kosong!")
	}

	contest := &Contest{Name: c.Name}

	return contest, nil
}
