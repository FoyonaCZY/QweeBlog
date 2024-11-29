package controllers

import (
	"fmt"
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/auth"
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/FoyonaCZY/QweeBlog/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigUpdate(c *gin.Context) {
	//验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("配置更改失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("配置更改失败: %s", err.Error()),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin {
		util.Error(fmt.Sprintf("配置更改失败: %s", "权限不足"))
		c.JSON(http.StatusForbidden, gin.H{
			"message": "权限不足",
		})
		return
	}

	var request config.Config
	if err := c.ShouldBindJSON(&request); err == nil {
		*config.Configs = request
		go config.UpdateConfig()
	} else {
		util.Error(fmt.Sprintf("配置更改失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("配置更改失败: %s", err.Error()),
		})
		return
	}
}
