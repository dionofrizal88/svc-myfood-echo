package controllers

import (
	"net/http"
	"svc-myfood-echo/models"

	"github.com/labstack/echo"
)

func FetchAllProduct(c echo.Context) error{
	productName := c.QueryParam("nama_makanan")

	result, err := models.FetchAllProduct(productName)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchProductByUuid(c echo.Context) error{
	uuid := c.Param("uuid")
	result, err := models.FetchProductByUuid(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if result.Data == nil{
		var res models.BadResponse
		res.Status = http.StatusBadRequest
		res.Message = "Failed"
		res.Errors = []string{
			"id tidak ditemukan",
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusOK, result)
}

func StoreProduct(c echo.Context) error{
	kode := c.FormValue("kode")
	nama := c.FormValue("nama")
	gambar := c.FormValue("gambar")

	result, err := models.StoreProduct(kode, nama, gambar)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchProductRecomendation(c echo.Context) error{

	result, err := models.FetchProductRecomendation()
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}