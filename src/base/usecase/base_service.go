package usecase

import (
	"toolbox/src/base/domain/entity"
	"toolbox/src/base/domain/services"
	"toolbox/src/common/logging"
)

type BaseService struct {
	DomainService services.IBaseDomainService
	Logger        logging.ILogger
}

func NewBaseService(s services.IBaseDomainService) *BaseService {
	return &BaseService{
		DomainService: s,
		Logger:        logging.GetLogger(),
	}
}

func (cs *BaseService) EncodeBase(req *entity.EncodeBaseRequest) (*entity.EncodeBaseResponse, error) {
	cs.Logger.Log(logging.EncodeBaseStart, nil)
	defer cs.Logger.Log(logging.EncodeBaseEnd, nil)

	base32, err := cs.DomainService.EncodeBase(req.PlainText, entity.Base32)
	if err != nil {
		cs.Logger.Log(logging.EncodeBaseError, err)
	}
	base64, err := cs.DomainService.EncodeBase(req.PlainText, entity.Base64)
	if err != nil {
		cs.Logger.Log(logging.EncodeBaseError, err)
	}

	return &entity.EncodeBaseResponse{
		Base32: base32,
		Base64: base64,
	}, nil
}

func (cs *BaseService) DecodeBase(req *entity.DecodeBaseRequest) (*entity.DecodeBaseResponse, error) {
	cs.Logger.Log(logging.DecodeBaseStart, nil)
	defer cs.Logger.Log(logging.DecodeBaseEnd, nil)

	plain32, err := cs.DomainService.DecodeBase(req.BaseText, entity.Base32)
	if err != nil {
		cs.Logger.Log(logging.DecodeBaseError, err)
	}
	plain64, err := cs.DomainService.DecodeBase(req.BaseText, entity.Base64)
	if err != nil {
		cs.Logger.Log(logging.DecodeBaseError, err)
	}

	return &entity.DecodeBaseResponse{
		Plain32: plain32,
		Plain64: plain64,
	}, nil
}
