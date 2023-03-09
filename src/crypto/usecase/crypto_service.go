package usecase

import (
	"toolbox/src/common/logging"
	"toolbox/src/crypto/domain/entity"
	"toolbox/src/crypto/domain/services"
)

type CryptoService struct {
	DomainService services.ICryptoDomainService
	Logger        logging.ILogger
}

func NewCryptoService(s services.ICryptoDomainService) *CryptoService {
	return &CryptoService{
		DomainService: s,
		Logger:        logging.GetLogger(),
	}
}

func (cs *CryptoService) EncryptAES(req *entity.EncryptRequest) (*entity.EncryptResponse, error) {
	cs.Logger.Log(logging.EncryptAESStart, nil)
	defer cs.Logger.Log(logging.EncryptAESEnd, nil)

	encryptedText, err := cs.DomainService.EncryptAES(req.PlainText, req.Key)
	if err != nil {
		return &entity.EncryptResponse{}, err
	}
	return &entity.EncryptResponse{EncryptedText: encryptedText}, nil
}

func (cs *CryptoService) DecryptAES(req *entity.DecryptRequest) (*entity.DecryptResponse, error) {
	cs.Logger.Log(logging.DecryptAESStart, nil)
	defer cs.Logger.Log(logging.DecryptAESEnd, nil)

	plainText, err := cs.DomainService.DecryptAES(req.EncryptedText, req.Key)
	if err != nil {
		return &entity.DecryptResponse{}, err
	}
	return &entity.DecryptResponse{PlainText: plainText}, nil
}
