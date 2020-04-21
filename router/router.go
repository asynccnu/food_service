package router

import (
	"net/http"

	"github.com/asynccnu/food_service/handler/food"
	"github.com/asynccnu/food_service/handler/restaurant"
	"github.com/asynccnu/food_service/handler/sd"
	"github.com/asynccnu/food_service/handler/search"
	"github.com/asynccnu/food_service/router/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	sear := g.Group("/api/v1/search")
	{
		sear.GET("/food", search.SearchFood)
		sear.GET("/restaurant", search.SearchRestaurant)
		sear.GET("/hot", search.HotSearch)
	}

	rest := g.Group("/api/v1/restaurant")
	{
		rest.GET("/detail/:id", restaurant.Get)
		rest.GET("/list", restaurant.List)
		rest.POST("/random", restaurant.Random) //get改成post
	}

	foo := g.Group("/api/v1/food")
	{
		foo.GET("/recommend", food.Recommend)
	}

	return g
}
