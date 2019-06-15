package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/web/{paramothersecond:path}", other2)

	v1 := app.Party("/v1/api")
	{ // braces are optional, it's just type of style, to group the routes visually.

		// http://localhost:8080
		v1.Get("/", func(ctx iris.Context) {
			ctx.HTML("Version 1 API. go to <a href='" + ctx.Path() + "/api" + "'>/api/users</a>")
		})

		usersAPI := v1.Party("/users")
		{
			// http://localhost:8080/api/users
			usersAPI.Get("/", func(ctx iris.Context) {
				ctx.Writef("All users")
			})
			// http://localhost:8080/api/users/42
			usersAPI.Get("/{userid:int}", func(ctx iris.Context) {
				ctx.Writef("user with id: %s", ctx.Params().Get("userid"))
			})
		}
	}

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

func other2(ctx iris.Context) {
	param := ctx.Params().Get("paramothersecond")
	ctx.Writef("from other2: %s", param)
}
