package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/pkg/controller"
	"net/http"
)

type LedgerHandler struct {
	LedgerController *controller.LedgerController
}

func NewLedgerHandler(r *gin.RouterGroup, lc *controller.LedgerController) {
	handler := LedgerHandler{LedgerController: lc}

	ledger := r.Group("/ledger")
	{
		ledger.GET("", handler.GetAll)
	}
}

func (h *LedgerHandler) GetAll(ctx *gin.Context) {
	res, err := h.LedgerController.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}
