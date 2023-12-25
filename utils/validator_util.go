package utils

import "github.com/go-playground/validator"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type Response struct {
	Success bool             `json:"success"`
	Message string           `json:"message,omitempty" default:""`
	Data    any              `json:"data,omitempty"`
	Error   []*ErrorResponse `json:"errors,omitempty"`
}

var validate = validator.New()

// func ValidateStruct[T any](vT T) []*ErrorResponse {
// 	errors := []*ErrorResponse{}
// 	err := validate.Struct(vT)
// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			element := &ErrorResponse{}
// 			element.FailedField = err.StructNamespace()
// 			element.Tag = err.Tag()
// 			element.Value = err.Param()
// 			errors = append(errors, element)
// 		}
// 	}
// 	if len(errors) == 0 {
// 		return nil
// 	}
// 	return errors
// }

func Resp(s bool, m string, d any, e []*ErrorResponse) *Response {
	return &Response{
		Success: s,
		Message: m,
		Data:    d,
		Error:   e,
	}
}
