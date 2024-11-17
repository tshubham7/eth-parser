package handler

import (
	"net/http"
	"strconv"

	"github.com/tshubham7/eth-parser/internal/pkg/constants"
	"github.com/tshubham7/eth-parser/internal/pkg/model"
)

var errorMap = map[string]model.ResponseError{
	constants.ErrorCodeInternalError: {
		Code:   constants.ErrorCodeInternalError,
		Title:  constants.ErrorInternalError,
		Detail: "something went wrong on our side",
		Status: strconv.Itoa(http.StatusInternalServerError),
	},
	constants.ErrorCodeExternalServerError: {
		Code:   constants.ErrorCodeExternalServerError,
		Title:  constants.ErrorInternalError,
		Detail: "something went wrong at external server side",
		Status: strconv.Itoa(http.StatusInternalServerError),
	},
}
