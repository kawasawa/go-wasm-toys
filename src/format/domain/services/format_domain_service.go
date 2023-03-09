package services

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

type IFormatDomainService interface {
	FormatJson(string) (string, error)
}

type FormatDomainService struct{}

func NewFormatDomainService() *FormatDomainService {
	return &FormatDomainService{}
}

func (s *FormatDomainService) FormatJson(plainText string) (string, error) {
	if plainText == "" {
		return "", errors.New("plainText is empty.")
	}

	var formatted bytes.Buffer
	err := json.Indent(&formatted, []byte(strings.TrimSpace(plainText)), "", "\t")
	if err != nil {
		return "", err
	}

	return formatted.String(), nil
}
