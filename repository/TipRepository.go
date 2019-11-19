package repository

import "github.com/encasol/tipsterchat/model"

type TipRepository interface {
	All() ([]model.Tip, error)
	Add(tip model.Tip) error
}
