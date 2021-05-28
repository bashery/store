package main

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// hella
func assets() string {
	if os.Getenv("USERNAME") != "fedor" {
		return "/root/store/assets"
	}
	return "assets"
}

func main() {

	db := setdb()
	defer db.Close()
    
    e := echo.New()
    //SetupRoutes()
    //e.GET("/a", contentHandler, contentRewrite)
    

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Renderer = templ()

    e.Static("/a", assets())
	e.Static("/fs", photoFold())

	e.GET("/", home)
    e.GET("/sign", signPage)
	e.POST("/sign", signup)
	e.GET("/login", loginPage)
	e.POST("/login", login)
    e.GET("/mystore", myStores)
	e.GET("/stores", stores)
	e.GET("/acount/:name", acount)

    e.GET("/:catigory", getProds) // ?? 
	
    e.GET("/product/:id", getOneProd)
    //e.GET("/:catigory/:id", getOneProd) // whech is beter ? :catigory or /product ?
	e.GET("/upload", uploadPage)
	e.POST("/upload", upload)
    


    // not found pages
    e.GET("/:ok/:ok/:ok", notFoundPage)
    e.GET("/:ok/", notFoundPage)


	e.Logger.Fatal(e.Start(":8080"))
}

