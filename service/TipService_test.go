package service_test

import (
	"github.com/encasol/tipsterchat/model"
	"github.com/encasol/tipsterchat/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

type MockedTipRepo struct {
	mock.Mock
}

func (mock MockedTipRepo) All() ([]model.Tip, error) {
	args := mock.Called()
	return args.Get(0).([]model.Tip), args.Error(1)
}

func (mock MockedTipRepo) Add(tip model.Tip) error {
	args := mock.Called(tip)
	return args.Error(0)
}

func TestListTipsEmptySet(t *testing.T) {
	repo := new(MockedTipRepo)
	svc := service.TipService{TipRepo: repo}
	emtyTips := []model.Tip{}

	repo.On("All").Return(emtyTips, nil)
	tip, err := svc.ListTips()
	assert.Equal(t, tip, emtyTips)
	assert.Equal(t, err, nil)
}

func TestListTipsPanic(t *testing.T) {
	repo := new(MockedTipRepo)
	svc := service.TipService{TipRepo: repo}
	emtyTips := []model.Tip{}

	repo.On("All").Return(emtyTips, 2)
	assert.Panics(t, func() { svc.ListTips() }, nil)
}

func TestAddTip(t *testing.T) {
	repo := new(MockedTipRepo)
	svc := service.TipService{TipRepo: repo}
	emtyTip := model.Tip{}

	repo.On("Add", emtyTip).Return(nil)
	err := svc.AddTip(emtyTip)
	assert.Equal(t, err, nil)
}
