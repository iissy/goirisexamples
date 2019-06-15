package main

import (
	"github.com/kataras/iris"
)

type picture struct {
	ID      string
	URL     string
	Subject string
}

func main() {
	app := iris.New()

	// 设置关注的视图目录，和文件后缀
	tmpl := iris.HTML("./views", ".html")
	// 默认的模版文件
	tmpl.Layout("shared/layout.html")
	// 是否每次请求都重新加载文件，这个在开发期间设置为true，在发布时设置为false
	// 可以方便每次修改视图文件而无需停止服务
	tmpl.Reload(true)
	// 设置页面的函数
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings, " + s + "!"
	})

	app.RegisterView(tmpl)

	// 包含部分视图，数据填充的页面
	app.Get("/", func(ctx iris.Context) {
		pic := picture{ID: "ueheh2yu", URL: "http://hrefs.cn", Subject: "go web is popular"}
		ctx.ViewData("title", "Home page")
		ctx.ViewData("part", pic)
		ctx.View("home/index.html")
	})

	// 无母模版的页面
	app.Get("/nolayout", func(ctx iris.Context) {
		// 不使用母模版
		ctx.ViewLayout(iris.NoLayout)
		if err := ctx.View("home/nolayout.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	// http://localhost:8080
	// http://localhost:8080/nolayout
	app.Run(iris.Addr(":8080"))
}
