package domain

import "errors"

type FaqReq struct {
	Category string `json:"category"`
	Title    string `json:"title"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Faq struct {
	ID       int    `json:"id" db:"id"`
	Category string `json:"category" db:"category"`
	Title    string `json:"title" db:"title"`
	Question string `json:"question" db:"question"`
	Answer   string `json:"answer" db:"answer"`
}

func (faq *FaqReq) Validate() error {
	switch {
	case faq.Category == "":
		return errors.New("FIELD CATEGORY TIDAK BOLEH KOSONG")
	case faq.Title == "":
		return errors.New("FIELD TITLE TIDAK BOLEH KOSONG")
	case faq.Question == "":
		return errors.New("FIELD QUESTION TIDAK BOLEH KOSONG")
	case faq.Answer == "":
		return errors.New("FIELD ANSWER TIDAK BOLEH KOSONG")
	}
	return nil
}
