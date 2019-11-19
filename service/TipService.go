package service

import (
	"github.com/encasol/tipsterchat/model"
	"github.com/encasol/tipsterchat/repository"
)

type ITipService interface {
	ListTips() ([]model.Tip, error)
	AddTip(tip model.Tip) error
}

type TipService struct {
	TipRepo repository.TipRepository
}

func (s TipService) ListTips() ([]model.Tip, error) {
	tips, err := s.TipRepo.All()
	if err != nil {
		panic(err)
	}

	return tips, nil
}

func (s TipService) AddTip(tip model.Tip) error {
	return s.TipRepo.Add(tip)
}
