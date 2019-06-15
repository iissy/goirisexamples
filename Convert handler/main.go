package main

import (
	"github.com/kataras/iris"
	"net/http"
)

func main() {
    app := iris.New()

    sillyHTTPHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
            println(r.RequestURI)
    })

    sillyConvertedToIon := iris.FromStd(sillyHTTPHandler)
    // FromStd can take (http.ResponseWriter, *http.Request, next http.Handler) too!
	app.Use(sillyConvertedToIon)

    app.Run(iris.Addr(":8080"))
}