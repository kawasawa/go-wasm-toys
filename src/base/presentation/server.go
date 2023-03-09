package presentation

import (
	"syscall/js"

	"toolbox/src/base/domain/entity"
	"toolbox/src/base/domain/services"
	"toolbox/src/base/usecase"
	"toolbox/src/common/elements"
)

func StartServer() {
	s, ja := di()
	addHandler(s, ja)
	<-make(chan struct{})
}

func di() (*usecase.BaseService, elements.IJsAdapter) {
	ds := services.NewBaseDomainService()
	s := usecase.NewBaseService(ds)
	ja := elements.GetJsAdapter()
	return s, ja
}

func addHandler(s *usecase.BaseService, ja elements.IJsAdapter) {
	js.Global().Set("invokeEncodeBase", invokeEncodeBase(s, ja))
	js.Global().Set("invokeDecodeBase", invokeDecodeBase(s, ja))
}

func invokeEncodeBase(s *usecase.BaseService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}

		result, err := s.EncodeBase(&entity.EncodeBaseRequest{PlainText: text})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse([]string{
			result.Base32,
			result.Base64,
		})
	})
}

func invokeDecodeBase(s *usecase.BaseService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}

		result, err := s.DecodeBase(&entity.DecodeBaseRequest{BaseText: text})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse([]string{
			result.Plain32,
			result.Plain64,
		})
	})
}
