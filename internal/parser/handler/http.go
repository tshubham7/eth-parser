package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/tshubham7/eth-parser/internal/parser/usecase"
	"github.com/tshubham7/eth-parser/internal/pkg/constants"
	"github.com/tshubham7/eth-parser/internal/pkg/helpers"
	"github.com/tshubham7/eth-parser/internal/pkg/model"
	"github.com/tshubham7/eth-parser/internal/pkg/utils"
)

type parserHandler struct {
	manager usecase.ParserUsecase
}

func NewParserHttpHandler(manager usecase.ParserUsecase) *parserHandler {
	return &parserHandler{manager: manager}
}

func (h *parserHandler) RequestGetCurrentBlockNumber(w http.ResponseWriter, r *http.Request) {
	ctx, logg := utils.NewLogger(context.Background())
	logg.Info("got request for get current block")

	block, err := h.manager.GetCurrentBlockNumber(ctx)
	if err != nil {
		logg.Errorf("operation failed: %v", err)
		var errResp model.ResponseError
		var ok bool
		errResp, ok = errorMap[err.Error()]
		if !ok {
			errResp = errorMap[constants.ErrorCodeInternalError]
		}
		status, _ := strconv.Atoi(errResp.Status)
		helpers.RespondWithStatus(w, status, errResp)
		return
	}
	resp := model.ResponseData{Data: block}
	helpers.RespondWithStatus(w, http.StatusOK, resp)
}

func (h *parserHandler) RequestPostSubscribe(w http.ResponseWriter, r *http.Request) {
	ctx, logg := utils.NewLogger(context.Background())
	logg.Info("got request for post subcription")

	address := r.URL.Query().Get("address")
	if address == "" {
		helpers.RespondWithStatus(w, http.StatusBadRequest, model.ResponseError{
			Code:   constants.ErrorCodeBadRequest,
			Title:  constants.ErrorBadRequest,
			Detail: "invalid address provided",
			Status: strconv.Itoa(http.StatusBadGateway),
		})
		return
	}
	if err := h.manager.Subscribe(ctx, address); err != nil {
		logg.Errorf("operation failed: %v", err)
		var errResp model.ResponseError
		var ok bool
		errResp, ok = errorMap[err.Error()]
		if !ok {
			errResp = errorMap[constants.ErrorCodeInternalError]
		}
		status, _ := strconv.Atoi(errResp.Status)
		helpers.RespondWithStatus(w, status, errResp)
		return
	}
	helpers.RespondWithStatus(w, http.StatusOK, nil)
}

func (h *parserHandler) RequestGetTransactions(w http.ResponseWriter, r *http.Request) {
	ctx, logg := utils.NewLogger(context.Background())
	logg.Info("got request for get transactions")

	address := r.URL.Query().Get("address")
	if address == "" {
		helpers.RespondWithStatus(w, http.StatusBadRequest, model.ResponseError{
			Code:   constants.ErrorCodeBadRequest,
			Title:  constants.ErrorBadRequest,
			Detail: "invalid address provide",
			Status: strconv.Itoa(http.StatusBadGateway),
		})
		return
	}

	transactions, err := h.manager.GetTransactions(ctx, address)
	if err != nil {
		logg.Errorf("operation failed: %v", err)
		var errResp model.ResponseError
		var ok bool
		errResp, ok = errorMap[err.Error()]
		if !ok {
			errResp = errorMap[constants.ErrorCodeInternalError]
		}
		status, _ := strconv.Atoi(errResp.Status)
		helpers.RespondWithStatus(w, status, errResp)
		return
	}

	resp := model.ResponseData{Data: transactions}
	helpers.RespondWithStatus(w, http.StatusOK, resp)
}

// Added for testing purpose
func (h *parserHandler) RequestAllTransactions(w http.ResponseWriter, r *http.Request) {
	ctx, logg := utils.NewLogger(context.Background())
	logg.Info("got request for get transactions")

	transactions, err := h.manager.GetAllTransactions(ctx)
	if err != nil {
		logg.Errorf("operation failed: %v", err)
		var errResp model.ResponseError
		var ok bool
		errResp, ok = errorMap[err.Error()]
		if !ok {
			errResp = errorMap[constants.ErrorCodeInternalError]
		}
		status, _ := strconv.Atoi(errResp.Status)
		helpers.RespondWithStatus(w, status, errResp)
		return
	}

	resp := model.ResponseData{Data: transactions}
	helpers.RespondWithStatus(w, http.StatusOK, resp)
}
