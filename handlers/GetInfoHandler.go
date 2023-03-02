package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetCountry(ctx *gin.Context) {
	//pageStr := ctx.Param("page")
	//limitStr := ctx.Param("limit")

	// todo need to consult
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	log.Println(pageStr, limitStr)
	if pageStr == "" {
		pageStr = "0"
	}
	if limitStr == "" {
		limitStr = "0"
	}

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	//log.Println(limit, page)
	offset := (page - 1) * 10
	app, len, err := h.Repository.GetCountries(offset, limit)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(len)
	ctx.JSON(http.StatusOK, app)
}
func (h *Handler) GetLanguage(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")

	if pageStr == "" {
		pageStr = "0"
	}
	if limitStr == "" {
		limitStr = "0"
	}
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetLanguages(offset, limit)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
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

func (h *Handler) GetCurrency(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetCurrency(offset, limit)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}

func (h *Handler) GetTest(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetTest(offset, limit)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}

func (h *Handler) GetAgent(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetAgent(offset, limit)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}

func (h *Handler) GetAccountAgent(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetAccountAgent(offset, limit)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}
