package main

import (
	"fmt"
	//storage "github.com/anurakhan/KeyValueDB/Storage"
	db "github.com/anurakhan/KeyValueDB/DataBase"
)

func main() {
	dataBase := new(db.DataBase)
	dataBase.Init()
	dataBase.Put("Arlan", "Cool!")
	dataBase.Put("Zhomart", "Awesome!")
	dataBase.Put("Serik", "Karakurt!#$")
	val, ok := dataBase.Get("Zhomart")
	fmt.Println("The DataBase Value: ", val, "IsPresent? ", ok)
	val, ok = dataBase.Get("LOLOLOLOL")
	fmt.Println("The DataBase Value: ", val, "IsPresent? ", ok)

}
