package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) AddTransfer(ctx *gin.Context) {
	var transfers models.Transfers

	if err := ctx.ShouldBindJSON(&transfers); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.AddTransfer(&transfers)

	if err != nil {
		log.Printf("%s in AddTest(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Transfer added!")
}

func (h *Handler) GetTransfer(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetTransfer(offset, limit)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, app)
}

func (h Handler) UpdateTransfer(ctx *gin.Context) {
	var transfer *models.Transfers

	if err := ctx.ShouldBindJSON(&transfer); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if transfer.EntityId == 0 && transfer.Entity == "" &&
		transfer.LangId == 0 && transfer.Value == "" {
		err := h.Repository.DeleteTransfers(transfer)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.Repository.UpdateTransfers(transfer)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}
