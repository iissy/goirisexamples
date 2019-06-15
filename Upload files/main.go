package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris"
)

// 5MB
const maxSize = 5 << 20

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))

	// 1. 第一种上传方式，token可以用来加密认证，这里并未使用到
	app.Get("/upload", func(ctx iris.Context) {
		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		ctx.View("upload_form.html", token)
	})

	// 处理上传请求，并保存文件到目录
	app.Post("/upload", iris.LimitRequestBodySize(maxSize+1<<20), func(ctx iris.Context) {
		file, info, err := ctx.FormFile("uploadfile")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}

		defer file.Close()
		fname := info.Filename

		out, err := os.OpenFile("./uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}
		defer out.Close()
		io.Copy(out, file)

		ctx.JSON(iris.Map{
			"status":  true,
			"message": "ok",
		})
	})

	// 2. 第二种文件上传方式，支持多文件
	app.Get("/upload2", func(ctx iris.Context) {
		ctx.View("upload_form2.html")
	})

	// 接收上传请求，支持多文件
	app.Post("/upload2", iris.LimitRequestBodySize(maxSize), func(ctx iris.Context) {
		ctx.UploadFormFiles("./uploads", beforeSave)
	})

	app.StaticWeb("/css", "./assets/css")
	app.StaticWeb("/js", "./assets/js")
	app.StaticWeb("/uploads", "./uploads")

	app.Run(iris.Addr(":8080"), iris.WithPostMaxMemory(maxSize))
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)
	file.Filename = ip + "-" + file.Filename

	ctx.Writef("|%s", "/uploads/"+file.Filename)
}
