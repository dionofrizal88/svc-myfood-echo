package models

import (
	"fmt"
	"net/http"
	"strings"
	"svc-myfood-echo/db"
)

type ProductBasket struct {
	Kode 		string `json:"kode"`
	Nama  		string `json:"nama"`
	Harga  		string `json:"harga"`
	Gambar     	string `json:"gambar"`
	Jumlah 		int8   `json:"jumlah"`
	Keterangan  string `json:"keterangan"`
	UUID        string `json:"uuid"`
}

func FetchActiveBasket()(Response, error){
	var obj ProductBasket
	var arrobj []ProductBasket
	var res Response

	con := db.CreateCon()	

	stmt := "SELECT p.nama, p.harga, p.kode, p.gambar, b.jumlah, b.keterangan, b.uuid from baskets b JOIN products p on p.id = b.id_product WHERE b.flag_order = 0"

	rows, err := con.Query(stmt)
	if err != nil{
		return res, err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Nama, &obj.Harga, &obj.Kode, &obj.Gambar, &obj.Jumlah, &obj.Keterangan, &obj.UUID)
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

func StoreProductToBasket(uuidProduct string, jumlah int, keterangan string) (Response, error) {
	var res Response
	var idProduct int

	con := db.CreateCon()

	// find id product
	sqlProduct := "SELECT id FROM products WHERE uuid = ?"

	stmtProduct, err := con.Prepare(sqlProduct)
	if err != nil {
		return res, err
	}

	product, err := stmtProduct.Query(uuidProduct)
	if err != nil {
		return res, err
	}

	for product.Next() {
		err = product.Scan(&idProduct)
	}
	// -------------------
	sqlStatement := "INSERT INTO baskets(jumlah, id_product, keterangan) VALUES(?,?,?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(jumlah, idProduct, keterangan)
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

func FindIdBasket(uuidBaskets []string) ([]int, error){
	con := db.CreateCon()

	ids := strings.Join(uuidBaskets, "','")
	sqlRaw := fmt.Sprintf(`SELECT id FROM baskets WHERE uuid IN ('%s')`, ids)
	result, err := con.Query(sqlRaw)

	if err != nil {
		return nil, err
	}

	// change flag order
	UpdateFlagOrder(ids)

	var allId []int
	var idBasket int

	for result.Next(){
		err = result.Scan(&idBasket)
		if err != nil {
			return nil, err
		}
		allId = append(allId, idBasket)
	}

	return allId, nil
}

func UpdateFlagOrder(ids string)error{
	con := db.CreateCon()

	sqlStatement := fmt.Sprintf(`UPDATE baskets SET flag_order = ? WHERE uuid IN ('%s')`, ids)

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	stmt.Exec(true)

	return nil
}

func DeleteProductBasket(uuid string) (Response, error){
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM baskets WHERE uuid = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(uuid)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}




