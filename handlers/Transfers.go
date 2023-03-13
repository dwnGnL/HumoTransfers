package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	pagination := GeneratePaginationFromRequest(ctx)
	TransferLists, err := h.Repository.GetTransfer(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.Repository.TotalPageTransfer(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = TransferLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
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
