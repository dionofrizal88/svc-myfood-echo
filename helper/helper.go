package helper

import "svc-myfood-echo/db"

func GenerateId() int{
	con := db.CreateCon()

	id, err := con.Query("SELECT UUID_SHORT() AS uuidShort;")
	if err != nil {
		panic(err.Error())
	}
	var uuidShort int
	for id.Next() {
		err := id.Scan(&uuidShort)
		if err != nil {
			panic(err.Error())
		}
	}
	defer id.Close()
	return uuidShort
}