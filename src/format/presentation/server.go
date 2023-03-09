package presentation

import (
	"syscall/js"

	"toolbox/src/common/elements"
	"toolbox/src/format/domain/entity"
	"toolbox/src/format/domain/services"
	"toolbox/src/format/usecase"
)

func StartServer() {
	s, ja := di()
	addHandler(s, ja)
	<-make(chan struct{})
}

func di() (*usecase.FormatService, elements.IJsAdapter) {
	ds := services.NewFormatDomainService()
	s := usecase.NewFormatService(ds)
	ja := elements.GetJsAdapter()
	return s, ja
}

func addHandler(s *usecase.FormatService, ja elements.IJsAdapter) {
	js.Global().Set("invokeFormatJson", invokeFormatJson(s, ja))
}

func invokeFormatJson(s *usecase.FormatService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}

		result, err := s.FormatJson(&entity.FormatRequest{PlainText: text})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse([]string{result.FormattedText})
	})
}
