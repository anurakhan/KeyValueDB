package db

type IDataBase interface {
	Put(key, val string)
	Get(key string) (string, bool)
	Delete(key string)
}

type DataBase struct {
	hashTable map[string]string
}

//initialization of database hashtable that contains the key value data in memory.
func (dataBase *DataBase) Init() {
	dataBase.hashTable = make(map[string]string)

}

//database get value by key query.
func (dataBase *DataBase) Get(key string) (string, bool) {
	val, ok := dataBase.hashTable[key]
	return val, ok
}

//database add new key value query.
func (dataBase *DataBase) Put(key, val string) {
	dataBase.hashTable[key] = val
}

//delete key value record from database.
func (dataBase *DataBase) Delete(key string) {
	delete(dataBase.hashTable, key)
}
