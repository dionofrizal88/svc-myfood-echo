package controllers

import (
	"net/http"
	"strconv"
	"svc-myfood-echo/models"

	"github.com/labstack/echo"
)

func FetchActiveBasket(c echo.Context) error{
	result, err := models.FetchActiveBasket()
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreProductToBasket(c echo.Context) error{
	uuidProduct := c.Param("uuid_product")
	jumlah := c.FormValue("jumlah")
	keterangan := c.FormValue("keterangan")

	conv_jumlah, err := strconv.Atoi(jumlah)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreProductToBasket(uuidProduct, conv_jumlah, keterangan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteProductToBasket(c echo.Context) error{
	uuid := c.Param("uuid_product")

	result, err := models.DeleteProductBasket(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
