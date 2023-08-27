package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
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

func main() {
	port := os.Getenv("PORT")
	app := gin.Default()

	app.POST("/graphql", graphqlHandler())
	app.GET("/", playgroundHandler())
	app.Use(middlewareAuth())
	app.StaticFS("/images", http.Dir("assets/images"))
	fmt.Println("Playground started at http://localhost:" + port)
	app.Run(":" + port)
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
