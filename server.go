package main

import (

	// "gorm.io/driver/sqlite"

	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/subrotokumar/stellerlink-backend/graph"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func middlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		customHeader := c.Request.Header.Get("Auth")
		auth := os.Getenv("Authorization")
		if customHeader != auth {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Authorization failed",
			})
		}
		c.Next()
	}
}

func main() {
	app := gin.Default()
	app.Use(middlewareAuth())
	app.StaticFS("/assets", http.Dir("assets/images"))
	app.POST("/graphql", graphqlHandler())
	app.GET("/", playgroundHandler())
	app.Run(":8080")
}
