package controllers

import (
	"fmt"
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/auth"
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/FoyonaCZY/QweeBlog/service/group"
	"github.com/FoyonaCZY/QweeBlog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GroupCreate(c *gin.Context) {
	// 验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("用户组创建失败: %s", err.Error()))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("用户组创建失败: %s", err.Error()),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin {
		util.Error(fmt.Sprintf("用户组创建失败: %s", "权限不足"))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("用户组创建失败: %s", "权限不足"),
		})
		return
	}

	var request group.CreateRequest
	if err := c.ShouldBindJSON(&request); err == nil {
		res, err := request.Create()
		if err != nil {
			util.Error(fmt.Sprintf("用户组创建失败: %s", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("用户组创建失败: %s", err.Error()),
			})
		} else {
			util.Info(fmt.Sprintf("用户组创建成功: %s", request.Name))
			c.JSON(http.StatusOK, res)
		}
	} else {
		util.Error(fmt.Sprintf("用户组创建失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户组创建失败，参数错误: %s", err.Error()),
		})
	}
}

func GroupUpdate(c *gin.Context) {
	// 验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("用户组更新失败: %s", err.Error()))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("用户组更新失败: %s", err.Error()),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin {
		util.Error(fmt.Sprintf("用户组更新失败: %s", "权限不足"))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("用户组更新失败: %s", "权限不足"),
		})
		return
	}

	var request group.UpdateRequest
	if err := c.ShouldBindJSON(&request); err == nil {
		//不允许修改初始管理员组和默认注册组
		if request.ID == 1 || request.ID == config.Configs.DefaultGroup.ID {
			util.Error(fmt.Sprintf("用户组更新失败: %s", "不允许修改初始管理员组或默认注册组"))
			c.JSON(http.StatusForbidden, gin.H{
				"message": fmt.Sprintf("用户组更新失败: %s", "不允许修改初始管理员组或默认注册组"),
			})
			return
		}

		res, err := request.Update()
		if err != nil {
			util.Error(fmt.Sprintf("用户组更新失败: %s", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("用户组更新失败: %s", err.Error()),
			})
		} else {
			util.Info(fmt.Sprintf("用户组更新成功: %s", request.Name))
			c.JSON(http.StatusOK, res)
		}
	} else {
		util.Error(fmt.Sprintf("用户组更新失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户组更新失败，参数错误: %s", err.Error()),
		})
	}
}

func GroupDelete(c *gin.Context) {
	// 验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("用户组删除失败: %s", err.Error()))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("用户组删除失败: %s", err.Error()),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin {
		util.Error(fmt.Sprintf("用户组删除失败: %s", "权限不足"))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("用户组删除失败: %s", "权限不足"),
		})
		return
	}

	var request group.DeleteRequest
	atoi, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.Error(fmt.Sprintf("用户组删除失败: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户组删除失败: %s", err.Error()),
		})
		return
	}
	request.ID = uint(atoi)
	//不允许删除初始管理员组和默认注册组
	if request.ID == 1 || request.ID == config.Configs.DefaultGroup.ID {
		util.Error(fmt.Sprintf("用户组删除失败: %s", "不允许删除初始管理员组或默认注册组"))
		c.JSON(http.StatusForbidden, gin.H{
			"message": fmt.Sprintf("用户组删除失败: %s", "不允许删除初始管理员组或默认注册组"),
		})
		return
	}

	res, err := request.Delete()
	if err != nil {
		util.Error(fmt.Sprintf("用户组删除失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("用户组删除失败: %s", err.Error()),
		})
	} else {
		util.Info(fmt.Sprintf("用户组删除成功: %d", request.ID))
		c.JSON(http.StatusOK, res)
	}
}

func GroupList(c *gin.Context) {
	// 验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("用户组列表获取失败: %s", err.Error()))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("用户组列表获取失败: %s", err.Error()),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin {
		util.Error(fmt.Sprintf("用户组列表获取失败: %s", "权限不足"))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("用户组列表获取失败: %s", "权限不足"),
		})
		return
	}

	groups, err := group.List()
	if err != nil {
		util.Error(fmt.Sprintf("用户组列表获取失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("用户组列表获取失败: %s", err.Error()),
		})
	} else {
		util.Info(fmt.Sprintf("用户组列表获取成功"))
		c.JSON(http.StatusOK, groups)
	}
}

func GroupInfo(c *gin.Context) {
	// 验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("用户组信息获取失败: %s", err.Error()))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("用户组信息获取失败: %s", err.Error()),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin {
		util.Error(fmt.Sprintf("用户组信息获取失败: %s", "权限不足"))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("用户组信息获取失败: %s", "权限不足"),
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		util.Error("用户组信息获取失败，参数错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户组信息获取失败，参数错误",
		})
		return
	}
	var request group.InfoRequest
	atoi, err := strconv.Atoi(id)
	if err != nil {
		util.Error(fmt.Sprintf("用户组信息获取失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("用户组信息获取失败，参数错误: %s", err.Error()),
		})
		return
	}
	request.ID = uint(atoi)
	res, err := request.Info()
	if err != nil {
		util.Error(fmt.Sprintf("用户组信息获取失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("用户组信息获取失败: %s", err.Error()),
		})
	} else {
		util.Info(fmt.Sprintf("用户组信息获取成功: %s", id))
		c.JSON(http.StatusOK, res)
	}
}
