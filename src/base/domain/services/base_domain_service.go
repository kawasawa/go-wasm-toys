package services

import (
	"encoding/base32"
	"encoding/base64"

	"toolbox/src/base/domain/entity"

	"github.com/pkg/errors"
)

type IBaseDomainService interface {
	EncodeBase(string, entity.BaseMethod) (string, error)
	DecodeBase(string, entity.BaseMethod) (string, error)
}

type BaseDomainService struct{}

func NewBaseDomainService() *BaseDomainService {
	return &BaseDomainService{}
}

func (s *BaseDomainService) EncodeBase(plainText string, method entity.BaseMethod) (string, error) {
	if plainText == "" {
		return "", errors.New("plainText is empty.")
	}

	encodedText := ""
	switch method {
	case entity.Base32:
		encodedText = base32.StdEncoding.EncodeToString([]byte(plainText))
	case entity.Base64:
		encodedText = base64.StdEncoding.EncodeToString([]byte(plainText))
	default:
		return "", errors.New("method out of range.")
	}
	return encodedText, nil
}

func (s *BaseDomainService) DecodeBase(baseText string, method entity.BaseMethod) (string, error) {
	if baseText == "" {
		return "", errors.New("baseText is empty.")
	}

	var decoded []byte
	var err error
	switch method {
	case entity.Base32:
		decoded, err = base32.StdEncoding.DecodeString(baseText)
	case entity.Base64:
		decoded, err = base64.StdEncoding.DecodeString(baseText)
	default:
		return "", errors.New("method out of range.")
	}

	return string(decoded), err
}
