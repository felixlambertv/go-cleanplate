package main

import (
	"github.com/felixlambertv/go-cleanplate/config"
	"github.com/felixlambertv/go-cleanplate/internal/app"
)

func main() {
	conf := config.GetInstance()

	app.Run(conf)
}
