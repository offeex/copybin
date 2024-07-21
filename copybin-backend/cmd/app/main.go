package main

import (
	"copybin/internal/config"
	"copybin/internal/rest"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	router, err := rest.NewRouter(cfg)
	if err != nil {
		panic(err)
	}

	err = router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
