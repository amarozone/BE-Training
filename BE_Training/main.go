package main
 
import (
    routes "CRUD_API/routes"
    "github.com/gin-gonic/gin"
)
 
func main() {
    router := gin.Default()
 
	router.POST("/laptops", routes.PostLaptops)
	router.GET("laptops/:laptopBrand", routes.GetLaptops)
	router.PUT("laptops/:laptopBrand", routes.UpdateLaptop)
	router.DELETE("laptops/:laptopBrand", routes.DeleteLaptop)
    router.Run("localhost: 3000")
	
}