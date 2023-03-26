package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) AddPaymentType(ctx *gin.Context) {
	var PaymentType models.PaymentType

	if err := ctx.ShouldBindJSON(&PaymentType); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.AddPaymentType(&PaymentType)
	if err != nil {
		log.Printf("%s in AddPaymentType(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "PaymentType added!")
}

func (h *Handler) GetPaymentType(ctx *gin.Context) {
	pagination := GeneratePaginationFromRequest(ctx)
	PAymentTypeList, err := h.Repository.GetPaymentType(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.Repository.TotalPagePaymentType(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = PAymentTypeList
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Handler) UpdatePaymentType(ctx *gin.Context) {
	var paymentType *models.PaymentType

	if err := ctx.ShouldBindJSON(&paymentType); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if paymentType.Name == "" {
		err := h.Repository.DeletePaymentType(paymentType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.Repository.UpdatePaymentType(paymentType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}
