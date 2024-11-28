package controllers

import (
	"fmt"
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/auth"
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
	var request user.LoginRequest
	if err := c.ShouldBindJSON(&request); err == nil {
		res, err := request.Login()
		if err != nil {
			util.Error(fmt.Sprintf("用户登录失败: %s", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("用户登录失败: %s", err.Error()),
			})
		} else {
			util.Info(fmt.Sprintf("用户登录成功: %s", request.Email))
			c.JSON(http.StatusOK, res)
		}
	} else {
		util.Error(fmt.Sprintf("用户登录失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户登录失败，参数错误: %s", err.Error()),
		})
	}
}

func UserUpdate(c *gin.Context) {
}

func UserInfo(c *gin.Context) {
	//验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("用户信息获取失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("用户信息获取失败: %s", err.Error()),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin {
		util.Error(fmt.Sprintf("用户信息获取失败，权限不足"))
		c.JSON(http.StatusForbidden, gin.H{
			"message": "用户信息获取失败，权限不足",
		})
		return
	}

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
