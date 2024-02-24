package response

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Code    int               `json:"code"`
	Message map[string]string `json:"message"`
}

func Error(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResponse := ErrorResponse{
		Code:    code,
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResponse)
}

func OK(w http.ResponseWriter) {
	Error(w, http.StatusOK, http.StatusText(http.StatusOK))
}

func Created(w http.ResponseWriter) {
	Error(w, http.StatusCreated, http.StatusText(http.StatusCreated))
}

func NotFound(w http.ResponseWriter) {
	Error(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func InternalError(w http.ResponseWriter) {
	Error(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func BadRequest(w http.ResponseWriter) {
	Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
}

func Unauthorized(w http.ResponseWriter) {
	Error(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func Forbidden(w http.ResponseWriter) {
	Error(w, http.StatusForbidden, http.StatusText(http.StatusForbidden))
}

func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(data)
}

func ValidationErrs(w http.ResponseWriter, errors validator.ValidationErrors) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity) // 422 Unprocessable Entity
	errorMap := make(map[string]string)

	for _, e := range errors {
		errorMap[e.Field()] = e.Tag()
	}

	response := ValidationErrorResponse{
		Code:    http.StatusUnprocessableEntity,
		Message: errorMap,
	}

	json.NewEncoder(w).Encode(response)
}
