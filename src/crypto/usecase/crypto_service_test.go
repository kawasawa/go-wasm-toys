package usecase

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"toolbox/src/common/logging"
	"toolbox/src/crypto/domain/entity"
	"toolbox/src/crypto/domain/services"
)

func TestNewCryptoService(t *testing.T) {
	ds := CryptoDomainServiceMock{}
	type args struct {
		DomainService services.ICryptoDomainService
	}
	for _, testcase := range []struct {
		name string
		args args
		want *CryptoService
	}{
		{
			name: "インスタンスが生成されること",
			args: args{DomainService: &ds},
			want: &CryptoService{DomainService: &ds, Logger: logging.GetLogger()},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			got := NewCryptoService(testcase.args.DomainService)
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", testcase.want) {
				t.Errorf("NewCryptoService(): got %v, want %v", got, testcase.want)
			}
			if !reflect.DeepEqual(got.DomainService, testcase.want.DomainService) {
				t.Errorf("NewCryptoService() DomainService: got %v, want %v", got.DomainService, testcase.want.DomainService)
			}
		})
	}
}

func Test_CryptoService_EncryptAES(t *testing.T) {
	type fields struct {
		DomainService services.ICryptoDomainService
	}
	type args struct {
		req *entity.EncryptRequest
	}
	type want struct {
		value *entity.EncryptResponse
		error error
	}
	for _, testcase := range []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "平文の文字列と暗号化キーを渡し、暗号化された文字列が返却されること",
			fields: fields{DomainService: &CryptoDomainServiceMock{
				EncryptAESMock: func(string, string) (string, error) {
					return "EncryptedText", nil
				},
			}},
			args: args{req: &entity.EncryptRequest{PlainText: "PlainText", Key: "Key"}},
			want: want{value: &entity.EncryptResponse{EncryptedText: "EncryptedText"}, error: nil},
		},
		{
			name: "暗号化処理でエラーが発生した場合、そのエラー情報が返却されること",
			fields: fields{DomainService: &CryptoDomainServiceMock{
				EncryptAESMock: func(string, string) (string, error) {
					return "", errors.New("Encryption error")
				},
			}},
			args: args{req: &entity.EncryptRequest{PlainText: "PlainText", Key: "Key"}},
			want: want{value: &entity.EncryptResponse{EncryptedText: ""}, error: errors.New("Encryption error")},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			s := NewCryptoService(testcase.fields.DomainService)
			gotVal, gotErr := s.EncryptAES(testcase.args.req)
			if !reflect.DeepEqual(gotVal, testcase.want.value) {
				t.Errorf("CryptoService.EncryptAES(): got %v, want %v", gotVal, testcase.want.value)
			}
			if fmt.Sprintf("%v", gotErr) != fmt.Sprintf("%v", testcase.want.error) {
				t.Errorf("CryptoService.EncryptAES(): error %v, want %v", gotErr, testcase.want.error)
			}
		})
	}
}

func Test_CryptoService_DecryptAES(t *testing.T) {
	type fields struct {
		DomainService services.ICryptoDomainService
	}
	type args struct {
		req *entity.DecryptRequest
	}
	type want struct {
		value *entity.DecryptResponse
		error error
	}
	for _, testcase := range []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "暗号化された文字列と暗号化キーを渡し、復号された文字列が返却されること",
			fields: fields{DomainService: &CryptoDomainServiceMock{
				DecryptAESMock: func(string, string) (string, error) {
					return "PlainText", nil
				},
			}},
			args: args{req: &entity.DecryptRequest{EncryptedText: "EncryptedText", Key: "Key"}},
			want: want{value: &entity.DecryptResponse{PlainText: "PlainText"}, error: nil},
		},
		{
			name: "復号処理でエラーが発生した場合、そのエラー情報が返却されること",
			fields: fields{DomainService: &CryptoDomainServiceMock{
				DecryptAESMock: func(string, string) (string, error) {
					return "", errors.New("Decryption error")
				},
			}},
			args: args{req: &entity.DecryptRequest{EncryptedText: "EncryptedText", Key: "Key"}},
			want: want{value: &entity.DecryptResponse{PlainText: ""}, error: errors.New("Decryption error")},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			s := NewCryptoService(testcase.fields.DomainService)
			gotVal, gotErr := s.DecryptAES(testcase.args.req)
			if !reflect.DeepEqual(gotVal, testcase.want.value) {
				t.Errorf("CryptoService.DecryptAES(): got %v, want %v", gotVal, testcase.want.value)
			}
			if fmt.Sprintf("%v", gotErr) != fmt.Sprintf("%v", testcase.want.error) {
				t.Errorf("CryptoService.DecryptAES(): error %v, want %v", gotErr, testcase.want.error)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Mock
// -----------------------------------------------------------------------------

type CryptoDomainServiceMock struct {
	EncryptAESMock func(string, string) (string, error)
	DecryptAESMock func(string, string) (string, error)
}

func (s *CryptoDomainServiceMock) EncryptAES(plainText string, key string) (string, error) {
	return s.EncryptAESMock(plainText, key)
}

func (s *CryptoDomainServiceMock) DecryptAES(encryptedText string, key string) (string, error) {
	return s.DecryptAESMock(encryptedText, key)
}
