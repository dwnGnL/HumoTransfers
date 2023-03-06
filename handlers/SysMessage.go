package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) AddSysMessage(ctx *gin.Context) {
	var message models.SysMessage

	if err := ctx.ShouldBindJSON(&message); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	message.Active = true
	err := h.Repository.AddSysMessage(&message)
	if err != nil {
		log.Printf("%s in AddSysMessage(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "SysMessage added!")
}

func (h *Handler) GetSysMessage(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetSysMessage(offset, limit)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}

func (h Handler) UpdateSysMessage(ctx *gin.Context) {
	var message *models.SysMessage

	if err := ctx.ShouldBindJSON(&message); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if message.Name == "" {
		err := h.Repository.DeleteSysMessage(message)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.Repository.UpdateSysMessage(message)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}

	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) SysMessageStatus(ctx *gin.Context) {
	var message *models.SysMessage

	if err := ctx.ShouldBindJSON(&message); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.SysMessageStatus(message)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
