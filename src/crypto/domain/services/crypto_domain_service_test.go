package services

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewCryptoDomainService(t *testing.T) {
	for _, testcase := range []struct {
		name string
		want *CryptoDomainService
	}{
		{
			name: "インスタンスが生成されること",
			want: &CryptoDomainService{},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			got := NewCryptoDomainService()
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", testcase.want) {
				t.Errorf("NewCryptoDomainService(): got %v, want %v", got, testcase.want)
			}
		})
	}
}

func Test_CryptoDomainService_EncryptAES(t *testing.T) {
	ignoreWant := "<ignore>"
	type args struct {
		plainText string
		key       string
	}
	type want struct {
		encryptedText string
		error         error
	}
	for _, testcase := range []struct {
		name string
		args args
		want want
	}{
		{
			name: "平文の文字列と暗号化キーを渡し、暗号化された文字列が返却されること",
			args: args{plainText: "PlainText", key: "Key"},
			want: want{encryptedText: ignoreWant, error: nil},
		},
		{
			name: "平文の文字列が空である場合、エラー情報が返却されること",
			args: args{plainText: "", key: "Key"},
			want: want{encryptedText: "", error: errors.New("plainText is empty.")},
		},
		{
			name: "暗号化キーが空である場合、エラー情報が返却されること",
			args: args{plainText: "PlainText", key: ""},
			want: want{encryptedText: "", error: errors.New("key is empty.")},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			ds := NewCryptoDomainService()
			gotVal, gotErr := ds.EncryptAES(testcase.args.plainText, testcase.args.key)
			if testcase.want.encryptedText != ignoreWant && gotVal != testcase.want.encryptedText {
				t.Errorf("CryptoDomainService.EncryptAES(): got %v, want %v", gotVal, testcase.want.encryptedText)
			}
			if fmt.Sprintf("%v", gotErr) != fmt.Sprintf("%v", testcase.want.error) {
				t.Errorf("CryptoDomainService.EncryptAES(): error %v, want %v", gotErr, testcase.want.error)
			}
		})
	}
}

func Test_CryptoDomainService_DecryptAES(t *testing.T) {
	type args struct {
		encryptedText string
		key           string
	}
	type want struct {
		plainText string
		error     error
	}
	for _, testcase := range []struct {
		name string
		args args
		want want
	}{
		{
			name: "暗号化された16進の文字列と暗号化キーを渡し、復号された文字列が返却されること",
			args: args{encryptedText: "439f21bb9ca932e28f8a776c8f8aa8b313abe53bc21a8db37cbd62cf91d93d770c71eee589", key: "Key"},
			want: want{plainText: "PlainText", error: nil},
		},
		{
			name: "暗号化された16進の文字列が空である場合、エラー情報が返却されること",
			args: args{encryptedText: "", key: "Key"},
			want: want{plainText: "", error: errors.New("encryptedText is empty.")},
		},
		{
			name: "暗号化キーが空である場合、エラー情報が返却されること",
			args: args{encryptedText: "439f21bb9ca932e28f8a776c8f8aa8b313abe53bc21a8db37cbd62cf91d93d770c71eee589", key: ""},
			want: want{plainText: "", error: errors.New("key is empty.")},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			ds := NewCryptoDomainService()
			gotVal, gotErr := ds.DecryptAES(testcase.args.encryptedText, testcase.args.key)
			if gotVal != testcase.want.plainText {
				t.Errorf("CryptoDomainService.DecryptAES(): got %v, want %v", gotVal, testcase.want.plainText)
			}
			if fmt.Sprintf("%v", gotErr) != fmt.Sprintf("%v", testcase.want.error) {
				t.Errorf("CryptoDomainService.DecryptAES(): error %v, want %v", gotErr, testcase.want.error)
			}
		})
	}
}
