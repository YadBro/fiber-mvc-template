package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/YadBro/fiber-mvc-template/bootstrap"
	"github.com/YadBro/fiber-mvc-template/pkg/env"
)

//go:embed views
var viewsfs embed.FS

func main() {
	app := bootstrap.NewApplication(viewsfs, env.Development)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "4000"))))
}
