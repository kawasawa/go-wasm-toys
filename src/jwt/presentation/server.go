package presentation

import (
	"strconv"
	"syscall/js"

	"toolbox/src/common/elements"
	"toolbox/src/jwt/domain/entity"
	"toolbox/src/jwt/domain/services"
	"toolbox/src/jwt/usecase"
)

func StartServer() {
	s, ja := di()
	addHandler(s, ja)
	<-make(chan struct{})
}

func di() (*usecase.JwtService, elements.IJsAdapter) {
	ds := services.NewJwtDomainService()
	s := usecase.NewJwtService(ds)
	ja := elements.GetJsAdapter()
	return s, ja
}

func addHandler(s *usecase.JwtService, ja elements.IJsAdapter) {
	js.Global().Set("invokeDecodeJwt", invokeDecodeJwt(s, ja))
}

func invokeDecodeJwt(s *usecase.JwtService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		tokenString, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		key, err := ja.GetElementValue(args[1].String())
		if err != nil {
			key = ""
		}

		result, err := s.DecodeJwt(&entity.DecodeJwtRequest{TokenString: tokenString, Key: key})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse([]string{
			result.Header,
			result.Payload,
			strconv.FormatBool(result.Verified),
		})
	})
}
