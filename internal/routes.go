package internal

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/valianx/videos/internal/application/handlers"
	"log"
	"net/http"
	"os"
)

func Routes(port string) *gin.Engine {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	fmt.Printf("connect to port %s\n", port)

	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(static.Serve("/", static.LocalFile("./views", false)))
	r.NoRoute(static.Serve("/", static.LocalFile("/views", false)))

	r.POST("/login", AuthMiddleware.LoginHandler)

	auth := r.Group("/api")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", AuthMiddleware.RefreshHandler)

	auth.Use(AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", HelloHandler)

		//sistema usuarios
		auth.GET("/users", handlers.FindUsers)
		r.POST("/users", handlers.CreateUser)
		auth.GET("/users/:id", handlers.FindUser)
		auth.PATCH("/users/:id", handlers.UpdateUser)
		auth.DELETE("/users/:id", handlers.DeleteUser)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}


	return r
}
