package main

import "github.com/kataras/iris"
import "github.com/kataras/iris/context"
import "net/http"
import "log"

func main() {
    app := iris.Default()

    // Query string parameters are parsed using the existing underlying request object.
    // The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe.
    app.Get("/welcome", basicAuth(sayHello))
    app.Get("/welcome2", sayHello)

    app.Run(iris.Addr(":8080"))
}

func sayHello(ctx iris.Context) {
    firstname := ctx.URLParamDefault("firstname", "Guest")
    // shortcut for ctx.Request().URL.Query().Get("lastname").
    lastname := ctx.URLParam("lastname") 

    ctx.Writef("Hello %s %s", firstname, lastname)
}

func basicAuth(h context.Handler) context.Handler {
	return func(ctx iris.Context) {
		if 1 == 2 {
            log.Print("yes")
			h(ctx)
		} else {
            log.Print("no")
			http.Error(ctx.ResponseWriter(), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}