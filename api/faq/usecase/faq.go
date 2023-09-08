package usecase

type FaqUsecaseImpl interface {
}

type FaqUsecase struct {
	uc *FaqUsecaseImpl
}

func NewFaqUsecase(uc *FaqUsecaseImpl) *FaqUsecase {
	return &FaqUsecase{uc}
}
