package controllers

import (
	"context"
	"fmt"
	"net/http"
	"restapi/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type contentStruct struct {
	Title  string `json:"title"`
	Auhtor string `json:"author"`
}

func CreateContent(ctx *gin.Context) {

	mongoClient := mongo.DbConnect().Database("test")

	db := mongoClient.Collection("loader")

	var content contentStruct

	if err := ctx.BindJSON(&content); err != nil {
		ctx.JSON(http.StatusNotFound, &err)
		return
	}

	_, err := db.InsertOne(context.TODO(), content)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"message": "item created",
	})

}

func GetAll(ctx *gin.Context) {

	mongoClient := mongo.DbConnect().Database("test")

	db := mongoClient.Collection("loader")

	cursor, err := db.Find(context.TODO(), bson.D{})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, &err)
	}

	result := []bson.D{}

	if err = cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err, "Error --- Parsing")
	}

	ctx.JSON(http.StatusOK, &result)

	defer cursor.Close(context.TODO())

}

func UpdateContent() {
	fmt.Println("Update")
}
