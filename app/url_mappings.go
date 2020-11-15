package app

import "github.com/ankitanwar/user-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUser)
}
