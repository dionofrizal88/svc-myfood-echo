package controllers

import (
	"fmt"
	"net/http"
	"svc-myfood-echo/models"

	"github.com/labstack/echo"
)

func CheckoutBasketProduct(c echo.Context) error{
	uuidBasket := &models.CheckoutBasket{}
	orderName := c.FormValue("nama_pemesan")
	tableNumber := c.FormValue("nomor_meja")

	// store data order_pivot (nama, nomor_meja dan generate id)
	orderPivot, err := models.StoreOrderPivot(orderName, tableNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	fmt.Println("id order_pivot :", orderPivot)

	// get all id basket and change flag_order true
	c.Bind(uuidBasket)
	fmt.Println("bind", uuidBasket)
	uuidSlice := uuidBasket.UUID

	allIdBasket, err := models.FindIdBasket(uuidSlice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// storeOrder simpan semua data order (id_basket, id_order_pivot)
	order, err := models.StoreOrder(orderPivot, allIdBasket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, order)
}

	// c.Bind(uuidBasket)
	// fmt.Println("bind", uuidBasket)
	// a := uuidBasket.UUID
	// fmt.Println(a)