package delivery

import (
	"encoding/json"
	"github.com/encasol/tipsterchat/model"
	"io"
)

type JsonProxy interface {
	DecodeJson(Body io.ReadCloser) (model.Tip, error)
	EncodeJson(tips []model.Tip) ([]byte, error)
}

type JsonDecoder struct {
}

func (j JsonDecoder) DecodeJson(Body io.ReadCloser) (model.Tip, error) {
	decoder := json.NewDecoder(Body)
	var tip model.Tip
	err := decoder.Decode(&tip)
	if err != nil {
		panic(err)
	}

	return tip, nil
}

func (j JsonDecoder) EncodeJson(tips []model.Tip) ([]byte, error) {
	return json.Marshal(tips)
}
