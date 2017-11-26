package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Config struct {
	Port         int  `env:"PORT" envDefault:"3000"`
	IsProduction bool `env:"PRODUCTION" envDefault:false`
}

func main() {

	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf(" asdfsd dfdfddfdf%+v\n", err)
	}

	router := gin.Default()
	store := sessions.NewCookieStore([]byte(RandToken(64)))
	router.Use(sessions.Sessions("postmurumsession", store))

	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", IndexHandler)
	router.GET("/login", LoginHandler)
	router.GET("/oauthcallback", AuthHandler)

	//dev purposes
	router.GET("/org", OrgHandler)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := router.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthorizeRequest())
	{

		authorized.POST("/submit", SubmitHandler)
		authorized.POST("/read", FieldHandler)

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

	router.Run("0.0.0.0:" + strconv.Itoa(cfg.Port))
}
