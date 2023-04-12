package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nishinoyama/kobuy-2/ent"
	"github.com/nishinoyama/kobuy-2/ent/user"
	"net/http"
)

type LedgerHandler struct {
	Client *ent.LedgerClient
}

func (h *LedgerHandler) GetAll(ctx *gin.Context) {
	cc := context.Background()
	l, err := h.Client.Query().
		WithReceiver(func(query *ent.UserQuery) {
			query.Select(user.FieldName)
		}).
		WithPayer(func(query *ent.UserQuery) {
			query.Select(user.FieldName)
		}).All(cc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, l)
}
