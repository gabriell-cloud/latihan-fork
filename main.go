package main

import (
 "log"
 "github.com/gofiber/fiber/v2"
 "uts/database"
 "uts/route"
)

func main() {
 // Create a Fiber instance
 app := fiber.New()

 // Connect to the database
 database.Connect()

 // Define routes
 app.Post("/insert", route.InsertData)
 app.Get("/getData", route.GetAllData)
 app.Get("/getDataUser/:id_user", route.GetUserByid)
 app.Get("/delete/:id_user", route.Delete)
 app.Put("/update/:id_user", route.Update)

 // Start the Fiber application
 log.Fatal(app.Listen(":3000"))
}
