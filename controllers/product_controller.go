package controllers

import (
	"go-api-project/database"
	"go-api-project/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetProducts(ginCtx *gin.Context) {
	page := ginCtx.DefaultQuery("page","1")
	limit := ginCtx.DefaultQuery("limit","5")

	pageInt , err := strconv.Atoi(page) // >>> Atoi means -> ASCII to integer
	if(err != nil || pageInt < 1){	pageInt = 1}
	limitInt , err := strconv.Atoi(limit)
	if(err != nil || limitInt < 1){	limitInt = 5}

	offset := (pageInt - 1) * limitInt

	// >> Get Total Product
	var total int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM products").Scan(&total)
	if(err != nil){
		ginCtx.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return
	}

	// >>> Fetch Product List With Pagination
	rows, err := database.DB.Query("SELECT id,name,price FROM products LIMIT ? OFFSET ?",limitInt,offset)

	if(err != nil){
		ginCtx.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return
	}

	defer rows.Close()

	var products []models.ProductModel
	for rows.Next(){
		var p models.ProductModel
		rows.Scan(&p.Id,&p.Name,&p.Price)
		products = append(products, p)
	}

	ginCtx.JSON(http.StatusOK,gin.H{
		"data": products,
		"limit": limitInt,
		"page": pageInt,
		"total":total,
	})
}