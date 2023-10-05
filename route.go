package route

import (
  "uts/database"
  "uts/models"

  "github.com/gofiber/fiber/v2"
  )
// buatlah endpoint Insert Data sesuai dengan Colection Postman
func InsertData(c *fiber.Ctx) error {

 // Get the user data from the request body
 user := models.User{}
 err := json.NewDecoder(r.Body).Decode(&user)
 if err != nil {
  http.Error(w, err.Error(), http.StatusBadRequest)
  return
 }

 // Create a new user in the database
 database.DB.Create(&user)

 // Respond with the created user
 json.NewEncoder(w).Encode(user)

 return nil
}

// Lengkapi Code Berikut untuk untuk Mengambil data untuk semua user - user
func GetAllData(c *fiber.Ctx) error {

 var users []models.User

 // Get all users from the database
 database.DB.Find(&users)

 // Respond with all users
 return c.JSON(fiber.Map{
  "data": users,
 })
}

//Lengkapi Code berikut Untuk Mengambil data dari id_user berdasarkan Parameter

func GetUserByid(c *fiber.Ctx) error {

 var user models.User

 // Get the user ID from the parameter
 id_user := c.Params("id_user")

 // Get the user from the database
 database.DB.Where("id_user = ?", id_user).First(&user)

 // Respond with the user
 return c.JSON(fiber.Map{
  "data": user,
 })
}
