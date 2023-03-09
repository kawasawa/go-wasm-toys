package presentation

import (
	"syscall/js"

	"toolbox/src/common/elements"
	"toolbox/src/hash/domain/entity"
	"toolbox/src/hash/domain/services"
	"toolbox/src/hash/usecase"
)

func StartServer() {
	s, ja := di()
	addHandler(s, ja)
	<-make(chan struct{})
}

func di() (*usecase.HashService, elements.IJsAdapter) {
	ds := services.NewHashDomainService()
	s := usecase.NewHashService(ds)
	ja := elements.GetJsAdapter()
	return s, ja
}

func addHandler(s *usecase.HashService, ja elements.IJsAdapter) {
	js.Global().Set("invokeHash", invokeHash(s, ja))
}

func invokeHash(s *usecase.HashService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}

		result, err := s.Hash(&entity.HashRequest{PlainText: text})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse([]string{
			result.MD5,
			result.SHA1,
			result.SHA256,
			result.SHA384,
			result.SHA512,
		})
	})
}
