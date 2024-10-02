package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/haikelz/asmaul-husna-api/internal/entities"
	"github.com/haikelz/asmaul-husna-api/internal/services"
)

type BaseResponse struct {
	StatusCode int `json:"statusCode"`
}

type HomeEndpoints struct {
	All    string `json:"all"`
	Urutan string `json:"urutan"`
	Latin  string `json:"latin"`
}

type HomeResponse struct {
	Author     string        `json:"author"`
	Repository string        `json:"repository"`
	Endpoints  HomeEndpoints `json:"endpoints"`
}

type AllAsmaulHusnaResponse struct {
	BaseResponse
	Total int                     `json:"total"`
	Data  []*entities.AsmaulHusna `json:"data"`
}

func main() {
	app := fiber.New()

	// caching
	app.Use(cache.New())

	// compression
	app.Use(compress.New())

	// handle CORS
	app.Use(cors.New())

	// helmet
	app.Use(helmet.New())

	// logger
	app.Use(logger.New(logger.Config{Format: "${ip} ${pid} ${locals:requestid} ${status} - ${method} ${path}\n"}))

	// healthcheck
	app.Use(healthcheck.New())

	// monitoring API
	app.Get("/metrics", monitor.New())

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(HomeResponse{Author: "Haikel Ilham Hakim", Repository: "https://github.com/haikelz/asmaul-husna-api", Endpoints: HomeEndpoints{All: "/api/all", Urutan: "/api/:urutan", Latin: "/api/latin/:latin"}})
	})

	app.Get("/api/all", func(c *fiber.Ctx) error {
		asmaulHusna, err := services.GetAsmaulHusna()

		if len(asmaulHusna) <= 0 {
			services.InsertLoopAsmaulHusna()
		}

		if err != nil {
			return err
		}

		return c.JSON(AllAsmaulHusnaResponse{BaseResponse: BaseResponse{StatusCode: 200}, Total: 99, Data: asmaulHusna})
	})

	app.Listen(":5000")
}
