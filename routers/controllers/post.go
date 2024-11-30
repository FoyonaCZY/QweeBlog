package controllers

import (
	"fmt"
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/auth"
	"github.com/FoyonaCZY/QweeBlog/service/post"
	"github.com/FoyonaCZY/QweeBlog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostPublish(c *gin.Context) {
	//验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error("发布文章失败: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "发布文章失败: " + err.Error(),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin && reqUser.Group.Type != models.GroupTypeEditor {
		util.Error("发布文章失败: 权限不足")
		c.JSON(http.StatusForbidden, gin.H{
			"message": "发布文章失败: 权限不足",
		})
		return
	}

	var request post.CreateRequest
	if err := c.ShouldBindJSON(&request); err == nil {
		response, err := request.Create()
		if err != nil {
			util.Error("发布文章失败: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "发布文章失败: " + err.Error(),
			})
			return
		}
		util.Info("发布文章成功")
		c.JSON(http.StatusOK, response)
	} else {
		util.Error("发布文章失败: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "发布文章失败: " + err.Error(),
		})
	}
}

func PostList(c *gin.Context) {
	pageID, err := strconv.Atoi(c.Param("pageid"))
	if err != nil {
		util.Error(fmt.Sprintf("文章列表获取失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("文章列表获取失败，参数错误: %s", err.Error()),
		})
		return
	}

	res, err := post.List(pageID)
	if err != nil {
		util.Error("获取文章列表失败: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取文章列表失败: " + err.Error(),
		})
		return
	}
	util.Info("获取文章列表成功")
	c.JSON(http.StatusOK, res)
}

func PostDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		util.Error("文章信息获取失败，参数错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文章信息获取失败，参数错误",
		})
		return
	}
	var request post.DetailRequest
	atoi, err := strconv.Atoi(id)
	if err != nil {
		util.Error(fmt.Sprintf("文章信息获取失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("文章信息获取失败，参数错误: %s", err.Error()),
		})
		return
	}
	request.ID = uint(atoi)

	res, err := request.Detail()
	if err != nil {
		util.Error("获取文章详情失败: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取文章详情失败: " + err.Error(),
		})
		return
	}
	util.Info("获取文章详情成功")
	c.JSON(http.StatusOK, res)
}

func PostDelete(c *gin.Context) {
	//验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error("删除文章失败: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "删除文章失败: " + err.Error(),
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		util.Error("文章信息获取失败，参数错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文章信息获取失败，参数错误",
		})
		return
	}
	atoi, err := strconv.Atoi(id)
	if err != nil {
		util.Error(fmt.Sprintf("文章信息获取失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("文章信息获取失败，参数错误: %s", err.Error()),
		})
		return
	}
	deleteID := uint(atoi)

	if reqUser.Group.Type != models.GroupTypeAdmin && (reqUser.Group.Type != models.GroupTypeEditor || reqUser.ID != deleteID) {
		util.Error("删除文章失败: 权限不足")
		c.JSON(http.StatusForbidden, gin.H{
			"message": "删除文章失败: 权限不足",
		})
		return
	}

	var request post.DeleteRequest
	request.ID = deleteID
	res, err := request.Delete()
	if err != nil {
		util.Error("删除文章失败: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "删除文章失败: " + err.Error(),
		})
		return
	}
	util.Info("删除文章成功")
	c.JSON(http.StatusOK, res)
}

func PostUpdate(c *gin.Context) {
	//验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error("删除文章失败: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "删除文章失败: " + err.Error(),
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		util.Error("文章信息获取失败，参数错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文章信息获取失败，参数错误",
		})
		return
	}
	atoi, err := strconv.Atoi(id)
	if err != nil {
		util.Error(fmt.Sprintf("文章信息获取失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("文章信息获取失败，参数错误: %s", err.Error()),
		})
		return
	}
	updateID := uint(atoi)

	if reqUser.Group.Type != models.GroupTypeAdmin && (reqUser.Group.Type != models.GroupTypeEditor || reqUser.ID != updateID) {
		util.Error("删除文章失败: 权限不足")
		c.JSON(http.StatusForbidden, gin.H{
			"message": "删除文章失败: 权限不足",
		})
		return
	}

	var request post.UpdateRequest
	if err := c.ShouldBindJSON(&request); err == nil {
		request.ID = updateID
		res, err := request.Update()
		if err != nil {
			util.Error("更新文章失败: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "更新文章失败: " + err.Error(),
			})
			return
		}
		util.Info("更新文章成功")
		c.JSON(http.StatusOK, res)
	} else {
		util.Error("更新文章失败: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "更新文章失败: " + err.Error(),
		})
	}
}

func PostCount(c *gin.Context) {
	res, err := post.Count()
	if err != nil {
		util.Error("获取文章总数失败: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取文章总数失败: " + err.Error(),
		})
		return
	}
	util.Info("获取文章总数成功")
	c.JSON(http.StatusOK, res)
}
