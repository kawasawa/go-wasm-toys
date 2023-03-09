package services

import (
	"net/url"
	"regexp"

	"github.com/pkg/errors"
)

type IUrlDomainService interface {
	EncodeUrl(string) (string, error)
	DecodeUrl(string) (string, error)
}

type UrlDomainService struct{}

func NewUrlDomainService() *UrlDomainService {
	return &UrlDomainService{}
}

func (s *UrlDomainService) EncodeUrl(urlStr string) (string, error) {
	if urlStr == "" {
		return "", errors.New("plainText is empty.")
	}

	encodedUrl := url.QueryEscape(urlStr)
	encodedUrl = regexp.MustCompile(`([^%])(\+)`).ReplaceAllString(encodedUrl, "$1%20")
	return encodedUrl, nil
}

func (s *UrlDomainService) DecodeUrl(urlStr string) (string, error) {
	if urlStr == "" {
		return "", errors.New("baseText is empty.")
	}

	decodedUrl, err := url.QueryUnescape(urlStr)
	if err != nil {
		return "", err
	}
	return string(decodedUrl), err
}
