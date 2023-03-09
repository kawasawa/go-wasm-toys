package usecase

import (
	"toolbox/src/common/logging"
	"toolbox/src/jwt/domain/entity"
	"toolbox/src/jwt/domain/services"
)

type JwtService struct {
	DomainService services.IJwtDomainService
	Logger        logging.ILogger
}

func NewJwtService(s services.IJwtDomainService) *JwtService {
	return &JwtService{
		DomainService: s,
		Logger:        logging.GetLogger(),
	}
}

func (cs *JwtService) DecodeJwt(req *entity.DecodeJwtRequest) (*entity.DecodeJwtResponse, error) {
	cs.Logger.Log(logging.DecodeJwtStart, nil)
	defer cs.Logger.Log(logging.DecodeJwtEnd, nil)

	header, payload, verified, err := cs.DomainService.DecodeJwt(req.TokenString, req.Key)
	if err != nil {
		return &entity.DecodeJwtResponse{}, err
	}

	return &entity.DecodeJwtResponse{
		Header:   header,
		Payload:  payload,
		Verified: verified,
	}, nil
}
