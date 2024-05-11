package main

import (
	"github.com/emmearn/gotlas/auth-middleware/internal"
)

func main() {
	middleware := internal.NewAuthMiddleware()

	err, api := internal.NewAPI(middleware)

	// qui devi passargli il router del server, dipendentemente dal server che usi (e.g. weaveworks)
	api.RegisterRoutes(server.Router)

	if err != server.Run(); err != nil {
		panic("uh oh")
	}
}
