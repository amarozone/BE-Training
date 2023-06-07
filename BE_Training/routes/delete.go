package routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/databases"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)
 
func DeleteLaptop(c *gin.Context) {
 ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
 var DB = database.ConnectDB()
 laptopBrand := c.Param("laptopBrand")
 
 var postCollection = getcollection.GetCollection(DB, "Laptops") 
 defer cancel()
 fmt.Println("hello", laptopBrand)
 result, err := postCollection.DeleteOne(ctx, bson.M{"Brand": laptopBrand}) 
 res := map[string]interface{}{"data": result}
 
 fmt.Println("hello", laptopBrand)
 if err != nil {
 c.JSON(http.StatusInternalServerError, gin.H{"message": err})
 return
 }
 
 if result.DeletedCount < 1 {
 c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
 return
 }
 
 c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully", "Data": res})
}
