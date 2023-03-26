package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) AddServices(ctx *gin.Context) {
	var service models.Services

	if err := ctx.ShouldBindJSON(&service); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	service.Active = true
	err := h.Repository.AddService(&service)
	if err != nil {
		log.Printf("%s in AddService(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "AddService added!")
}

func (h *Handler) GetServices(ctx *gin.Context) {
	pagination := GeneratePaginationFromRequest(ctx)

	ServicesList, err := h.Repository.GetAccountAgent(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.Repository.TotalPageServices(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = ServicesList
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Handler) UpdateService(ctx *gin.Context) {
	var service *models.Services

	if err := ctx.ShouldBindJSON(&service); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if service.VendorId == 0 && service.Name == "" &&
		service.Type == "" {
		err := h.Repository.DeleteService(service)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	} else {
		err := h.Repository.UpdateService(service)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) ServiceStatus(ctx *gin.Context) {
	var services *models.Services

	if err := ctx.ShouldBindJSON(&services); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.ServicesStatus(services)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
