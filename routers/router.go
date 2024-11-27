package routers

import (
	"github.com/FoyonaCZY/QweeBlog/routers/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//用户相关
	user := r.Group("/user")
	{
		//用户注册
		user.POST("/register", controllers.UserRegister)

		//用户登录
		user.POST("/login", controllers.UserLogin)
	}

	//protected routes
	v1 := r.Group("/api/v1")
	{

		//文章相关
		post := v1.Group("/post")
		{
			//发布文章
			post.POST("/publish", controllers.PostPublish)

			//获取文章列表
			post.GET("/list", controllers.PostList)

			//获取文章详情
			post.GET("/detail", controllers.PostDetail)

			//删除文章
			post.DELETE("/delete/:id", controllers.PostDelete)

			//更新文章
			post.PUT("/update/:id", controllers.PostUpdate)
		}
	}
	return r
}
