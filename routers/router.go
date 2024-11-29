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

			//更新用户信息
			user.PUT("/update", controllers.UserUpdate)

			//删除用户
			user.DELETE("/delete/:id", controllers.UserDelete)
		}
	}

	/*

		业务相关

	*/
	v1 := r.Group("/api/v1")
	{
		/*

			文章相关

		*/
		posts := v1.Group("/posts")
		{

			//获取文章列表
			posts.GET("/list", controllers.PostList)

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
	}
	return r
}
