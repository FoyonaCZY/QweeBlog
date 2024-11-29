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
	//验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("用户信息更改失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("用户信息更改失败: %s", err.Error()),
		})
		return
	}

	var request user.UpdateRequest
	if err := c.ShouldBindJSON(&request); err == nil {
		//验证请求者权限
		if (reqUser.ID != request.ID || reqUser.GroupID != request.GroupID) && reqUser.Group.Type != models.GroupTypeAdmin {
			util.Error(fmt.Sprintf("用户信息更改失败，权限不足"))
			c.JSON(http.StatusForbidden, gin.H{
				"message": "用户信息更改失败，权限不足",
			})
			return
		}

		if request.ID == 1 && request.GroupID != 1 {
			util.Error(fmt.Sprintf("用户信息更改失败，无法更改初始管理员的用户组"))
			c.JSON(http.StatusForbidden, gin.H{
				"message": "用户信息更改失败，无法更改初始管理员的用户组",
			})
			return
		}

		res, err := request.Update()
		if err != nil {
			util.Error(fmt.Sprintf("用户信息更改失败: %s", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("用户信息更改失败: %s", err.Error()),
			})
		} else {
			util.Info(fmt.Sprintf("用户信息更改成功: %s", request.Email))
			c.JSON(http.StatusOK, res)
		}
	} else {
		util.Error(fmt.Sprintf("用户信息更改失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户信息更改失败，参数错误: %s", err.Error()),
		})
	}
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

func UserDelete(c *gin.Context) {
	//验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("用户信息删除失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("用户信息删除失败: %s", err.Error()),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin {
		util.Error(fmt.Sprintf("用户信息删除失败，权限不足"))
		c.JSON(http.StatusForbidden, gin.H{
			"message": "用户信息删除失败，权限不足",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if id == 1 {
		util.Error(fmt.Sprintf("用户信息删除失败，无法删除初始管理员"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户信息删除失败，无法删除初始管理员",
		})
		return
	}
	if err != nil {
		util.Error(fmt.Sprintf("用户信息删除失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户信息删除失败，参数错误: %s", err.Error()),
		})
		return
	}
	request := user.DeleteRequest{ID: uint(id)}
	res, err := request.Delete()
	if err != nil {
		util.Error(fmt.Sprintf("用户信息删除失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("用户信息删除失败: %s", err.Error()),
		})
	} else {
		util.Info(fmt.Sprintf("用户信息删除成功: %d", res.ID))
		c.JSON(http.StatusOK, res)
	}
}

func UserList(c *gin.Context) {
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

	res, err := user.List()
	if err != nil {
		util.Error(fmt.Sprintf("用户信息获取失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("用户信息获取失败: %s", err.Error()),
		})
	} else {
		util.Info(fmt.Sprintf("用户信息获取成功"))
		c.JSON(http.StatusOK, res)
	}
}

func UserActivate(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		util.Error(fmt.Sprintf("用户激活失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户激活失败，参数错误: %s", err.Error()),
		})
		return
	}
	token := c.Query("token")

	reqUser, err := models.GetUserByID(uint(id))
	if err != nil {
		util.Error(fmt.Sprintf("用户激活失败，用户不存在: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户激活失败，用户不存在: %s", err.Error()),
		})
		return
	}

	if reqUser.Status != models.UserStatusNotActive {
		util.Error(fmt.Sprintf("用户激活失败，用户已激活或被封禁"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户激活失败，用户已激活或被封禁",
		})
		return
	}

	if reqUser.ActivationToken != token {
		util.Error(fmt.Sprintf("用户激活失败，激活码错误"))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户激活失败，激活码错误",
		})
		return
	}

	reqUser.Status = models.UserStatusActive
	err = models.DB.Save(&reqUser).Error
	if err != nil {
		util.Error(fmt.Sprintf("用户激活失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("用户激活失败: %s", err.Error()),
		})
	} else {
		util.Info(fmt.Sprintf("用户激活成功: %d", id))
		c.JSON(http.StatusOK, gin.H{
			"message": "用户激活成功，请重新登录",
		})
	}
}
