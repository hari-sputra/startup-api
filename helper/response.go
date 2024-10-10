package helper

import "github.com/go-playground/validator/v10"

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	responseJ := response{
		Meta: meta,
		Data: data,
	}

	return responseJ

}

func ErrorValidationFormatter(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {

		var errMsg string

		switch e.Field() {
		case "Name":
			errMsg = "Name is required!"
		case "Occupation":
			errMsg = "Occupation is required!"
		case "Email":
			if e.Tag() == "required" {
				errMsg = "Email is required!"
			} else if e.Tag() == "email" {
				errMsg = "Email format not valid!"
			}
		case "Password":
			errMsg = "Password is Required"
		default:
			errMsg = e.Error()
		}

		errors = append(errors, errMsg)
	}

	return errors
}
