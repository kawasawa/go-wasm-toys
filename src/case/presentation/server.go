package presentation

import (
	"syscall/js"

	"toolbox/src/case/domain/entity"
	"toolbox/src/case/domain/services"
	"toolbox/src/case/usecase"
	"toolbox/src/common/elements"
)

func StartServer() {
	s, ja := di()
	addHandler(s, ja)
	<-make(chan struct{})
}

func di() (*usecase.CaseService, elements.IJsAdapter) {
	ds := services.NewCaseDomainService()
	s := usecase.NewCaseService(ds)
	ja := elements.GetJsAdapter()
	return s, ja
}

func addHandler(s *usecase.CaseService, ja elements.IJsAdapter) {
	js.Global().Set("invokeConvertCase", invokeConvertCase(s, ja))
}

func invokeConvertCase(s *usecase.CaseService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}

		result, err := s.ConvertCase(&entity.CaseRequest{PlainText: text})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse([]string{
			result.Camel,
			result.Pascal,
			result.Snake,
			result.Kebab,
			result.Upper,
			result.Lower,
		})
	})
}
