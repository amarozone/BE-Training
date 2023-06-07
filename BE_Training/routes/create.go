package routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/databases"
	model "CRUD_API/model"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)
func PostLaptops(c *gin.Context) {
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "Laptops")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	laptop := new(model.Laptops)
	defer cancel()
	
	if err := c.BindJSON(&laptop); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"message": err})
	log.Fatal(err)
	return
	}
	fmt.Println("hello")
	laptopPayload := model.Laptops{
		Brand: laptop.Brand,
		Model: laptop.Model,
		Year: laptop.Year,
		Price: laptop.Price,
	}
	
	result, err := postCollection.InsertOne(ctx, laptopPayload)
	
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
   }