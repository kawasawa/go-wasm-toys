package services

import (
	"errors"
	"math/rand"
	"time"
)

type IRandomDomainService interface {
	GenerateRandom(string, int, int) ([]string, error)
}

type RandomDomainService struct{}

func NewRandomDomainService() *RandomDomainService {
	return &RandomDomainService{}
}

func (s *RandomDomainService) GenerateRandom(characters string, length int, count int) ([]string, error) {
	if characters == "" {
		return nil, errors.New("characters is empty.")
	}
	if length <= 0 {
		return nil, errors.New("length less than or equal to 0.")
	}
	if count <= 0 {
		return nil, errors.New("count less than or equal to 0.")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]string, 0, count)
	for j := 0; j < count; j++ {
		value := make([]byte, 0, length)
		for j := 0; j < length; j++ {
			value = append(value, byte(characters[r.Intn(len(characters))]))
		}
		result = append(result, string(value))
	}
	return result, nil
}
