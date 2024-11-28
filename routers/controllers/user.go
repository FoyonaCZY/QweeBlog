package controllers

import (
	"fmt"
	"github.com/FoyonaCZY/QweeBlog/service/user"
	"github.com/FoyonaCZY/QweeBlog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserRegister(c *gin.Context) {
	var request user.RegisterRequest
	if err := c.ShouldBindJSON(&request); err == nil {
		res, err := request.Register()
		if err != nil {
			util.Error(fmt.Sprintf("用户注册失败: %s", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("用户注册失败: %s", err.Error()),
			})
		} else {
			if res.ID == 0 {
				util.Info(fmt.Sprintf("用户注册失败, 参数错误: %s", request.Email))
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("用户注册失败, 参数错误: %s", request.Email),
				})
				return
			}
			util.Info(fmt.Sprintf("用户注册成功: %s", request.Email))
			c.JSON(http.StatusOK, res)
		}
	} else {
		util.Error(fmt.Sprintf("用户注册失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户注册失败，参数错误: %s", err.Error()),
		})
	}
}

func UserLogin(c *gin.Context) {
}

func UserUpdate(c *gin.Context) {
}

func UserInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		util.Error("用户信息获取失败，参数错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户信息获取失败，参数错误",
		})
		return
	}
	var request user.InfoRequest
	atoi, err := strconv.Atoi(id)
	if err != nil {
		util.Error(fmt.Sprintf("用户信息获取失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户信息获取失败，参数错误: %s", err.Error()),
		})
		return
	}
	request.ID = uint(atoi)
	res, err := request.Info()
	if err != nil {
		util.Error(fmt.Sprintf("用户信息获取失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("用户信息获取失败: %s", err.Error()),
		})
	} else {
		util.Info(fmt.Sprintf("用户信息获取成功: %s", id))
		c.JSON(http.StatusOK, res)
	}
}
