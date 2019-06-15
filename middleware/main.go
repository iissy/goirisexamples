package main

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/hero"
    "time"
    "log"
)

const refreshEvery = 10 * time.Second

func main() {
    app := iris.New()
    app.Use(mymid)
    app.Use(iris.Cache(refreshEvery))

    helloHandler := hero.Handler(hello)
    // Per route middleware, you can add as many as you desire.
    app.Get("/builtin/{to:string}", helloHandler)

    // Per route middleware, you can add as many as you desire.
    app.Get("/benchmark", greet)

    // Listen and serve on http://0.0.0.0:8080
    app.Run(iris.Addr(":8080"))
}

func hello(to string) string {
    log.Print("i am hello")
	return "Hello " + to
}

func greet(ctx iris.Context) {
	ctx.Header("X-Custom", "my  custom header")
	ctx.Writef("Hello World! %s", time.Now())
}

func mymid(ctx iris.Context) {
	log.Print("i am mid")
    ctx.Next()
}