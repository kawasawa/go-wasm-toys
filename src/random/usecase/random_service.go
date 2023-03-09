package usecase

import (
	"toolbox/src/common/logging"
	"toolbox/src/random/domain/entity"
	"toolbox/src/random/domain/services"
)

type RandomService struct {
	DomainService services.IRandomDomainService
	Logger        logging.ILogger
}

func NewRandomService(s services.IRandomDomainService) *RandomService {
	return &RandomService{
		DomainService: s,
		Logger:        logging.GetLogger(),
	}
}

func (cs *RandomService) GenerateRandom(req *entity.GenerateRandomRequest) (*entity.GenerateRandomResponse, error) {
	cs.Logger.Log(logging.GenerateRandomStart, nil)
	defer cs.Logger.Log(logging.GenerateRandomEnd, nil)

	passwords, err := cs.DomainService.GenerateRandom(req.Characters, req.Length, req.Count)
	if err != nil {
		return &entity.GenerateRandomResponse{}, err
	}
	return &entity.GenerateRandomResponse{Randoms: passwords}, nil
}
