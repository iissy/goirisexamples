package main

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/context"
)

func main() {
    app := iris.New()

    app.Get("/", func(ctx context.Context) /* or iris.Context, it's the same for Go 1.9+. */ {

        // request header "referer" or url parameter "referer".
        r := ctx.GetReferrer()
        switch r.Type {
        case context.ReferrerSearch:
            ctx.Writef("Search %s: %s\n", r.Label, r.Query)
            ctx.Writef("Google: %s\n", r.GoogleType)
        case context.ReferrerSocial:
            ctx.Writef("Social %s\n", r.Label)
        case context.ReferrerIndirect:
            ctx.Writef("Indirect: %s\n", r.URL)
        }
    })

    app.Run(iris.Addr(":8080"))
}