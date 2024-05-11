package main

import (
	"github.com/emmearn/gotlas/auth-middleware/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	middleware := internal.NewAuthMiddleware()

	err, api := internal.NewAPI(middleware)

	router := gin.Default()
	api.RegisterRoutes(router)

	if err != router.Run(); err != nil {
		panic("uh oh")
	}
}
