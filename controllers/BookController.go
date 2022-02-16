package controllers

import (
	"bookManagementSystem/configs"
	"bookManagementSystem/dtos"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)


var bookCollection = configs.GetCollection(configs.DbClient, "book")
var validate = validator.New()


func CreateBook() gin.HandlerFunc{
	bookRequest := new(dtos.BookRequestDto)
	return func(c *gin.Context){

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := c.BindJSON(bookRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.BookErrorResponse{
				Status: http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		validatorError := validate.Struct(bookRequest)
		if validatorError != nil {
			c.JSON(http.StatusBadRequest, dtos.BookErrorResponse{
				Status: http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{"data": validatorError.Error()}})
			return
		}

		one, err := bookCollection.InsertOne(ctx, bookRequest)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, one)
	}
}
