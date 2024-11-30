package controllers

import (
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/auth"
	"github.com/FoyonaCZY/QweeBlog/service/category"
	"github.com/FoyonaCZY/QweeBlog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CategoryCreate(c *gin.Context) {
	//验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error("创建分类失败: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "创建分类失败: " + err.Error(),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin && reqUser.Group.Type != models.GroupTypeEditor {
		util.Error("创建分类失败: 权限不足")
		c.JSON(http.StatusForbidden, gin.H{
			"message": "创建分类失败: 权限不足",
		})
		return
	}

	var request category.CreateRequest
	if err := c.ShouldBindJSON(&request); err == nil {
		response, err := request.Create()
		if err != nil {
			util.Error("创建分类失败: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "创建分类失败: " + err.Error(),
			})
			return
		}
		util.Info("创建分类成功")
		c.JSON(http.StatusOK, response)
	} else {
		util.Error("创建分类失败: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "创建分类失败: " + err.Error(),
		})
	}
}

func CategoryUpdate(c *gin.Context) {
	//验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error("更新分类失败: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "更新分类失败: " + err.Error(),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin && reqUser.Group.Type != models.GroupTypeEditor {
		util.Error("更新分类失败: 权限不足")
		c.JSON(http.StatusForbidden, gin.H{
			"message": "更新分类失败: 权限不足",
		})
		return
	}

	var request category.UpdateRequest
	if err := c.ShouldBindJSON(&request); err == nil {
		response, err := request.Update()
		if err != nil {
			util.Error("更新分类失败: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "更新分类失败: " + err.Error(),
			})
			return
		}
		util.Info("更新分类成功")
		c.JSON(http.StatusOK, response)
	} else {
		util.Error("更新分类失败: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "更新分类失败: " + err.Error(),
		})
	}
}

func CategoryList(c *gin.Context) {
	res, err := category.List()
	if err != nil {
		util.Error("获取分类列表失败: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取分类列表失败: " + err.Error(),
		})
		return
	}
	util.Info("获取分类列表成功")
	c.JSON(http.StatusOK, res)
}

func CategoryDelete(c *gin.Context) {
	//验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error("删除分类失败: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "删除分类失败: " + err.Error(),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin && reqUser.Group.Type != models.GroupTypeEditor {
		util.Error("删除分类失败: 权限不足")
		c.JSON(http.StatusForbidden, gin.H{
			"message": "删除分类失败: 权限不足",
		})
		return
	}
	var request category.DeleteRequest
	atoi, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.Error("删除分类失败: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "删除分类失败: " + err.Error(),
		})
		return
	}
	request.ID = uint(atoi)

	response, err := request.Delete()
	if err != nil {
		util.Error("删除分类失败: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "删除分类失败: " + err.Error(),
		})
		return
	}
	util.Info("删除分类成功")
	c.JSON(http.StatusOK, response)
}
