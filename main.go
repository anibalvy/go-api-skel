package main

import (
	"api_skel/config"
	"api_skel/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)



func main() {
    fmt.Println("starting server api")
    var err = config.LoadEnvs()
    if err != nil {
        fmt.Println("error loading config")
    }

    // create fiber server
    app := fiber.New()
    app.Use(logger.New())
    app.Use(requestid.New())
    app.Use(cors.New( cors.Config{
        AllowHeaders: "Origin, Content-type, Accept, Authorization",
        AllowCredentials: true,
    }))

    app.Get("/", func(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{ "status": 200, "data" : "api 001"})
	})


    app.Post("/get-token", config.Get_token)

    routes.Setup(app)


    app.Use(func(c *fiber.Ctx) error {
	    return c.Status(fiber.StatusNotFound).SendString("path not found")
    })

    log.Fatal(app.Listen(":3000"))

}
