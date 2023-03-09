package services

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"toolbox/src/hash/domain/entity"

	"github.com/pkg/errors"
)

type IHashDomainService interface {
	Hash(string, entity.HashMethod) (string, error)
}

type HashDomainService struct{}

func NewHashDomainService() *HashDomainService {
	return &HashDomainService{}
}

func (s *HashDomainService) Hash(plainText string, method entity.HashMethod) (string, error) {
	if plainText == "" {
		return "", errors.New("plainText is empty.")
	}

	hashedText := ""
	switch method {
	case entity.MD5:
		hashedText = fmt.Sprintf("%x", md5.Sum([]byte(plainText)))
	case entity.SHA1:
		hashedText = fmt.Sprintf("%x", sha1.Sum([]byte(plainText)))
	case entity.SHA256:
		hashedText = fmt.Sprintf("%x", sha256.Sum256([]byte(plainText)))
	case entity.SHA384:
		hashedText = fmt.Sprintf("%x", sha512.Sum384([]byte(plainText)))
	case entity.SHA512:
		hashedText = fmt.Sprintf("%x", sha512.Sum512([]byte(plainText)))
	default:
		return "", errors.New("method out of range.")
	}
	return hashedText, nil
}
