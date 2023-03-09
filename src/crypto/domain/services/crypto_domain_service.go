package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"
)

// 参考
// Go で DB にいれるデータを暗号化したい
// https://kirikiriyamama.hatenablog.com/entry/2022/08/16/113409

type ICryptoDomainService interface {
	EncryptAES(string, string) (string, error)
	DecryptAES(string, string) (string, error)
}

type CryptoDomainService struct{}

func NewCryptoDomainService() *CryptoDomainService {
	return &CryptoDomainService{}
}

func (s *CryptoDomainService) EncryptAES(plainText string, key string) (string, error) {
	if plainText == "" {
		return "", errors.New("plainText is empty.")
	}
	if key == "" {
		return "", errors.New("key is empty.")
	}

	// 暗号化キーをSHA-256変換
	hashedKeyBytes := sha256.Sum256([]byte(key))

	// AES暗号化インターフェースの取得
	block, err := aes.NewCipher(hashedKeyBytes[:])
	if err != nil {
		return "", errors.WithStack(err)
	}

	// AES-GCMモードを指定
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.WithStack(err)
	}

	// 初期化ベクトルを生成
	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", errors.WithStack(err)
	}

	// テキストを暗号化し、初期化ベクトルを結合
	encryptedBytes := aead.Seal(nonce, nonce, []byte(plainText), nil)

	// 文字列に変換し返却
	encryptedText := fmt.Sprintf("%x", encryptedBytes)
	return encryptedText, nil
}

func (s *CryptoDomainService) DecryptAES(encryptedText string, key string) (string, error) {
	if encryptedText == "" {
		return "", errors.New("encryptedText is empty.")
	}
	if key == "" {
		return "", errors.New("key is empty.")
	}

	// 暗号化キーをSHA-256変換
	hashedKeyBytes := sha256.Sum256([]byte(key))

	// 文字列をバイト配列に変換
	encryptedBytes, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", errors.WithStack(err)
	}

	// AES暗号化インターフェースの取得
	block, err := aes.NewCipher(hashedKeyBytes[:])
	if err != nil {
		return "", errors.WithStack(err)
	}

	// AES-GCMモードを指定
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.WithStack(err)
	}

	// 初期化ベクトルと暗号化文字列に分離
	if len(encryptedBytes) < aead.NonceSize() {
		return "", errors.New("encryptedText is too short.")
	}
	nonce, cipherText := encryptedBytes[:aead.NonceSize()], encryptedBytes[aead.NonceSize():]

	// 暗号化文字列を復号
	plainText, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(plainText), nil
}
