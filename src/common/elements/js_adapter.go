package elements

import (
	"errors"
	"fmt"
	"strconv"
	"syscall/js"
)

type IJsAdapter interface {
	GetElementValue(string) (string, error)
	CreateResponse([]string) map[string]interface{}
	CreateErrorResponse(error) map[string]interface{}
}

type JsAdapter struct{}

var adapter IJsAdapter = newJsAdapter()

func newJsAdapter() *JsAdapter {
	return &JsAdapter{}
}

func GetJsAdapter() IJsAdapter {
	return adapter
}

func (a JsAdapter) GetElementValue(elementID string) (string, error) {
	document := js.Global().Get("document")
	if !document.Truthy() {
		return "", errors.New("failed to get document object")
	}
	element := document.Call("getElementById", elementID)
	if !element.Truthy() {
		return "", fmt.Errorf("failed to getElementById: %s", elementID)
	}
	value := element.Get("value")
	if !value.Truthy() {
		return "", fmt.Errorf("failed to Get value: %s", elementID)
	}
	return value.String(), nil
}

func (a JsAdapter) CreateResponse(value []string) map[string]interface{} {
	response := map[string]interface{}{}
	response["count"] = len(value)
	for i, v := range value {
		response["value"+strconv.Itoa(i)] = v
	}
	return response
}

func (a JsAdapter) CreateErrorResponse(err error) map[string]interface{} {
	return map[string]interface{}{"error": err.Error()}
}
