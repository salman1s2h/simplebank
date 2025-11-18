package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/salman1s2h/simplebank/db/sqlc"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest

	// Use ctx, not router
	fmt.Printf("%v\n", ctx)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	fmt.Println(arg)
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) getAccount(ctx *gin.Context) {
	fmt.Println("Request received at /ping")
	idParam := ctx.Param("id")
	var id int64
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccountByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// func (server *Server) listAccounts(ctx *gin.Context) {
// 	// Placeholder implementation
// 	ctx.JSON(http.StatusOK, gin.H{"message": "listAccounts endpoint"})
// }

func (server *Server) listAccounts(ctx *gin.Context) {
	pageStr := ctx.GetHeader("page")
	limitStr := ctx.GetHeader("limit")

	page, err1 := strconv.Atoi(pageStr)
	limit, err2 := strconv.Atoi(limitStr)
	if err1 != nil || err2 != nil || page < 1 || limit < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	if limit > 12 {
		limit = 12
	}

	offset := (page - 1) * limit

	fmt.Println("limit:", limit)
	fmt.Println("offset:", offset)
	fmt.Println("page:", page)

	args := db.GetAccountsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	accounts, err := server.store.GetAccounts(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accounts"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"page":    page,
		"limit":   limit,
		"results": len(accounts),
		"data":    accounts,
	})
}
