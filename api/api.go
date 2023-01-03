package api

import (
	_ "github.com/shakhboznorbekov/token/user_api/api/docs"
	"github.com/shakhboznorbekov/token/user_api/api/handler"
	"github.com/shakhboznorbekov/token/user_api/config"
	"github.com/shakhboznorbekov/token/user_api/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(cfg *config.Config, r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandlerV1(cfg, storage)

	r.Use(customCORSMiddleware())

	r.POST("/login", handlerV1.Login)

	r.POST("/user", handlerV1.CreateUser)
	r.GET("/user/:id", handlerV1.GetUserById)
	r.GET("/user", handlerV1.GetUserList)
	// r.PUT("/user/:id", handlerV1.UpdateUser)
	// r.DELETE("/user/:id", handlerV1.DeleteUser)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
