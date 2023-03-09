package usecase

import (
	"toolbox/src/case/domain/entity"
	"toolbox/src/case/domain/services"
	"toolbox/src/common/logging"
)

type CaseService struct {
	DomainService services.ICaseDomainService
	Logger        logging.ILogger
}

func NewCaseService(s services.ICaseDomainService) *CaseService {
	return &CaseService{
		DomainService: s,
		Logger:        logging.GetLogger(),
	}
}

func (cs *CaseService) ConvertCase(req *entity.CaseRequest) (*entity.CaseResponse, error) {
	cs.Logger.Log(logging.ConvertCaseStart, nil)
	defer cs.Logger.Log(logging.ConvertCaseEnd, nil)

	camel, err := cs.DomainService.ConvertCase(req.PlainText, entity.Camel)
	if err != nil {
		return &entity.CaseResponse{}, err
	}
	pascal, err := cs.DomainService.ConvertCase(req.PlainText, entity.Pascal)
	if err != nil {
		return &entity.CaseResponse{}, err
	}
	snake, err := cs.DomainService.ConvertCase(req.PlainText, entity.Snake)
	if err != nil {
		return &entity.CaseResponse{}, err
	}
	kebab, err := cs.DomainService.ConvertCase(req.PlainText, entity.Kebab)
	if err != nil {
		return &entity.CaseResponse{}, err
	}
	upper, err := cs.DomainService.ConvertCase(req.PlainText, entity.Upper)
	if err != nil {
		return &entity.CaseResponse{}, err
	}
	lower, err := cs.DomainService.ConvertCase(req.PlainText, entity.Lower)
	if err != nil {
		return &entity.CaseResponse{}, err
	}

	return &entity.CaseResponse{
		Camel:  camel,
		Pascal: pascal,
		Snake:  snake,
		Kebab:  kebab,
		Upper:  upper,
		Lower:  lower,
	}, nil
}
