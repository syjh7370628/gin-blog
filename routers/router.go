package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "gin-blog/docs"
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/export"
	"gin-blog/pkg/qrcode"
	"gin-blog/pkg/upload"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"
	"net/http"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//gin.SetMode(setting.ServerSetting.RunMode)
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	//新增获取token的方法
	r.GET("/auth", api.GetLogin)
	//验证码
	r.GET("/getCaptcha", api.GetCaptcha)
	r.GET("/verifyCaptcha", api.VerifyCaptcha)
	r.GET("/show/:source", api.GetCaptchaPng)

	// 加载static文件夹下所有的文件
	r.LoadHTMLGlob("views/**/**/*")
	//swagger接口
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//图片上传
	r.POST("/upload", api.UploadImage)

	admin := r.Group("/admin")
	{
		//后台首页
		admin.GET("/index", GetAdminIndex)
		//后台登陆页面
		admin.GET("/login", GetLogin)
		//后台博客列表
		admin.GET("/list", GetAdminBlogList)
		//写博客页面
		admin.GET("/writeBlog", GetWriteBlog)

		//后台首页
		admin.GET("/html/:path", GetAdminPath)
	}

	blog := r.Group("/blog")
	{
		//博客首页
		blog.GET("/index", GetBlogIndex)
		//博客首页
		blog.GET("/detail", GetBlogDetail)
		//博客友情链接
		blog.GET("/link", GetBlogLink)
		//博客留言
		blog.GET("/gustbook", GetBlogGustbook)
		//博客归档
		blog.GET("/archives", GetBlogArchives)
		//博客归档
		blog.GET("/hardware", GetBlogArchives) //博客归档
		blog.GET("/software", GetBlogArchives) //博客归档
		blog.GET("/life", GetBlogArchives)
		//博客里程碑
		blog.GET("/roadmap", GetBlogRoadmap)
		//博客搜索
		blog.GET("/search", GetBlogSearch)
		//关于我们
		blog.GET("/about", GetBlogAbout)

	}

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取用户列表
		apiv1.GET("/auths", api.GetAuths)
		//获取指定用户
		apiv1.GET("/auths/:id", api.GetAuth)
		//新建用户
		apiv1.POST("/auths", api.AddAuth)
		//更新指定用户
		apiv1.PUT("/auths/:id", api.EditAuth)
		//删除指定用户
		apiv1.DELETE("/auths/:id", api.DeleteAuth)

		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//获取标签列表
		apiv1.GET("/tags/:id", v1.GetTag)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.GET("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)

		//获取评论列表
		apiv1.GET("/comments", v1.GetComments)
		//获取指定评论
		apiv1.GET("/comments/:id", v1.GetComment)
		//新建评论
		apiv1.POST("/comments", v1.AddComment)
		//更新指定评论
		apiv1.PUT("/comments/:id", v1.EditComment)
		//删除指定评论
		apiv1.DELETE("/comments/:id", v1.DeleteComment)

		//获取留言列表
		apiv1.GET("/messages", v1.GetMessages)
		//获取指定留言
		apiv1.GET("/messages/:id", v1.GetMessages)
		//新建留言
		apiv1.POST("/messages", v1.AddMessage)
		//更新指定留言
		apiv1.PUT("/messages/:id", v1.EditMessage)
		//删除指定留言
		apiv1.DELETE("/messages/:id", v1.DeleteMessage)

		//获取分类列表
		apiv1.GET("/categorys", v1.GetCategorys)
		//获取指定分类
		apiv1.GET("/categorys/:id", v1.GetCategory)
		//新建分类
		apiv1.POST("/categorys", v1.AddCategory)
		//更新指定分类
		apiv1.PUT("/categorys/:id", v1.EditCategory)
		//删除指定分类
		apiv1.DELETE("/categorys/:id", v1.DeleteCategory)
	}

	return r
}
