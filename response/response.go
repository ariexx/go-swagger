package response

import "reflect"

type EmptyObject map[string]string

type ResponseJsonSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseFail struct {
	Status   string `json:"status" default:"fail"`
	Data     any    `json:"data"`
	DataFail any    `json:"data_fail"`
	Message  string `json:"message"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message any    `json:"message"`
}

func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Array, reflect.Chan, reflect.Map, reflect.Interface:
		return reflect.ValueOf(i).IsNil()
	default:
		return false
	}
}

func isArrayEmpty(i interface{}) bool {
	if reflect.TypeOf(i).Kind() == reflect.Slice {
		return reflect.ValueOf(i).Len() == 0
	}
	return false
}

// ApiResponseSuccess is a function to return response success
func ApiResponseSuccess(status string, data any) *ResponseJsonSuccess {
	//return empty string if data is nil
	if isNilFixed(data) {
		data = EmptyObject{}
	}
	//return array if data is empty
	if isArrayEmpty(data) {
		data = []EmptyObject{}
	}
	return &ResponseJsonSuccess{
		Status: status,
		Data:   data,
	}
}

// ApiResponseError is a function to return response error
func ApiResponseError[T any](message T, code string) *ResponseError {
	return &ResponseError{
		Status:  "error",
		Code:    code,
		Message: message,
	}
}

// ApiResponseFail is a function to return response fail
func ApiResponseFail(data any, message string) *ResponseFail {
	if isNilFixed(data) {
		data = EmptyObject{}
	}

	if isArrayEmpty(data) {
		data = []EmptyObject{}
	}

	return &ResponseFail{
		Status:   "fail",
		Data:     data,
		DataFail: data,
		Message:  message,
	}
}

type ProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
