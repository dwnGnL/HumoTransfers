package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) AddVendor(ctx *gin.Context) {
	var Vendor models.Vendor

	if err := ctx.ShouldBindJSON(&Vendor); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	Vendor.Active = true
	err := h.Repository.AddVendor(&Vendor)
	if err != nil {
		log.Printf("%s in AddVendor(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Vendor added!")
}

func (h *Handler) GetVendor(ctx *gin.Context) {
	pagination := GeneratePaginationFromRequest(ctx)
	VendorList, err := h.Repository.GetVendor(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.Repository.TotalPageVendor(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = VendorList
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Handler) UpdateVendor(ctx *gin.Context) {
	var vendor *models.Vendor

	if err := ctx.ShouldBindJSON(&vendor); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if vendor.Name == "" {
		err := h.Repository.DeleteVendor(vendor)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.Repository.UpdateVendor(vendor)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) VendorStatus(ctx *gin.Context) {
	var vendor *models.Vendor

	if err := ctx.ShouldBindJSON(&vendor); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.VendorStatus(vendor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
