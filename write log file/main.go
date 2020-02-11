package main

import (
	"io"
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

// 按天生成日志文件
func todayFilename() string {
	today := time.Now().Format("20060102")
	return today + ".log"
}

// 创建打开文件
func newLogFile() *os.File {
	filename := todayFilename()
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func main() {
	f := newLogFile()
	defer f.Close()

	app := iris.New()
	// 同时写文件日志与控制台日志
	app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))

	requestLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,
		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},
		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(requestLogger)

	app.Get("/info", func(ctx iris.Context) {
		ctx.Application().Logger().Infof("hello: %s", "i am info.")
		ctx.WriteString("write info")
	})

	app.Get("/error", func(ctx iris.Context) {
		ctx.Application().Logger().Errorf("hello: %s", "i am error.")
		ctx.WriteString("write error")
	})

	app.Run(
		iris.Addr(":8080"),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
