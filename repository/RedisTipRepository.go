package repository

import (
	"errors"
	"fmt"
	"github.com/encasol/tipsterchat/model"
	"github.com/go-redis/redis"
)

type RedisTipRepository struct {
	Connection *redis.Client
}

func (redis RedisTipRepository) All() ([]model.Tip, error) {
	fmt.Println("Funciona la interficie")
	return nil, errors.New("Failed All()")
}
