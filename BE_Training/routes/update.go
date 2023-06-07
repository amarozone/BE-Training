package routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/databases"
	model "CRUD_API/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)
 
func UpdateLaptop(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "Laptops")
	
	laptopBrand := c.Param("laptopBrand")
	var laptop model.Laptops
	
	defer cancel()
	

	if err := c.BindJSON(&laptop); err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	return
	}
	
	edited := bson.M{"Brand": laptop.Brand, "Model": laptop.Model, "Year":laptop.Year, "Price":laptop.Price}
	
	result, err := postCollection.UpdateOne(ctx, bson.M{"Brand": laptopBrand}, bson.M{"$set": edited})
	
	res := map[string]interface{}{"data": result}
	
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	return
	}
	
	if result.MatchedCount < 1 {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
	return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": res})
   }
