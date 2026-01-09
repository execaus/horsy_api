package dto

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DefaultLimit = 10
	DefaultPage  = 1
)

type GetListParams struct {
	Limit  int32
	Page   int32
	Search string
}

func (p *GetListParams) BindFromContext(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", strconv.Itoa(DefaultLimit))
	pageStr := c.DefaultQuery("page", strconv.Itoa(DefaultPage))
	p.Search = c.DefaultQuery("search", "")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = DefaultLimit
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = DefaultPage
	}

	p.Limit = int32(limit)
	p.Page = int32(page)
}
