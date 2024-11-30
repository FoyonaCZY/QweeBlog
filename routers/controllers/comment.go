package controllers

import (
	"fmt"
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/auth"
	"github.com/FoyonaCZY/QweeBlog/service/comment"
	"github.com/FoyonaCZY/QweeBlog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentListAll(c *gin.Context) {
	// 获取评论列表
	res, err := comment.List()
	if err != nil {
		util.Error(fmt.Sprintf("获取评论列表失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("获取评论列表失败: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func CommentList(c *gin.Context) {
	// 获取评论列表
	postID, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		util.Error(fmt.Sprintf("获取评论列表失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("获取评论列表失败，参数错误: %s", err.Error()),
		})
		return
	}
	id := uint(postID)

	res, err := comment.ListByPostID(id)
	if err != nil {
		util.Error(fmt.Sprintf("获取评论列表失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("获取评论列表失败: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func CommentCreate(c *gin.Context) {
	// 验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("评论创建失败: %s", err.Error()))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("评论创建失败: %s", err.Error()),
		})
		return
	}

	var request comment.CreateRequest
	if err := c.ShouldBindJSON(&request); err == nil {
		request.UserID = reqUser.ID
		response, err := request.Create()
		if err != nil {
			util.Error(fmt.Sprintf("评论创建失败: %s", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("评论创建失败: %s", err.Error()),
			})
			return
		}
		util.Info("评论创建成功")
		c.JSON(http.StatusOK, response)
	} else {
		util.Error(fmt.Sprintf("评论创建失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("评论创建失败: %s", err.Error()),
		})
	}
}

func CommentDelete(c *gin.Context) {
	// 验证请求者权限
	reqUser, err := auth.CurrentUser(c)
	if err != nil {
		util.Error(fmt.Sprintf("评论删除失败: %s", err.Error()))
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("评论删除失败: %s", err.Error()),
		})
		return
	}
	if reqUser.Group.Type != models.GroupTypeAdmin {
		util.Error(fmt.Sprintf("评论删除失败: %s", "权限不足"))
		c.JSON(http.StatusForbidden, gin.H{
			"message": fmt.Sprintf("评论删除失败: %s", "权限不足"),
		})
		return
	}

	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.Error(fmt.Sprintf("评论删除失败，参数错误: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("评论删除失败，参数错误: %s", err.Error()),
		})
		return
	}
	var request comment.DeleteRequest
	request.ID = uint(commentID)
	res, err := request.Delete()
	if err != nil {
		util.Error(fmt.Sprintf("评论删除失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("评论删除失败: %s", err.Error()),
		})
	} else {
		util.Info(fmt.Sprintf("评论删除成功: %d", res.ID))
		c.JSON(http.StatusOK, res)
	}
}
