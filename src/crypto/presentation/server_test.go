package presentation

// import (
// 	"reflect"
// 	"testing"

// 	"toolbox/src/crypto/domain/services"
// 	"toolbox/src/crypto/usecase"
// )

// func Test_di(t *testing.T) {
// 	type want struct {
// 		value usecase.CryptoService
// 		error error
// 	}

// 	for _, testcase := range []struct {
// 		name string
// 		want want
// 	}{
// 		{
// 			name: "インスタンスを取得できること",
// 			want: want{
// 				value: *usecase.NewCryptoService(services.NewCryptoDomainService()),
// 				error: nil,
// 			},
// 		},
// 	} {
// 		t.Run(testcase.name, func(t *testing.T) {
// 			gotVal, gotVal1, gotErr := di()

// 			gotValue := reflect.ValueOf(gotVal).Elem()
// 			wantValue := reflect.ValueOf(testcase.want.value).Elem()

// 			for i := 0; i < gotValue.NumField(); i++ {
// 				gotType := gotValue.Field(i).Elem().Type().String()
// 				wantType := wantValue.Field(i).Elem().Type().String()
// 				if gotType != wantType {
// 					t.Errorf("di() = %v, want %v", gotType, wantType)
// 				}
// 			}
// 			if gotErr != testcase.want.error {
// 				t.Errorf("di(): error %v, want %v", gotErr, testcase.want.error)
// 			}
// 		})
// 	}
// }
