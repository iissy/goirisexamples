package main

import "github.com/kataras/iris"

func main() {
    app := iris.Default()

    app.Post("/form_post", func(ctx iris.Context) {
        message := ctx.FormValue("message")
        nick := ctx.FormValueDefault("nick", "anonymous")

        ctx.JSON(iris.Map{
            "status":  "posted",
            "message": message,
            "nick":    nick,
        })
    })

    app.Run(iris.Addr(":8080"))
}