package main

import (
	"github.com/kataras/iris"

	"github.com/gorilla/securecookie"
)

var (
	// hashKey是必需的，用于使用HMAC验证cookie值。 建议使用32或64字节的密钥。
	hashKey = []byte("the-big-and-secret-fash-key-here")
	// blockKey是可选的，用于加密cookie值 - 将其设置为nil以不使用加密。
	// 如果设置，则长度必须对应于加密算法的块大小。
	// 对于默认使用的AES，有效长度为16,24或32字节，用于选择AES-128，AES-192或AES-256。
	blockKey = []byte("lot-secret-of-characters-big-too")
	sc       = securecookie.New(hashKey, blockKey)
)

func newApp() *iris.Application {
	app := iris.New()

	// Set A Cookie.
	app.Get("/cookies/{name}/{value}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		value := ctx.Params().Get("value")

		ctx.SetCookieKV(name, value, iris.CookieEncode(sc.Encode)) // <--

		ctx.Writef("cookie added: %s = %s", name, value)
	})

	// Retrieve A Cookie.
	app.Get("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		value := ctx.GetCookie(name, iris.CookieDecode(sc.Decode)) // <--

		ctx.WriteString(value)
	})

	// Delete A Cookie.
	app.Delete("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		ctx.RemoveCookie(name) // <--

		ctx.Writef("cookie %s removed", name)
	})

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}
