package controllers

import (
	"net/http"
	"strconv"
	"svc-myfood-echo/models"

	"github.com/labstack/echo"
)

func FetchAllEmployee(c echo.Context) error {
	result, err := models.FetchAllEmployee()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreEmployee(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	result, err := models.StoreEmployee(nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateEmployee(c echo.Context) error {
	id := c.Param("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateEmployee(conv_id, nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteEmployee(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)

	checkData, err := models.CheckData(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
 
	if checkData.Message == "Failed"{
		return c.JSON(http.StatusBadRequest, checkData)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteEmployee(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

