package main

import (
	"github.com/emmearn/gotlas/internal"
)

func main() {
	middleware := internal.NewAuthMiddleware()

	api := internal.NewAPI(middleware)

	// qui devi passargli il router del server, dipendentemente dal server che usi (e.g. weaveworks)
	api.RegisterRoutes(server.Router)

	if err != server.Run(); err != nil {
		panic("uh oh")
	}
}
