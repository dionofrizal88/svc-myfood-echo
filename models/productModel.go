package models

import (
	"net/http"
	"svc-myfood-echo/db"
)

type Products struct {
	ID	 		int64  `json:"id"`
	Kode 		string `json:"kode"`
	Nama  		string `json:"nama"`
	Harga  		string `json:"harga"`
	Gambar     	string `json:"gambar"`
	FlagReady  	string `json:"flag_ready"`
	FlagAktif   string `json:"flag_aktif"`
	UUID        string `json:"uuid"`
}

func FetchAllProduct(productName string)(Response, error){
	var obj Products
	var arrobj []Products
	var res Response
	
	con := db.CreateCon()

	nameFoodSearch := "%"
	if(len(productName) > 0){
		nameFoodSearch += productName + "%"
	}

	stmt := "SELECT kode, nama, harga, gambar, flag_ready, flag_aktif, uuid FROM products WHERE flag_aktif = true AND nama LIKE ?"

	rows, err := con.Query(stmt, nameFoodSearch)
	if err != nil{
		return res, err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Kode, &obj.Nama, &obj.Harga, &obj.Gambar, &obj.FlagReady, &obj.FlagAktif, &obj.UUID)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func FetchProductByUuid(uuid string)(Response, error){
	var obj Products
	var arrobj []Products
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT kode, nama, harga, gambar, flag_ready, flag_aktif, uuid from products WHERE uuid = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Query(uuid)
	if err != nil {
		return res, err
	}

	for result.Next() {
		err = result.Scan(&obj.Kode, &obj.Nama, &obj.Harga, &obj.Gambar, &obj.FlagReady, &obj.FlagAktif, &obj.UUID)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	if(len(arrobj) == 0){
		res.Data = nil
	}else{
		res.Data = arrobj
	}

	return res, nil
}

func StoreProduct(kode string, nama string, gambar string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT products (kode, nama, gambar) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(kode, nama, gambar)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}