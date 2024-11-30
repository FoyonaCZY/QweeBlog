package routers

import (
	"github.com/FoyonaCZY/QweeBlog/middlewares"
	"github.com/FoyonaCZY/QweeBlog/routers/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//设置为发布模式，避免日志冲突
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	/*

		用户相关

	*/
	user := r.Group("/user")
	{
		//用户注册
		user.POST("/register", controllers.UserRegister)

		//用户登录
		user.POST("/login", controllers.UserLogin)

		userProtected := user.Group("")
		{
			userProtected.Use(middlewares.JwtAuth())

			//获取用户信息
			user.GET("/info/:id", controllers.UserInfo)

			//获取用户列表
			user.GET("/list", controllers.UserList)

			//更新用户信息
			user.PUT("/update", controllers.UserUpdate)

			//删除用户
			user.DELETE("/delete/:id", controllers.UserDelete)
		}
	}
	//激活用户
	r.GET("/activate", controllers.UserActivate)

	/*

		用户组相关

	*/
	group := r.Group("/group")
	{
		group.Use(middlewares.JwtAuth())

		//创建用户组
		group.POST("/create", controllers.GroupCreate)

		//获取用户组列表
		group.GET("/list", controllers.GroupList)

		//删除用户组
		group.DELETE("/delete/:id", controllers.GroupDelete)

		//更新用户组
		group.PUT("/update", controllers.GroupUpdate)

		//获取用户组信息
		group.GET("/info/:id", controllers.GroupInfo)
	}

	/*

		配置相关

	*/
	config := r.Group("/config")
	{
		config.Use(middlewares.JwtAuth())

		//修改配置
		config.POST("/update", controllers.ConfigUpdate)
	}

	/*

		文章相关

	*/
	posts := r.Group("/posts")
	{
		//获取文章页数
		posts.GET("/count", controllers.PostCount)

		//获取文章列表
		posts.GET("/list/:pageid", controllers.PostList)

		//获取文章详情
		posts.GET("/detail/:id", controllers.PostDetail)

		postProtected := posts.Group("")
		{
			postProtected.Use(middlewares.JwtAuth())

			//发布文章
			posts.POST("/publish", controllers.PostPublish)

			//删除文章
			posts.DELETE("/delete/:id", controllers.PostDelete)

			//更新文章
			posts.PUT("/update/:id", controllers.PostUpdate)
		}
	}

	/*

		评论相关

	*/
	//comments := v1.Group("/comments")
	//{
	//	//发布评论
	//	comments.POST("/publish",middlewares.JwtAuthMiddleware(), controllers.CommentPublish)
	//
	//	//获取评论列表
	//	comments.GET("/listall", controllers.CommentListall)
	//
	//	//获取文章评论列表
	//	comments.GET("/list/:post_id", controllers.CommentList)
	//
	//	//删除评论
	//	comments.DELETE("/delete/:id", middlewares.JwtAuthMiddleware(),controllers.CommentDelete)
	//
	//	//更新评论
	//	//comments.PUT("/update/:id", middlewares.JwtAuthMiddleware(),controllers.CommentUpdate)
	//}
	//
	///*
	//
	//	标签相关
	//
	//*/
	//tags := v1.Group("/tags")
	//{
	//	//创建标签
	//	tags.POST("/create", middlewares.JwtAuthMiddleware(),controllers.TagCreate)
	//
	//	//获取标签列表
	//	tags.GET("/list", controllers.TagList)
	//
	//	//删除标签
	//	tags.DELETE("/delete/:id", middlewares.JwtAuthMiddleware(),controllers.TagDelete)
	//
	//	//更新标签
	//	tags.PUT("/update/:id", middlewares.JwtAuthMiddleware(),controllers.TagUpdate)
	//}
	return r
}
