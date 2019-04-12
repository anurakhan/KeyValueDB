package main

import (
	"fmt"
	//storage "github.com/anurakhan/KeyValueDB/Storage"
	db "github.com/anurakhan/KeyValueDB/DataBase"
)

func main() {
	var dataBase db.IDataBase = new(db.DataBase)
	dataBase.InitFromDisk("DataBaseData")
	val, ok := dataBase.Get("Arlan")
	fmt.Println("The DataBase Value: ", val, "IsPresent? ", ok)
	val, ok = dataBase.Get("Zhomart")
	fmt.Println("The DataBase Value: ", val, "IsPresent? ", ok)
	val, ok = dataBase.Get(1)
	fmt.Println("The DataBase Value: ", val, "IsPresent? ", ok)

}
