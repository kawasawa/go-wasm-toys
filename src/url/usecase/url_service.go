package usecase

import (
	"toolbox/src/common/logging"
	"toolbox/src/url/domain/entity"
	"toolbox/src/url/domain/services"
)

type UrlService struct {
	DomainService services.IUrlDomainService
	Logger        logging.ILogger
}

func NewUrlService(s services.IUrlDomainService) *UrlService {
	return &UrlService{
		DomainService: s,
		Logger:        logging.GetLogger(),
	}
}

func (cs *UrlService) EncodeUrl(req *entity.EncodeUrlRequest) (*entity.EncodeUrlResponse, error) {
	cs.Logger.Log(logging.EncodeUrlStart, nil)
	defer cs.Logger.Log(logging.EncodeUrlEnd, nil)

	encodedUrl, err := cs.DomainService.EncodeUrl(req.UrlStr)
	if err != nil {
		cs.Logger.Log(logging.EncodeBaseError, err)
	}

	return &entity.EncodeUrlResponse{EncodedUrl: encodedUrl}, nil
}

func (cs *UrlService) DecodeUrl(req *entity.DecodeUrlRequest) (*entity.DecodeUrlResponse, error) {
	cs.Logger.Log(logging.DecodeUrlStart, nil)
	defer cs.Logger.Log(logging.DecodeUrlEnd, nil)

	decodedUrl, err := cs.DomainService.DecodeUrl(req.UrlStr)
	if err != nil {
		return &entity.DecodeUrlResponse{}, err
	}

	return &entity.DecodeUrlResponse{DecodedUrl: decodedUrl}, nil
}
