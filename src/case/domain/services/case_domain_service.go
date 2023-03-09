package services

import (
	"regexp"
	"strings"

	"toolbox/src/case/domain/entity"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

type ICaseDomainService interface {
	ConvertCase(string, entity.CaseMethod) (string, error)
}

type CaseDomainService struct{}

func NewCaseDomainService() *CaseDomainService {
	return &CaseDomainService{}
}

func (s *CaseDomainService) ConvertCase(plainText string, method entity.CaseMethod) (string, error) {
	if plainText == "" {
		return "", errors.New("plainText is empty.")
	}

	convertedText := ""
	switch method {
	case entity.Camel:
		convertedText = strcase.ToLowerCamel(plainText)
	case entity.Pascal:
		convertedText = strcase.ToCamel(plainText)
	case entity.Snake:
		convertedText = strcase.ToSnake(plainText)
	case entity.Kebab:
		convertedText = strcase.ToKebab(plainText)
	case entity.Upper:
		convertedText = strings.ToUpper(plainText)
	case entity.Lower:
		convertedText = strings.ToLower(plainText)
	default:
		return "", errors.New("the method" + string(method) + "out of range.")
	}
	return convertedText, nil
}

var (
	match1 = regexp.MustCompile("(.)([A-Z][a-z]+)")
	match2 = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func ConvertToCamelCase(text string) string {
	tmp := text
	tmp = match1.ReplaceAllString(tmp, "${1}_${2}")
	tmp = match2.ReplaceAllString(tmp, "${1}_${2}")
	return strings.ToLower(tmp)
}

func ConvertToSnakeCase(text string) string {
	tmp := text
	tmp = match1.ReplaceAllString(tmp, "${1}_${2}")
	tmp = match2.ReplaceAllString(tmp, "${1}_${2}")
	return strings.ToLower(tmp)
}

func ConvertToKebabCase(text string) string {
	tmp := text
	tmp = match1.ReplaceAllString(tmp, "${1}-${2}")
	tmp = match2.ReplaceAllString(tmp, "${1}-${2}")
	return strings.ToLower(tmp)
}
