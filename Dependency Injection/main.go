package main

import "github.com/kataras/iris"
import "github.com/kataras/iris/hero"

func main() {
	app := iris.Default()
	
	//1. Path Parameters - Built'n Dependencies
	helloHandler := hero.Handler(hello)
	app.Get("/builtin/{to:string}", helloHandler)

	//2. Services - Static Dependencies
	hero.Register(&myTestService{
		prefix: "Service: Hello",
	})
	helloServiceHandler := hero.Handler(helloService)
	app.Get("/service/{to:string}", helloServiceHandler)

	//3. Per-Request - Dynamic Dependencies
	hero.Register(func(ctx iris.Context) (form LoginForm) {
		ctx.ReadForm(&form)
		return
	})
	loginHandle := hero.Handler(login)
	app.Post("/dynamic/login", loginHandle)

    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr(":8080"))
}

//1. Path Parameters - Built'n Dependencies
func hello(to string) string {
	return "Hello " + to
}

//2. Services - Static Dependencies
type Service interface {
	SayHello(to string) string
}

type myTestService struct {
	prefix string
}

func (s *myTestService) SayHello(to string) string {
	return s.prefix + " " + to
}

func helloService(to string, service Service) string {
	return service.SayHello(to)
}

//3. Per-Request - Dynamic Dependencies
type LoginForm struct {
	Username string
	Password string
}

func login(form LoginForm) string {
	return "Hello " + form.Username
}