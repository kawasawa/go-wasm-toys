package usecase

import (
	"toolbox/src/common/logging"
	"toolbox/src/hash/domain/entity"
	"toolbox/src/hash/domain/services"
)

type HashService struct {
	DomainService services.IHashDomainService
	Logger        logging.ILogger
}

func NewHashService(s services.IHashDomainService) *HashService {
	return &HashService{
		DomainService: s,
		Logger:        logging.GetLogger(),
	}
}

func (cs *HashService) Hash(req *entity.HashRequest) (*entity.HashResponse, error) {
	cs.Logger.Log(logging.HashStart, nil)
	defer cs.Logger.Log(logging.HashEnd, nil)

	md5, err := cs.DomainService.Hash(req.PlainText, entity.MD5)
	if err != nil {
		return &entity.HashResponse{}, err
	}
	sha1, err := cs.DomainService.Hash(req.PlainText, entity.SHA1)
	if err != nil {
		return &entity.HashResponse{}, err
	}
	sha256, err := cs.DomainService.Hash(req.PlainText, entity.SHA256)
	if err != nil {
		return &entity.HashResponse{}, err
	}
	sha384, err := cs.DomainService.Hash(req.PlainText, entity.SHA384)
	if err != nil {
		return &entity.HashResponse{}, err
	}
	sga512, err := cs.DomainService.Hash(req.PlainText, entity.SHA512)
	if err != nil {
		return &entity.HashResponse{}, err
	}

	return &entity.HashResponse{
		MD5:    md5,
		SHA1:   sha1,
		SHA256: sha256,
		SHA384: sha384,
		SHA512: sga512,
	}, nil
}
