package presentation

import (
	"syscall/js"

	"toolbox/src/common/elements"
	"toolbox/src/crypto/domain/entity"
	"toolbox/src/crypto/domain/services"
	"toolbox/src/crypto/usecase"
)

func StartServer() {
	s, ja := di()
	addHandler(s, ja)
	<-make(chan struct{})
}

func di() (*usecase.CryptoService, elements.IJsAdapter) {
	ds := services.NewCryptoDomainService()
	s := usecase.NewCryptoService(ds)
	ja := elements.GetJsAdapter()
	return s, ja
}

func addHandler(s *usecase.CryptoService, ja elements.IJsAdapter) {
	js.Global().Set("invokeEncryptAES", invokeEncryptAES(s, ja))
	js.Global().Set("invokeDecryptAES", invokeDecryptAES(s, ja))
}

func invokeEncryptAES(s *usecase.CryptoService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		key, err := ja.GetElementValue(args[1].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}

		result, err := s.EncryptAES(&entity.EncryptRequest{PlainText: text, Key: key})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse([]string{result.EncryptedText})
	})
}

func invokeDecryptAES(s *usecase.CryptoService, ja elements.IJsAdapter) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text, err := ja.GetElementValue(args[0].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		key, err := ja.GetElementValue(args[1].String())
		if err != nil {
			return ja.CreateErrorResponse(err)
		}

		result, err := s.DecryptAES(&entity.DecryptRequest{EncryptedText: text, Key: key})
		if err != nil {
			return ja.CreateErrorResponse(err)
		}
		return ja.CreateResponse([]string{result.PlainText})
	})
}
