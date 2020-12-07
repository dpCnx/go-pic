package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
	"web/models"
)

// GetPic godoc
// @Summary 获取图片
// @Description 获取图片
// @Tags 获取图片
// @ID /getpic
// @Accept  json
// @Produce  json
// @Param page query int true "当前页数"
// @Param page_size query int true "每页个数"
// @Success 1000 {object} models.ResponseData "success"
// @Router /getpic [get]
func GetPic(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		zap.L().Error(err.Error())
		models.ResponseError(c, models.CodeInvalidParams)
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		zap.L().Error(err.Error())
		models.ResponseError(c, models.CodeInvalidParams)
		return
	}

	if page == 2 {
		time.Sleep(800 * time.Millisecond)
	}

	p := &models.Pic{}
	res, err := p.QuarkPic(page, pageSize)
	if err != nil {
		zap.L().Error(err.Error())
		models.ResponseError(c, models.CodeServerBusy)
		return
	}

	models.ResponseSuccess(c, res)
}
