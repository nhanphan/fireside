package main

import (
	"fmt"

	"github.com/nhanphan/fireside/backend/api"
	"github.com/nhanphan/fireside/backend/routes"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

func main() {
	setupConf()

	// set the configs
	iris.Config.Render.Template.Directory = "../frontend/templates"
	iris.Config.Render.Template.Layout = "layout.html"

	// set the favicon
	iris.Favicon("../frontend/public/images/favicon.ico")

	// set static folder(s)
	iris.Static("/public", "../frontend/public", 1)

	// set the global middlewares
	iris.Use(logger.New(iris.Logger))

	// set the custom errors
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", nil, iris.NoLayout)
	})

	// register the routes & the public API
	registerRoutes()
	registerAPI()

	// start the server
	iris.Listen(viper.GetString("host") + ":" + viper.GetString("port"))
}

func setupConf() {
	viper.SetDefault("host", "127.0.0.1")
	viper.SetDefault("port", "8080")

	viper.SetConfigType("json")
	viper.SetConfigName("config")

	// TODO: should add this in the future for production
	// viper.AddConfigPath("/etc/fireside/")

	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func registerRoutes() {
	// register index using a 'Handler'
	iris.Handle("GET", "/", routes.Index())

	// this is other way to declare a route
	// using a 'HandlerFunc'
	iris.Get("/about", routes.About)

	// Dynamic route

	iris.Get("/profile/:username", routes.Profile)("user-profile")
	// user-profile is the custom,optional, route's Name: with this we can use the {{ url "user-profile" $username}} inside userlist.html

	iris.Get("/all", routes.UserList)
}

func registerAPI() {
	// this is other way to declare routes using the 'API'
	iris.API("/users", api.UserAPI{})
}
