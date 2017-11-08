package main

import (
	"fmt"
	"github.com/RedPatchTechnologies/postmurum-server/handlers"
	"github.com/RedPatchTechnologies/postmurum-server/middleware"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Config struct {
	Port         int  `env:"PORT" envDefault:"3000"`
	IsProduction bool `env:"PRODUCTION" envDefault:"DEBUG"`
}

func main() {

	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", cfg)

	router := gin.Default()
	store := sessions.NewCookieStore([]byte(handlers.RandToken(64)))
	router.Use(sessions.Sessions("postmurumsession", store))
	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.IndexHandler)
	router.GET("/login", handlers.LoginHandler)
	router.GET("/oauthcallback", handlers.AuthHandler)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := router.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.POST("/login", handlers.FieldHandler)
		authorized.POST("/submit", handlers.FieldHandler)
		authorized.POST("/read", handlers.FieldHandler)

		// nested group
		//testing := authorized.Group("testing")
		//testing.GET("/analytics", analyticsEndpoint)
	}

	/*
		authorized := router.Group("/sites")
		authorized.Use(middleware.AuthorizeRequest())
		{
			authorized.GET("/sites", handlers.FieldHandler)
		}
	*/

	router.Run("127.0.0.1:" + strconv.Itoa(cfg.Port))

	//router := web.New(Context{}). // Create your router
	//				Middleware(web.LoggerMiddleware).     // Use some included middleware
	// Middleware(web.ShowErrorsMiddleware). // ...
	// 	Middleware((*Context).SetHelloCount). // Your own middleware!
	//
	// 	Get("/", (*Context).SayHello) // Add a route

	//currentRoot, _ := os.Getwd()
	//router.Middleware(web.StaticMiddleware(path.Join(currentRoot, "public"), web.StaticOption{IndexFile: "index.html"}))

	/*
		router.Get("/users", (*Context).UsersList)
		router.Post("/users", (*Context).UsersCreate)
		router.Put("/users/:id", (*Context).UsersUpdate)
		router.Delete("/users/:id", (*Context).UsersDelete)
		router.Patch("/users/:id", (*Context).UsersUpdate)

	*/

	//http.ListenAndServe("localhost:3000", router) // Start the server!
}
