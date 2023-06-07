package routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/databases"
	model "CRUD_API/model"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetLaptops(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "Laptops")
	laptopBrand := c.Param("laptopBrand")
	var result model.Laptops
	
	defer cancel()
	
   
	fmt.Println("hello",laptopBrand)
	err := postCollection.FindOne(ctx, bson.M{"Brand": laptopBrand}).Decode(&result)
	
	res := map[string]interface{}{"data": result}
	
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
   }

