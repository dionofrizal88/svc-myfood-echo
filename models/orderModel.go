package models

import (
	"fmt"
	"net/http"
	"svc-myfood-echo/db"
	"svc-myfood-echo/helper"
)

type CheckoutBasket struct {
	UUID     		[]string   `json:"uuid_keranjang" form:"uuid_keranjang"` 
	OrderName     	 string   `json:"nama_pemesan" form:"nama_pemesan"` 
	TableNumber      string   `json:"nomor_meja" form:"nomor_meja"` 
}

func StoreOrderPivot(orderName string, numberTable string) (int, error){
	con := db.CreateCon()

	id_generate := helper.GenerateId()

	sqlStatement := "INSERT order_pivot (id, nomor_meja, nama_pemesan) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return 0, err
	}

	stmt.Exec(id_generate, numberTable, orderName)

	return id_generate, nil
}

func StoreOrder(idOrderPivot int, idBasket []int) (Response, error){
	var res Response

	con := db.CreateCon()

	sqlOrder := "INSERT orders (id_order_pivot, id_baskets) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlOrder)
	if err != nil {
		return res, err
	}

	for _, id := range idBasket{
		fmt.Println(id)
		stmt.Exec(idOrderPivot, id)
		if err != nil {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"id order pivot": 0,
	}
	
	return res, nil
}

