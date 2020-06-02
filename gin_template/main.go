package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 静态文件：HTML页面上用到的样式文件.css js文件 图片
func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/static", "./statics")
	// gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func (str string) template.HTML {
			return template.HTML(str)
		},
	})
	// r.LoadHTMLFiles("templates/index.tmpl") // 模板解析
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func (c *gin.Context)  {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title":"liwenzhou.com",
		})
	})
	r.GET("/users/index", func (c *gin.Context)  {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	r.Run(":9090") // 启动server
}