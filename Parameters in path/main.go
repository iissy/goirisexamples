package main

import "github.com/kataras/iris"

import (
    "regexp"
)

func main() {
    app := iris.Default()

    // This handler will match /user/john but will not match neither /user/ or /user.
    app.Get("/user/{name}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        ctx.Writef("Hello %s", name)
    })

    // This handler will match /users/42
    // but will not match /users/-1 because uint should be bigger than zero
    // neither /users or /users/.
    app.Get("/users/{id:uint64}", func(ctx iris.Context) {
        id := ctx.Params().GetUint64Default("id", 0)
        ctx.Writef("User with ID: %d", id)
    })

    // len(name) <=255 otherwise this route will fire 404 Not Found
    // and this handler will not be executed at all.
    app.Get("/profile/{name:alphabetical max(255)}", func(ctx iris.Context){
        name := ctx.Params().Get("name")
        ctx.Writef("Hello %s", name)
    })

    latLonExpr := "^-?[0-9]{1,3}(?:\\.[0-9]{1,10})?$"
    latLonRegex, _ := regexp.Compile(latLonExpr)

    // Register your custom argument-less macro function to the :string param type.
    // MatchString is a type of func(string) bool, so we use it as it is.
    app.Macros().Get("string").RegisterFunc("coordinate", latLonRegex.MatchString)

    app.Get("/coordinates/{lat:string coordinate()}/{lon:string coordinate()}", func(ctx iris.Context) {
        ctx.Writef("Lat: %s | Lon: %s", ctx.Params().Get("lat"), ctx.Params().Get("lon"))
    })

    app.Macros().Get("string").RegisterFunc("range", func(minLength, maxLength int) func(string) bool {
        return func(paramValue string) bool {
            return len(paramValue) >= minLength && len(paramValue) <= maxLength
        }
    })
    
    app.Get("/limitchar/{name:string range(1,200) else 400}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        ctx.Writef(`Hello %s | the name should be between 1 and 200 characters length
        otherwise this handler will not be executed`, name)
    })

    app.Macros().Get("string").RegisterFunc("has", func(validNames []string) func(string) bool {
        return func(paramValue string) bool {
            for _, validName := range validNames {
                if validName == paramValue {
                    return true
                }
            }
    
            return false
        }
    })
    
    app.Get("/static_validation/{name:string has([kataras,gerasimos,maropoulos])}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        ctx.Writef(`Hello %s | the name should be "kataras" or "gerasimos" or "maropoulos"
        otherwise this handler will not be executed`, name)
    })

    // However, this one will match /user/john/send and also /user/john/everything/else/here
    // but will not match /user/john neither /user/john/.
    app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        action := ctx.Params().Get("action")
        message := name + " is " + action
        ctx.WriteString(message)
    })

    app.Run(iris.Addr(":8080"))
}