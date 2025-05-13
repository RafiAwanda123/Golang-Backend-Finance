package routes

import (
	"github.com/RafiAwanda123/Finance-UMKM/handlers"
	"github.com/RafiAwanda123/Finance-UMKM/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Auth routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", handlers.Signup)
		authGroup.POST("/login", handlers.Login)
	}

	// Protected API routes
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.JWTAuth())
	{
		// Finance routes
		financeGroup := apiGroup.Group("/finance")
		{
			financeGroup.GET("/", handlers.GetAllFinance)
			financeGroup.GET("/info", handlers.GetFinanceInfo)
			financeGroup.GET("/info/:id", handlers.GetFinanceByID)
			financeGroup.POST("/add", handlers.AddFinance)
			financeGroup.PUT("/edit", handlers.UpdateFinance)
			financeGroup.DELETE("/delete", handlers.DeleteFinance)
		}

		// Analysis route
		apiGroup.GET("/analysis/:item", handlers.AnalysisHandler)
	}

	return r
}
