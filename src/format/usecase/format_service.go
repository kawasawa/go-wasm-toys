package usecase

import (
	"toolbox/src/common/logging"
	"toolbox/src/format/domain/entity"
	"toolbox/src/format/domain/services"
)

type FormatService struct {
	DomainService services.IFormatDomainService
	Logger        logging.ILogger
}

func NewFormatService(s services.IFormatDomainService) *FormatService {
	return &FormatService{
		DomainService: s,
		Logger:        logging.GetLogger(),
	}
}

func (cs *FormatService) FormatJson(req *entity.FormatRequest) (*entity.FormatResponse, error) {
	cs.Logger.Log(logging.FormatJsonStart, nil)
	defer cs.Logger.Log(logging.FormatJsonEnd, nil)

	encryptedText, err := cs.DomainService.FormatJson(req.PlainText)
	if err != nil {
		return &entity.FormatResponse{}, err
	}
	return &entity.FormatResponse{FormattedText: encryptedText}, nil
}
