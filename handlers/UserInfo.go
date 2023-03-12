package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) AddUserInfo(ctx *gin.Context) {
	var UserInfo models.UserInfo

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	UserInfo.Active = true
	err := h.Repository.AddUserInfo(&UserInfo)
	if err != nil {
		log.Printf("%s in AddCountry", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Country added!")
}

func (h *Handler) GetUserInfo(ctx *gin.Context) {
	pagination := GeneratePaginationFromRequest(ctx)
	UserLists, err := h.Repository.GetUserInfo(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.Repository.TotalPageUserInfo(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = UserLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Handler) UpdateUserInfo(ctx *gin.Context) {
	var UserInfo *models.UserInfo

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if UserInfo.Icon == "" && UserInfo.Name == "" && UserInfo.Sort == 0 {
		err := h.Repository.DeleteUserInfo(UserInfo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.Repository.UpdateUserInfo(UserInfo)
		log.Println("work&")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}

	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) UserInfoStatus(ctx *gin.Context) {
	var UserInfo *models.UserInfo

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.UserInfoStatus(UserInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
