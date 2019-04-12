package db

import (
	storage "github.com/anurakhan/KeyValueDB/Storage"
)

type IDataBase interface {
	Put(key, val interface{})
	Get(key interface{}) (interface{}, bool)
	Delete(key interface{})
	SaveToDisk() bool
	Init(path string)
	InitFromDisk(path string)
}

type DataBase struct {
	path      string
	hashTable map[interface{}]interface{}
}

//initialization of database hashtable that contains the key value data in memory.
func (dataBase *DataBase) Init(path string) {
	dataBase.hashTable = make(map[interface{}]interface{})
	dataBase.path = path
}

//initialization of database hashtable that contains the key value data in memory.
func (dataBase *DataBase) InitFromDisk(path string) {
	dataBase.path = path
	cntrl := storage.DataBaseDiskController{dataBase.path}
	data, _ := cntrl.ReadFromDisk()
	dataBase.hashTable = data
}

//database get value by key query.
func (dataBase *DataBase) Get(key interface{}) (interface{}, bool) {
	val, ok := dataBase.hashTable[key]
	return val, ok
}

//database add new key value query.
func (dataBase *DataBase) Put(key, val interface{}) {
	dataBase.hashTable[key] = val
}

//delete key value record from database.
func (dataBase *DataBase) Delete(key interface{}) {
	delete(dataBase.hashTable, key)
}

//methods that saves database data to disk.
func (dataBase *DataBase) SaveToDisk() bool {
	cntrl := storage.DataBaseDiskController{dataBase.path}
	ok := cntrl.SaveToDisk(dataBase.hashTable)
	return ok
}
