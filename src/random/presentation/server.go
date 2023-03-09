package presentation

import (
	"strconv"
	"syscall/js"

	"toolbox/src/common/elements"
	"toolbox/src/random/domain/entity"
	"toolbox/src/random/domain/services"
	"toolbox/src/random/usecase"
)

func StartServer() {
	s, ja := di()
	addHandler(s, ja)
	<-make(chan struct{})
}

func di() (*usecase.RandomService, elements.IJsAdapter) {
	ds := services.NewRandomDomainService()
	s := usecase.NewRandomService(ds)
	ja := elements.GetJsAdapter()
	return s, ja
}

func addHandler(s *usecase.RandomService, ja elements.IJsAdapter) {
	js.Global().Set("invokeGenerateRandom", invokeGenerateRandom(s, ja))
}

func invokeGenerateRandom(s *usecase.RandomService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		characters, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		temp, err := ja.GetElementValue(args[1].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		length, err := strconv.Atoi(temp)
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		temp, err = ja.GetElementValue(args[2].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		count, err := strconv.Atoi(temp)
		if err != nil {
			return ja.CreateErrorResponse(err)
		}

		result, err := s.GenerateRandom(&entity.GenerateRandomRequest{Characters: characters, Length: length, Count: count})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse(result.Randoms)
	})
}
