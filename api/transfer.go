package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/salman1s2h/simplebank/db/sqlc"
)

type TransferTxParams struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var reqBody TransferTxParams
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(400, "invalid request body")
		return
	}
	Currency := ctx.GetHeader("currency")

	fmt.Println("Currency in header: ", Currency)

	if !server.validAccount(ctx, reqBody.FromAccountID, reqBody.Currency) {
		return
	}

	if !server.validAccount(ctx, reqBody.ToAccountID, reqBody.Currency) {
		return
	}
	arg := db.TransferTxParams{
		FromAccountID: reqBody.FromAccountID,
		ToAccountID:   reqBody.ToAccountID,
		Amount:        reqBody.Amount,
	}
	rst, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, rst)
}

// func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
// 	account, err := server.store.GetAccountByID(ctx, accountID)
// 	if err != nil {
// 		if err == db.ErrNoRows {
// 			ctx.JSON(404, "account not found")
// 			return false
// 		}

// 		ctx.JSON(http.StatusInternalServerError, errResponse(err))
// 		return false

// 	}
// 	if account.Currency != currency {
// 		err := fmt.Errorf("account %d currency mismatch: %s vs %s", accountID, account.Currency, currency)
// 		ctx.JSON(http.StatusBadRequest, errResponse(err))
// 		return false
// 	}
// }

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccountByID(ctx, accountID)
	fmt.Println("account     :  ", account)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account %d currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
