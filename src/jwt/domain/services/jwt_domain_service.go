package services

import (
	"encoding/base64"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type IJwtDomainService interface {
	DecodeJwt(string, string) (string, string, bool, error)
}

type JwtDomainService struct{}

func NewJwtDomainService() *JwtDomainService {
	return &JwtDomainService{}
}

func (s *JwtDomainService) DecodeJwt(tokenString string, key string) (string, string, bool, error) {
	if tokenString == "" {
		return "", "", false, errors.New("token is empty.")
	}

	// ヘッダ、ペイロード、シグニチャに分割
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return "", "", false, errors.New("invalid token format.")
	}

	// ヘッダ、ペイロードを Base64 でデコード
	header, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", "", false, errors.New("failed to decode header.")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", "", false, errors.New("failed to decode payload.")
	}

	// シグニチャを渡された暗号鍵で検証
	// シグニチャ部はヘッダ〜ペイロード部を暗号鍵で暗号化したものと一致する
	verified := false
	if key != "" {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if token.Valid && err == nil {
			verified = true
		}
	}

	return string(header), string(payload), verified, err
}
