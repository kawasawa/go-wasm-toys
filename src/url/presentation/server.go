package presentation

import (
	"syscall/js"

	"toolbox/src/common/elements"
	"toolbox/src/url/domain/entity"
	"toolbox/src/url/domain/services"
	"toolbox/src/url/usecase"
)

func StartServer() {
	s, ja := di()
	addHandler(s, ja)
	<-make(chan struct{})
}

func di() (*usecase.UrlService, elements.IJsAdapter) {
	ds := services.NewUrlDomainService()
	s := usecase.NewUrlService(ds)
	ja := elements.GetJsAdapter()
	return s, ja
}

func addHandler(s *usecase.UrlService, ja elements.IJsAdapter) {
	js.Global().Set("invokeEncodeUrl", invokeEncodeUrl(s, ja))
	js.Global().Set("invokeDecodeUrl", invokeDecodeUrl(s, ja))
}

func invokeEncodeUrl(s *usecase.UrlService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}

		result, err := s.EncodeUrl(&entity.EncodeUrlRequest{UrlStr: text})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse([]string{result.EncodedUrl})
	})
}

func invokeDecodeUrl(s *usecase.UrlService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}

		result, err := s.DecodeUrl(&entity.DecodeUrlRequest{UrlStr: text})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse([]string{result.DecodedUrl})
	})
}
